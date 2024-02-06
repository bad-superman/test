package api

import (
	"fmt"

	"github.com/bad-superman/test/logging"
)

const (
	_sprdOrderURL                = "POST /api/v5/sprd/order"
	_sprdCancelOrderURL          = "POST /api/v5/sprd/cancel-order"
	_sprdMassCancelURL           = "POST /api/v5/sprd/mass-cancel"
	_sprdaMendCancelURL          = "POST /api/v5/sprd/amend-order"
	_sprdOrdersPendingURL        = "/api/v5/sprd/orders-pending"
	_sprdOrdersHistoryURL        = "/api/v5/sprd/orders-history"
	_sprdOrdersHistoryArchiveURL = "/api/v5/sprd/orders-history-archive"
	_sprdTradesURL               = "/api/v5/sprd/trades"
	_sprdSpreadsURL              = "/api/v5/sprd/spreads?baseCcy=%s&instId=%s&sprdId=%s&state=%s"
)

func (o *OkexClient) SprdSpreads(baseCcy, instId, sprdId, state string) ([]SpreadData, error) {
	res := new(SprdSpreadsResp)
	url := fmt.Sprintf(_sprdSpreadsURL, baseCcy, instId, sprdId, state)
	err := o.get(url, res)
	if err != nil {
		logging.Errorf("OkexClient|SprdSpreads error,err:%v", err)
		return res.Data, err
	}
	return res.Data, err
}
