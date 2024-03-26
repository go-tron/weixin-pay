package weixinPay

import (
	"fmt"
	"strconv"
)

type RefundReq struct {
	TransactionId     string  `json:"transactionId" validate:"required"`
	TxnAmount         float64 `json:"txnAmount" validate:"required"`
	OrigTransactionId string  `json:"origTransactionId" validate:"required"`
	OrigTxnAmount     float64 `json:"origTxnAmount" validate:"required"`
	NotifyUrl         string  `json:"notifyUrl"`
}

func (wxp *WeixinPay) Refund(req *RefundReq) (map[string]interface{}, error) {
	txnAmt, err := strconv.ParseFloat(fmt.Sprintf("%.2f", req.TxnAmount*100), 64)
	if err != nil {
		return nil, err
	}
	origTxnAmount, err := strconv.ParseFloat(fmt.Sprintf("%.2f", req.OrigTxnAmount*100), 64)
	if err != nil {
		return nil, err
	}
	var data = map[string]interface{}{
		"out_trade_no":  req.OrigTransactionId,
		"out_refund_no": req.TransactionId,
		"notify_url":    wxp.NotifyUrl,
		"amount": map[string]interface{}{
			"refund":   txnAmt,
			"total":    origTxnAmount,
			"currency": "CNY",
		},
	}

	if req.NotifyUrl != "" {
		data["notify_url"] = req.NotifyUrl
	}

	res, err := wxp.Execute("Refund", data, &map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return *res.(*map[string]interface{}), nil
}
