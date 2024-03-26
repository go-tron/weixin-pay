package weixinPay

func (wxp *WeixinPay) MarketingCallback(data map[string]interface{}) (map[string]interface{}, error) {
	res, err := wxp.Execute("MarketingCallback", data, &map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return *res.(*map[string]interface{}), nil
}

func (wxp *WeixinPay) MarketingStocks(stockId int) (map[string]interface{}, error) {
	var data = map[string]interface{}{
		"stock_id":            stockId,
		"stock_creator_mchid": wxp.MchId,
	}
	res, err := wxp.Execute("MarketingStocks", data, &map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return *res.(*map[string]interface{}), nil
}
