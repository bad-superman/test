package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/bad-superman/test/conf"
	"github.com/bad-superman/test/logging"
	v2 "github.com/bad-superman/test/sdk/dingtalk/v2"
	"github.com/bad-superman/test/sdk/okex"
	okex_api "github.com/bad-superman/test/sdk/okex/api"
)

const (
	_filledInfo   = "### 成交信息\n #### side:%s pos_side:%s price:%v\n"
	_positionInfo = "### 持仓信息\n #### long:%d short:%d\n"
	_pendingInfo  = "### 挂单信息\n #### ask pos_slide:%s price:%0.4f\n#### bid pos_slide:%s price:%0.4f\n"
)

func init() {
	conf.Init("./config.toml")
}

type GRIDTrade struct {
	client    *okex_api.OkexClient
	dClient   *v2.Manager
	instId    string
	longPos   int64
	shortPos  int64
	initPrice float64
}

func (g *GRIDTrade) GetBalance() (longPos int64, shortPos int64, err error) {
	postions, err := g.client.AccountPositions(okex.FuturesInstrument, []string{g.instId}, nil)
	if err != nil {
		logging.Errorf("GRIDTrade|GetBalance error,err:%v", err)
		return longPos, shortPos, err
	}
	for _, p := range postions {
		if p.InstID != g.instId {
			continue
		}
		if p.PosSide == okex.PositionLongSide {
			longPos = int64(p.Pos)
		}
		if p.PosSide == okex.PositionShortSide {
			shortPos = int64(p.Pos)
		}
	}
	g.longPos = longPos
	g.shortPos = shortPos
	return longPos, shortPos, nil
}

// 生成订单id
// side+time longbuy230205104101
func (g *GRIDTrade) GetClOrderID(side, posSide string) string {
	t := time.Now().Format("060102150405")
	return side + posSide + t
}

func (g *GRIDTrade) GetOrderPrice(price float64) (float64, float64) {
	ask := float64(price * 1.05)
	bid := float64(price / 1.05)
	// 空仓的情况，卖出价格上调
	if g.longPos == 0 && g.shortPos == 0 {
		ask = float64(price * 1.5)
	}
	return ask, bid
}

// 开多：买入开多（side 填写 buy； posSide 填写 long ）
// 开空：卖出开空（side 填写 sell； posSide 填写 short ）
// 平多：卖出平多（side 填写 sell；posSide 填写 long ）
// 平空：买入平空（side 填写 buy； posSide 填写 short ）
func (g *GRIDTrade) GetPosSide() (
	okex.PositionSide,
	okex.PositionSide) {
	askPosSide := okex.PositionShortSide
	bidPosSide := okex.PositionLongSide
	// 有多仓
	// 卖出平多
	if g.longPos > 0 {
		askPosSide = okex.PositionLongSide
	}
	// 有空仓
	// 买入平空
	if g.shortPos > 0 {
		bidPosSide = okex.PositionShortSide
	}
	return askPosSide, bidPosSide
}

func (g *GRIDTrade) InitOrderPrice() (price float64, err error) {
	// 获取当前挂单
	pendingOrders, err := g.client.PendingOrders(&okex_api.PendingOrderReq{
		InstId: g.instId,
	})
	if err != nil {
		logging.Errorf("GRIDTrade|InitOrder PendingOrders error,err:%v", err)
		return
	}
	// 取消当前订单
	cancelReq := make([]okex_api.CancelOrderReq, 0)
	for _, order := range pendingOrders {
		cancelReq = append(cancelReq, okex_api.CancelOrderReq{
			InstID: g.instId,
			OrdID:  order.OrdID,
		})
	}
	if len(cancelReq) > 0 {
		err = g.client.CancelBatchOrder(cancelReq)
		if err != nil {
			logging.Errorf("GRIDTrade|InitOrder CancelBatchOrder error,err:%v", err)
			return
		}
	}
	// 获取最后成交
	orders, err := g.client.FillsHistory(okex.FuturesInstrument, g.instId)
	if err != nil {
		logging.Errorf("GRIDTrade|InitOrder FillsHistory error,err:%v", err)
		return
	}
	if len(orders) == 0 {
		return 0, nil
	}
	if orders[0].FillSz != 1 {
		return 0, nil
	}
	price = float64(orders[0].FillPx)
	return price, nil
}

func (g *GRIDTrade) Trading() {
	// 需要初始化订单,价格
	initPrice, err := g.InitOrderPrice()
	if err != nil {
		logging.Panicf("InitOrderPrice error,err:%v", err)
	}
	// 指定初始价格
	if initPrice == 0 && g.initPrice != 0 {
		initPrice = g.initPrice
	} else {
		logging.Panic("InitOrderPrice is zero")
	}

	// 更新持仓
	g.GetBalance()
	askOrder, bidOrder, err := g.UpdateOrders(initPrice)
	if err != nil {
		logging.Panicf("UpdateOrders error,err:%v", err)
	}
	content := ""
	content += fmt.Sprintf(_positionInfo, g.longPos, g.shortPos)
	content += fmt.Sprintf(_pendingInfo,
		askOrder.PosSide, askOrder.Px,
		bidOrder.PosSide, bidOrder.Px,
	)
	// 发送钉钉通知
	mark := v2.NewMarkDown()
	mark.Markdown.Title = "那就不拖拉拉夫斯基"
	mark.Markdown.Text = content
	g.dClient.SendMsg(nil, mark)
	for {
		content := ""
		askOrderInfo, err := g.client.GetOrderInfo(g.instId, "", askOrder.ClOrdID)
		if err != nil {
			continue
		}
		bidOrderInfo, err := g.client.GetOrderInfo(g.instId, "", bidOrder.ClOrdID)
		if err != nil {
			continue
		}
		// 未成交
		if askOrderInfo.State == okex.OrderLive && bidOrderInfo.State == okex.OrderLive {
			time.Sleep(10 * time.Second)
			continue
		}
		// 卖单成了
		if askOrderInfo.State == okex.OrderFilled {
			initPrice = float64(askOrderInfo.Px)
			content += fmt.Sprintf(_filledInfo, askOrderInfo.Side, askOrderInfo.PosSide, askOrderInfo.FillPx)
		} else if askOrderInfo.State != okex.OrderCancel {
			err = g.client.CancelOrder(g.instId, askOrderInfo.OrdID, askOrderInfo.ClOrdID)
			if err != nil {
				time.Sleep(10 * time.Second)
				continue
			}
		}
		if bidOrderInfo.State == okex.OrderFilled {
			initPrice = float64(bidOrderInfo.Px)
			content += fmt.Sprintf(_filledInfo, bidOrderInfo.Side, bidOrderInfo.PosSide, bidOrderInfo.FillPx)
		} else if bidOrderInfo.State != okex.OrderCancel {
			err = g.client.CancelOrder(g.instId, bidOrderInfo.OrdID, bidOrderInfo.ClOrdID)
			if err != nil {
				time.Sleep(10 * time.Second)
				continue
			}
		}
		// 更新持仓
		_, _, err = g.GetBalance()
		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		}
		content += fmt.Sprintf(_positionInfo, g.longPos, g.shortPos)
		askTmpOrder, bidTmpOrder, err := g.UpdateOrders(initPrice)
		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		}
		askOrder, bidOrder = askTmpOrder, bidTmpOrder
		content += fmt.Sprintf(_pendingInfo,
			askOrder.PosSide, askOrder.Px,
			bidOrder.PosSide, bidOrder.Px,
		)
		// 发送钉钉通知
		mark := v2.NewMarkDown()
		mark.Markdown.Title = "那就不拖拉拉夫斯基"
		mark.Markdown.Text = content
		g.dClient.SendMsg(nil, mark)
	}
}

func (g *GRIDTrade) UpdateOrders(price float64) (askOrder, bidOrder okex_api.Order, err error) {
	askPrice, bidPrice := g.GetOrderPrice(price)
	askSide, bidSide := okex.OrderSell, okex.OrderBuy
	askPosSide, bidPosSide := g.GetPosSide()
	askOrderID := g.GetClOrderID(string(askSide), string(askPosSide))
	bidOrderID := g.GetClOrderID(string(bidSide), string(bidPosSide))
	logging.Infof("UpdateOrders,ask side:%v posSide:%s price:%0.4f", askSide, askPosSide, askPrice)
	logging.Infof("UpdateOrders,bid side:%v posSide:%s price:%0.4f", bidSide, bidPosSide, bidPrice)
	askOrder = okex_api.Order{
		InstID:  g.instId,
		ClOrdID: askOrderID,
		TdMode:  okex.TradeCrossMode,
		Side:    askSide,
		PosSide: askPosSide,
		Sz:      1,
		Px:      okex.JSONFloat64(askPrice),
		OrdType: okex.OrderLimit,
	}
	bidOrder = okex_api.Order{
		InstID:  g.instId,
		ClOrdID: bidOrderID,
		TdMode:  okex.TradeCrossMode,
		Side:    bidSide,
		PosSide: bidPosSide,
		Sz:      1,
		Px:      okex.JSONFloat64(bidPrice),
		OrdType: okex.OrderLimit,
	}
	_, err = g.client.TradeBatchOrders([]okex_api.Order{
		askOrder,
		bidOrder,
	})
	if err != nil {
		logging.Panicf("UpdateOrders error,err:%v", err)
		return askOrder, bidOrder, err
	}
	logging.Infof("UpdateOrders update ok,askOrderID:%s,bidOrderID:%s",
		askOrderID, bidOrderID)
	return askOrder, bidOrder, err
}

func main() {
	var initPrice = flag.Float64("price", 0, "-price init trade price")
	var instId = flag.String("instid", "", "-instid grid trade instance BTC-USD-230630")
	flag.Parse()
	fmt.Printf("init price:%0.4f,instid:%s", *initPrice, *instId)
	c := conf.GetConfig()
	gridTrade := &GRIDTrade{
		client:    okex_api.NewOkexClientByName(c, "test"),
		dClient:   v2.New(c.DTalkToken),
		instId:    *instId,
		initPrice: *initPrice,
	}
	defer func() {
		// 发送钉钉通知
		mark := v2.NewMarkDown()
		mark.Markdown.Title = "那就不拖拉拉夫斯基"
		mark.Markdown.Text = "## grid trading script exit!!!!!!"
		gridTrade.dClient.SendMsg(nil, mark)
	}()
	gridTrade.Trading()
}
