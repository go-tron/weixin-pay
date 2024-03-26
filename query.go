package weixinPay

type PayQueryReq struct {
	TransactionId string `json:"transactionId" validate:"required"`
}

func (wxp *WeixinPay) PayQuery(req *PayQueryReq) (*PayResult, error) {
	var data = map[string]interface{}{
		"mchid":        wxp.MchId,
		"out_trade_no": req.TransactionId,
	}
	res, err := wxp.Execute("PayQuery", data, &PayResult{})
	if err != nil {
		return nil, err
	}
	return res.(*PayResult), nil
}
