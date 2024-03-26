package weixinPay

import (
	"fmt"
	"strconv"
)

type ProfitsharingReceiversAddReq struct {
	ReceiveOpenId string `json:"receiveOpenId" validate:"required"`
}

func (wxp *WeixinPay) ProfitsharingReceiversAdd(req *ProfitsharingReceiversAddReq) (map[string]interface{}, error) {
	var data = map[string]interface{}{
		"appid":   wxp.AppId,
		"type":    "PERSONAL_OPENID",
		"account": req.ReceiveOpenId,
		//"name":            "",
		"relation_type": "USER",
		//"custom_relation": "",
	}
	res, err := wxp.Execute("ProfitsharingReceiversAdd", data, &map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return *res.(*map[string]interface{}), nil
}

type ProfitsharingOrdersReq struct {
	TransactionId       string  `json:"transactionId" validate:"required"`
	WeixinTransactionId string  `json:"weixinTransactionId" validate:"required"`
	ReceiveOpenId       string  `json:"receiveOpenId" validate:"required"`
	ReceiveAmount       float64 `json:"receiveAmount" validate:"required"`
}

func (wxp *WeixinPay) ProfitsharingOrders(req *ProfitsharingOrdersReq) (map[string]interface{}, error) {

	receiveAmount, err := strconv.ParseFloat(fmt.Sprintf("%.2f", req.ReceiveAmount*100), 64)
	if err != nil {
		return nil, err
	}

	var data = map[string]interface{}{
		"appid":          wxp.AppId,
		"transaction_id": req.WeixinTransactionId,
		"out_order_no":   req.TransactionId,
		"receivers": []map[string]interface{}{
			{
				"type":    "PERSONAL_OPENID",
				"account": req.ReceiveOpenId,
				//"name":        "",
				"amount":      receiveAmount,
				"description": "分账测试",
			},
		},
		"unfreeze_unsplit": true,
	}
	res, err := wxp.Execute("ProfitsharingOrders", data, &map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return *res.(*map[string]interface{}), nil
}
