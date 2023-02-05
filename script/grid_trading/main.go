package main

import (
	"time"

	"github.com/bad-superman/test/conf"
	"github.com/bad-superman/test/logging"
	"github.com/bad-superman/test/sdk/okex"
	okex_api "github.com/bad-superman/test/sdk/okex/api"
)

func init() {
	conf.Init("./config.toml")
}

type GRIDTrade struct {
	client   *okex_api.OkexClient
	instId   string
	longPos  int64
	shortPos int64
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
	ask := float64(int(price * 1.05))
	bid := float64(int(price * 0.95))
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
	price = float64(orders[0].FillPx)
	return price, nil
}

func (g *GRIDTrade) Trading() {
	// 需要初始化订单,价格
	initPrice, err := g.InitOrderPrice()
	if err != nil {
		logging.Panicf("InitOrderPrice error,err:%v", err)
	}

	// 更新持仓
	g.GetBalance()
	askOrderID, bidOrderID, err := g.UpdateOrders(initPrice)
	if err != nil {
		logging.Panicf("UpdateOrders error,err:%v", err)
	}
	for {
		askOrderInfo, err := g.client.GetOrderInfo(g.instId, "", askOrderID)
		if err != nil {
			continue
		}
		bidOrderInfo, err := g.client.GetOrderInfo(g.instId, "", bidOrderID)
		if err != nil {
			continue
		}
		// 未成交
		if askOrderInfo.State == okex.OrderLive && bidOrderInfo.State == okex.OrderLive {
			time.Sleep(30 * time.Second)
			continue
		}
		// 卖单成了
		if askOrderInfo.State == okex.OrderFilled {
			initPrice = float64(askOrderInfo.Px)
		} else {
			err = g.client.CancelOrder(g.instId, askOrderInfo.InstID, askOrderInfo.ClOrdID)
			if err != nil {
				time.Sleep(10 * time.Second)
				continue
			}
		}
		if bidOrderInfo.State == okex.OrderFilled {
			initPrice = float64(bidOrderInfo.Px)
		} else {
			err = g.client.CancelOrder(g.instId, bidOrderInfo.InstID, bidOrderInfo.ClOrdID)
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
		askTmpOrderID, bidTmpOrderID, err := g.UpdateOrders(initPrice)
		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		}
		askOrderID, bidOrderID = askTmpOrderID, bidTmpOrderID
	}
}

func (g *GRIDTrade) UpdateOrders(price float64) (askOrderID, bidOrderID string, err error) {
	askPrice, bidPrice := g.GetOrderPrice(price)
	askSide, bidSide := okex.OrderSell, okex.OrderBuy
	askPosSide, bidPosSide := g.GetPosSide()
	askOrderID = g.GetClOrderID(string(askSide), string(askPosSide))
	bidOrderID = g.GetClOrderID(string(bidSide), string(bidPosSide))
	logging.Infof("UpdateOrders,ask side:%v posSide:%s price:%f", askSide, askPosSide, askPrice)
	logging.Infof("UpdateOrders,bid side:%v posSide:%s price:%f", bidSide, bidPosSide, bidPrice)
	askOrder := okex_api.Order{
		InstID:  g.instId,
		ClOrdID: askOrderID,
		TdMode:  okex.TradeCrossMode,
		Side:    askSide,
		PosSide: askPosSide,
		Sz:      1,
		Px:      okex.JSONFloat64(askPrice),
		OrdType: okex.OrderLimit,
	}
	bidOrder := okex_api.Order{
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
		return askOrderID, bidOrderID, err
	}
	logging.Infof("UpdateOrders update ok,askOrderID:%s,bidOrderID:%s",
		askOrderID, bidOrderID)
	return askOrderID, bidOrderID, err
}

func main() {
	c := conf.GetConfig()
	gridTrade := &GRIDTrade{
		client: okex_api.NewOkexClientByName(c, "test"),
		instId: "BTC-USD-230331",
	}
	gridTrade.Trading()
}
