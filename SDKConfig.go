package weixinPay

type SDKConfig struct {
	Method string
	Url    string
}

var WeixinPublicCert = `-----BEGIN CERTIFICATE-----
MIID3DCCAsSgAwIBAgIUEiRdYwA9kEZAifa7zcAoXepN8FowDQYJKoZIhvcNAQEL
BQAwXjELMAkGA1UEBhMCQ04xEzARBgNVBAoTClRlbnBheS5jb20xHTAbBgNVBAsT
FFRlbnBheS5jb20gQ0EgQ2VudGVyMRswGQYDVQQDExJUZW5wYXkuY29tIFJvb3Qg
Q0EwHhcNMjEwMTIyMDc1MjU5WhcNMjYwMTIxMDc1MjU5WjBuMRgwFgYDVQQDDA9U
ZW5wYXkuY29tIHNpZ24xEzARBgNVBAoMClRlbnBheS5jb20xHTAbBgNVBAsMFFRl
bnBheS5jb20gQ0EgQ2VudGVyMQswCQYDVQQGDAJDTjERMA8GA1UEBwwIU2hlblpo
ZW4wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCvAndQZEaQUWNt0yol
y640BNo/fKC1YaYiaM6VTP5R91bcedi2CnzQaxqyWqqryaryZJYNlHamGRX79WeQ
mptWZtl9UfmtWXWgAibcqTsU2+OrE0tppNRQMxlMYiY3W+6Hc5k6Q4HG3UgtmErV
jN7EOAcOcpLAxBEn5yqwc60oH/7Qjw7yO6Vy3+gKYmTb2uRUrwEpMoi+7t28duZQ
ZbVXbXVxYeFSCrcA3VBDqyvf2lGb1P6fGnj263HoGXtKQ3xUS/NaPZbjYhXxqKsv
MvCA7plC6+1TkHvCRMmSbthcga+dgnta/70S/SwZyHOorUVoE6aRZY9dOxhwZua0
mJ2TAgMBAAGjgYEwfzAJBgNVHRMEAjAAMAsGA1UdDwQEAwIE8DBlBgNVHR8EXjBc
MFqgWKBWhlRodHRwOi8vZXZjYS5pdHJ1cy5jb20uY24vcHVibGljL2l0cnVzY3Js
P0NBPTFCRDQyMjBFNTBEQkMwNEIwNkFEMzk3NTQ5ODQ2QzAxQzNFOEVCRDIwDQYJ
KoZIhvcNAQELBQADggEBAJNGlUwliCdSufxeBYAaB52SGtBe+jDWRkOHJZVbM/vi
jdNG53rw3hXwujMYKAnCC089WloDjqPa9IhXrN80YGYWIa7j0tKQiilIqnUumW8e
pT6i8tr+EKg0kHkKi0JUh6KRSEbjLyfLNeSGewOgbGcwpN2iIj4VTxmoKZaEfLVL
m7ERwKcyqsM/4Wce3n7gLHnBcuM93XsiU/Ztb9jt4/6KzxsWm9Rk4iwUdXh9+Rrs
4LeccpVhDHIg9mrkEt22TLL96ZPQd3b6VmSy95WcrWMqJFfWCXdCr/+9h7CVxs9e
PrXxzWPLWTfpX6GtSl/x+eUADs5iwDLANzVb4j2KkmU=
-----END CERTIFICATE-----`

var SDKConfigMap = map[string]*SDKConfig{
	"GetCert": {
		Method: "GET",
		Url:    "https://api.mch.weixin.qq.com/v3/certificates",
	},
	"JsApiPay": {
		Method: "POST",
		Url:    "https://api.mch.weixin.qq.com/v3/pay/transactions/jsapi",
	},
	"PayQuery": {
		Method: "GET",
		Url:    "https://api.mch.weixin.qq.com/v3/pay/transactions/out-trade-no/{out_trade_no}",
	},
	"Refund": {
		Method: "POST",
		Url:    "https://api.mch.weixin.qq.com/v3/refund/domestic/refunds",
	},
	"RefundQuery": {
		Method: "GET",
		Url:    "https://api.mch.weixin.qq.com/v3/refund/domestic/refunds/{out_refund_no}",
	},
	"MarketingCallback": {
		Method: "POST",
		Url:    "https://api.mch.weixin.qq.com/v3/marketing/favor/callbacks",
	},
	"MarketingStocks": {
		Method: "GET",
		Url:    "https://api.mch.weixin.qq.com/v3/marketing/favor/stocks/{stock_id}",
	},
	"ProfitsharingReceiversAdd": {
		Method: "POST",
		Url:    "https://api.mch.weixin.qq.com/v3/profitsharing/receivers/add",
	},
	"ProfitsharingOrders": {
		Method: "POST",
		Url:    "https://api.mch.weixin.qq.com/v3/profitsharing/orders",
	},
}
