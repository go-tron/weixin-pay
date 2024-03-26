package weixinPay

import (
	"fmt"
	"github.com/go-tron/local-time"
	"github.com/go-tron/random"
	"strconv"
)

type JsApiPayReq struct {
	TransactionId string                 `json:"transactionId" validate:"required"`
	TxnAmount     float64                `json:"txnAmount" validate:"required"`
	OpenId        string                 `json:"openId" validate:"required"`
	Description   string                 `json:"description" validate:"required"`
	ExtraData     map[string]interface{} `json:"extraData"`
	ExpireTime    *localTime.Time        `json:"expireTime"`
	NotifyUrl     string                 `json:"notifyUrl"`
}

type JsApiPayResult struct {
	PrepayId string `json:"prepay_id"`
}

type JsApiPayRes struct {
	AppId     string `json:"appId"`
	Timestamp int64  `json:"timestamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

func (wxp *WeixinPay) JsApiPay(req *JsApiPayReq) (*JsApiPayRes, error) {

	txnAmt, err := strconv.ParseFloat(fmt.Sprintf("%.2f", req.TxnAmount*100), 64)
	if err != nil {
		return nil, err
	}

	var data = map[string]interface{}{
		"appid":        wxp.AppId,
		"mchid":        wxp.MchId,
		"description":  req.Description,
		"attach":       req.Description,
		"out_trade_no": req.TransactionId,
		"notify_url":   wxp.NotifyUrl,
		"payer": map[string]interface{}{
			"openid": req.OpenId,
		},
		"amount": map[string]interface{}{
			"total":    txnAmt,
			"currency": "CNY",
		},
		//"settle_info": map[string]interface{}{
		//	"profit_sharing": true,
		//},
	}

	if req.ExtraData != nil {
		if goodsTag := req.ExtraData["goodsTag"]; goodsTag != nil {
			data["goods_tag"] = goodsTag
		}
	}

	if req.ExpireTime != nil {
		data["time_expire"] = req.ExpireTime.RFC3339()
	}

	if req.NotifyUrl != "" {
		data["notify_url"] = req.NotifyUrl
	}

	res, err := wxp.Execute("JsApiPay", data, &JsApiPayResult{})
	if err != nil {
		return nil, err
	}
	result := res.(*JsApiPayResult)

	jsApiConfig := &JsApiPayRes{
		AppId:     wxp.AppId,
		Timestamp: localTime.Now().Unix(),
		NonceStr:  random.String(10),
		Package:   "prepay_id=" + result.PrepayId,
		SignType:  "RSA",
		PaySign:   "",
	}

	var message = wxp.AppId + "\n" + strconv.FormatInt(jsApiConfig.Timestamp, 10) + "\n" + jsApiConfig.NonceStr + "\n" + jsApiConfig.Package + "\n"
	paySign, err := wxp.Sign(message)
	if err != nil {
		return nil, ErrorSign
	}
	jsApiConfig.PaySign = paySign
	return jsApiConfig, nil
}
