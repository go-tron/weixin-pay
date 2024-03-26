package weixinPay

type RefundQueryReq struct {
	TransactionId string `json:"transactionId" validate:"required"`
}

func (wxp *WeixinPay) RefundQuery(req *RefundQueryReq) (*RefundResult, error) {
	var data = map[string]interface{}{
		"out_refund_no": req.TransactionId,
	}
	res, err := wxp.Execute("RefundQuery", data, &RefundResult{})
	if err != nil {
		return nil, err
	}
	return res.(*RefundResult), nil
}
