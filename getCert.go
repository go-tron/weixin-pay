package weixinPay

func (wxp *WeixinPay) GetCert() (*CertificateData, error) {
	res, err := wxp.Execute("GetCert", nil, &CertificateData{})
	if err != nil {
		return nil, err
	}
	return res.(*CertificateData), nil
}
