package weixinPay

type merchants interface {
	GetMerchantById(string) (*WeixinPay, error)
}

type Merchants struct {
	Merchants merchants
}

func (m *Merchants) VerifySign(appId string, timestamp, nonce, body, sign string) error {
	merchant, err := m.Merchants.GetMerchantById(appId)
	if err != nil {
		return err
	}
	return merchant.VerifySign(timestamp, nonce, body, sign)
}
func (m *Merchants) PayQuery(appId string, req *PayQueryReq) (*PayResult, error) {
	merchant, err := m.Merchants.GetMerchantById(appId)
	if err != nil {
		return nil, err
	}
	return merchant.PayQuery(req)
}
func (m *Merchants) Refund(appId string, req *RefundReq) (map[string]interface{}, error) {
	merchant, err := m.Merchants.GetMerchantById(appId)
	if err != nil {
		return nil, err
	}
	return merchant.Refund(req)
}
func (m *Merchants) RefundQuery(appId string, req *RefundQueryReq) (*RefundResult, error) {
	merchant, err := m.Merchants.GetMerchantById(appId)
	if err != nil {
		return nil, err
	}
	return merchant.RefundQuery(req)
}
func (m *Merchants) JsApiPay(appId string, req *JsApiPayReq) (*JsApiPayRes, error) {
	merchant, err := m.Merchants.GetMerchantById(appId)
	if err != nil {
		return nil, err
	}
	return merchant.JsApiPay(req)
}

func (m *Merchants) RefundResult(appId string, req *Resource) (*RefundResult, error) {
	merchant, err := m.Merchants.GetMerchantById(appId)
	if err != nil {
		return nil, err
	}
	return merchant.RefundResult(req)
}

func (m *Merchants) PayResult(appId string, req *Resource) (*PayResult, error) {
	merchant, err := m.Merchants.GetMerchantById(appId)
	if err != nil {
		return nil, err
	}
	return merchant.PayResult(req)
}
