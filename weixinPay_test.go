package weixinPay

import (
	"encoding/json"
	localTime "github.com/go-tron/local-time"
	"github.com/go-tron/logger"
	"testing"
	"time"
)

var weixin = New(&Config{
	AppId:         "wx6c8124f1fbafb1f3",
	MchId:         "1603022966",
	CertSN:        "5EDCD372816E7CBAD91804CA1EE3BC566E145CEC",
	PrivateKeyPem: "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC1UQZcO5d6Emzp\nne0L2ZxdCjbHc4nkFNTgqfuEwuDScClZgF+KXn2w91/5GsD3wLD6dLl6cD3ToFb7\np25USCgejbyYza//c/PlJUN84fYtwKd84ILweASq+028tMWbGhbPOw4R5QN8FHlZ\nLNYG2Njaj87OvP/r0B16Cu8DTA/rDCcl/5HZlZsMypKHDmxtDwBRPTYArFvaZZu1\nBeX7wcVdW5LU90VEaPy6l3ERwesiUEJ0gQjZFBHiyF/iAga2pFrZK7rGQ3ryGl0H\n6TesqFsp6BzUGW63M6kRWmO/b60DaJFKqH7hFDFxS54pxgfZv4HEr9Ka4XjaNoi9\nD1TPNne7AgMBAAECggEBAIgWcWSmoZHQ5IgPrYx4X0EB2o2m7XcQH5skWhCSSDYK\nUy7HCG6Nb190vEb2yqDpsqu5EkGQnwcI0GB/kXKW2e3cyhISR6e/Ou7hIh9IZgJ/\nF/bFd+HO4woGJpmdQLeiRD5z/6J0tkHaCB5jZZysA09AIqPO/XLbuFQSgsSBwWzT\nZdcztWSrP8WKdmFwsKKk1pgCrzS7wKwtjSvzcUF/6oOyyJBuPVZRviGdlWQ8tw2V\nupsQg/miiLvZczq41v4nAqpvlud7QtLD8ZR+b9RqWcf9B83cCJWb0xKDX2JUy0XM\nv5tAquMwPieLt2dYBe6w8ldk3ND9q/LfpKXPYq8M85kCgYEA2tEEv1fmLbuIhsPC\np8RtjGE8oPpRZnnyx4q9W9lv24WjeGl8x34VU5IiC5mUC4MaxfqGZCeQPjlFCFBH\nOlIxAKBGLH1MNUzLe3KgX6TjDA3cPAv9/dmvN9IeQFVTYOKwDzYnEHBIp3y7EjgG\nPqZR9zpY15afWXLtGVIa5nKdh58CgYEA1CCtgJz2pawC6V4N2dLGCJI6SEiWQBN6\nsi7ULZ2PRKq7CtYyMASQ5MHUOIUUlBwHDoDH1m5lal1LHkyJX6FrmUrQbecck2JK\nxeoIfVVyGYyVG/IvjGYFl3NsJ6GxWUQQVfJXpk44scvV5qWF5/yK8inqn54RKkB4\nIGQ/A1SkSmUCgYEAtqFb8AQCNsteEPTUw8e7kz4ZJ1all/1Sd99BWqbpqHQq0zZg\nEfUXAbBnP/1Hxi//qZwGjRNEXdrY1i6CtJejFJ2w1DMj5xyYfQlX91wcsJPk7C0q\nKbSFfPafjxxoQeYSAjA1fI/q4/fD/1nJRIL2yHznZ9DsYPD+GXMgxpSFDIsCgYAk\nDD2Pypy7gKSqBbqy1neiwz62Q+eMkgLavsx9x/WtxJmueMHkmRIKXcnzpOHfXXfx\nhf7vuKjxT1NRnc4Ge0busOEEnC6l+SEdyuyQZ/HQ16wLKLfd0wSGPS4W+gpKUh+4\ni0tLzqUhybLa1CwSRT9Tcb4WS+U82eHQF1kB9uNIZQKBgFL/QM46hwFGgn2aP/pm\neTkV0c6eJ1+/witZDJnW7YuXssCNHgls5O5LpJN/POGmwDzpfJkf4owJWPmlG4Lq\ndDwQkrq+N+unW5EJ44xAl4w4qFMIxWZc9/DTy4o5w3N1pZfNIBn04qm9nvbSnu7D\nqEs6wPphZPu4sGyEnZgZqJyX\n-----END PRIVATE KEY-----\n",
	APIV3Key:      "bHpnmR1h66DkWMVqtsBoDKeIMFULAXpm",
	Logger:        logger.NewZap("weixinPay", "info"),
})

func TestWeixinPay_MarketingCallback(t *testing.T) {
	result, err := weixin.MarketingCallback(map[string]interface{}{
		"mchid":      weixin.MchId,
		"notify_url": "https://weixin.eioos.com/marketing/callback/" + weixin.AppId,
		"switch":     true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}

func TestWeixinPay_GetCert(t *testing.T) {
	result, err := weixin.GetCert()
	if err != nil {
		t.Fatal(err)
	}
	var certificate = result.Data[0].EncryptCertificate
	pem, err := AESGCMDecrypter(weixin.APIV3Key, certificate.Nonce, certificate.Ciphertext, certificate.AssociatedData)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("pem", string(pem))
}

func TestWeixinPay_JsApiPay(t *testing.T) {
	result, err := weixin.JsApiPay(&JsApiPayReq{
		TransactionId: "32854923295488778251",
		TxnAmount:     1.00,
		OpenId:        "oasi95rPit953LHRYfaifGnTuqgs",
		Description:   "测试",
		ExpireTime:    localTime.Now().Add(time.Duration(time.Minute)).Ptr(),
		NotifyUrl:     "https://weixin.eioos.com/result/" + weixin.AppId,
	})
	if err != nil {
		t.Fatal(err)
	}

	s, _ := json.Marshal(result)
	t.Log("result", string(s))
}

func TestWeixinPay_PayQuery(t *testing.T) {
	result, err := weixin.PayQuery(&PayQueryReq{
		TransactionId: "12854923295488778251",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}

func TestWeixinPay_Refund(t *testing.T) {
	result, err := weixin.Refund(&RefundReq{
		TransactionId:     "228549232954887782494",
		TxnAmount:         0.12,
		OrigTransactionId: "12854923295488778249",
		OrigTxnAmount:     1.00,
		NotifyUrl:         "https://weixin.eioos.com/result/" + weixin.AppId,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}

func TestWeixinPay_RefundQuery(t *testing.T) {
	result, err := weixin.RefundQuery(&RefundQueryReq{
		TransactionId: "1355914720434982912",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}

func TestWeixinPay_VerifySign(t *testing.T) {
	var (
		timestamp = "1612164325"
		nonce     = "1sPICivwYIlWpoFL4H31eb6yPQVfDoaj"
		sign      = "OBcsIdHRpZdvzMfgFOStp8ZcCxZsPMagOoxhe6l+fZ3Y79RsQFmrx4yUozgOF8QfnOTWwg/hs0bk1duZepLSS+1yXTkuJxPJ8cK440JuE2Sty2DaDVVTkQ6XnAjplotQglGJorq9qS8l5NkIhzz6kCtG/SskAzTbIsZvUOMcCetBMtlmPRkE9Vt2SFaB3S+PNnwtU6QL2kSr3znOw0eRKsSd/21OT7zt4cnjjH4ync4O0jrpe+exzu/rp1KMSautDXX6n+ZC9sy76JJBcaKbeeRkwDiJ5TIwWqDzyczRTntoenbFdN/+CpmCqZlgRb9sExTHr1xK2qnRsmgI85SsmA=="
		body      = `{"id":"92815fd7-a930-5214-bf2b-a12190a07711","create_time":"2021-02-01T15:25:25+08:00","resource_type":"encrypt-resource","event_type":"COUPON.USE","summary":"代金券核销通知","resource":{"original_type":"coupon","algorithm":"AEAD_AES_256_GCM","ciphertext":"WZPiMxUjBLihsuDF703s+RhdnSa9RlShXtrbkkiUWFjwLvQ25y5BkUd/fGGhlw+daGq65psQFNFgt8NTBIxZunIhocpIgUX+Z4mxiZ1XUmQNkljqUix+1sd/bkPD9nAEnyln+6DSgXUlIlUoYUbI+tQ8N13ld+qMMcqGF7fwdNsPj2/2Z4hdGmIqi9p1guzCR9fmiy6yzCuT/8NMRyVnXgIVBxgDkdunugh4d50OdMjV4cEhqeNRlY1BrBAuEUQeMFxbXMvtJVsdjOt1t5Il0gWAkBDw5U3sPS6LIYfo240KsVf81NpAoPuZF/yDMB1r6IDGMPiU8HDpilysS2swI/MJdEyL03dkqQEbzS4hMT/rcxM+lMQGxvAOrr+RvrJopSvu22fTLNbvYxlweSQbM7L5AYFgKtggS2UrwTtiXYzvkl2H5fPEnnqPUqpMoBZzzhTbRH7kO+riklyam59G2wXjS8Ospr2+HOplM1OrFCC1vEzzEDAhL1A/4J4NQw8jmhJl//bmVLKr0k2u0j+MgUCz8dGl/yssZHLEXsilh7zCgVow+DGW2nrN4vpStHNYC3N53SY9u/jb0G16jcFS0msRHv7yQTrORnF/9INMxt9IDlJtzRTwTWYAuzJ815U3+UwCftNfux5LrJHviAKW1HMF/ZTP5hA0+4m6IMeC1UXPiHgok/h6lHQUQDWhtv2ekFSF+/745iC5ZL32iy+q/jfrdy1otGrqfg+/czIiuZduw+b54Av58i2No2Lb2CZ0NATDUB1dtFYQvQ==","associated_data":"coupon","nonce":"Wf87iOcuq4Cg"}}`
	)
	err := weixin.VerifySign(timestamp, nonce, body, sign)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("succeed")
}

func TestWeixinPay_PayResult(t *testing.T) {
	var data = `{"id":"2d0bd4eb-f862-5a14-8206-8c7cebdd01a6","create_time":"2021-01-27T15:43:36+08:00","resource_type":"encrypt-resource","event_type":"TRANSACTION.SUCCESS","summary":"支付成功","resource":{"original_type":"transaction","algorithm":"AEAD_AES_256_GCM","ciphertext":"2L7Jdtb29djOXSjsv+oBNcmOkhqXoL9ozfDqSwj9yIycGd9UQidZ+thcJVQijjd0UmeOIJTSt+PPFVMlpwVXT96zLJqqdci1P1J1cWdrJAWpalALFythFB+1y9OTiFgrQuAPXMBf8oUAg6kEWNUT2NyKXXBlkzWn9VSQoYVobVVA6Ftexryq+sU6HHIuIuxyqfT9eZ+8e6qOFAkLpt3UszO3IQwmJqXLNcvsYjfTqTfK3zAkzxgxcnvuuYQrDoTP/FUlrFh0aT5NKDKdhaiz7bXOZ1OHISGRCz6e76Sr3yAaUQDunh4JvF4i1ypZ/MfqfFF66w+u6GBSUcHErLyRwNwj+P7iFVQ3sbMSqEZ0tgvHTfrm+RLgGsSLLUkgfrspbiwBsYx0UsNAlTF8YXfYUbj5gNcVM5xGKdIzXc6RaiWKfHC9xRwGGdqPajsaT++MXRLU85zvTlbK+3xLdGM1KE46LxA5JPs+QX0ErKO+x1nmOozvVfJ0gNfnVSKg3dYg7D7llWC6t1N236GK/8ye8h6Yq2gZ0BL7GarSkck38Qb9naYY5JamCqmB/Bdf95TmVUpnA58wStMeRLvKgamt1Kjd3Wb8Tyi0VvwZqvgeNjZ+wufhwkFzJTKOC+hchyVTI8KeVFlO70g0gD0PAWbsRNvTxgosvCmhc+wnsjAxhfhVCZusmFD7W3qOyJYfkjUR0KsqQR/Nyl+v7Ejfb1xLU1SXrgjstJ77cxEYvrdtfCNvSYYo0qa5iKm6MLS26QVqU0zzmt3cN7R6o9jdt563WRbaYw5ZkctJBWLtP29BpQZ2QMBaAxTbLboIvZxt+X3x+RaYNSCIVOvk2HKTTRpj2mmsfD1L7XEx9M9qBdu+x0v7MxI531nW","associated_data":"transaction","nonce":"kCQ1WHMeORAI"}}`

	cipherResult := &CipherData{}
	if err := json.Unmarshal([]byte(data), cipherResult); err != nil {
		t.Fatal(err)
	}

	result, err := weixin.PayResult(&cipherResult.Resource)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}

func TestWeixinPay_RefundResult(t *testing.T) {
	var data = `{"id":"1e8b205f-7462-50b3-ad27-4ff18c8cc199","create_time":"2021-01-27T22:04:32+08:00","resource_type":"encrypt-resource","event_type":"REFUND.SUCCESS","summary":"退款成功","resource":{"original_type":"refund","algorithm":"AEAD_AES_256_GCM","ciphertext":"F2EuImVDX9s/DeWobzHhtgD3DokRZowHDHNEajs5YGEaWXjGM1JQc48tW5Y6Pz1Zn8GwBbMYUhFEH+00p8xMHzioQYJJCvN3cfS0/PU8Ww/qyq989NkrEyhlxbul1AtjmK8Qxb89Jxvd58azNHQLlD0jZPuJenDkrKdagDklHjhMtxq+DgV6RXbbt1fmsNeZ6ZFgLXhA+vlZ6Y1cvQmBorY/tp5nOJ9aECJKE1N2ysLhodu5GzSmWcgGGqbWd8YhWATW+PVE/cYCZpcNe867c4gmUBQN8BziOCJ8SuiVskJpPDCrWt6BVxpp86IorDtq5l4R7DWd5NXwEYWhcAEcg6KY3cTYyHc9PPkHYBvskSwpnK45","associated_data":"refund","nonce":"thqwv6pmSx2V"}}`

	cipherResult := &CipherData{}
	if err := json.Unmarshal([]byte(data), cipherResult); err != nil {
		t.Fatal(err)
	}

	result, err := weixin.RefundResult(&cipherResult.Resource)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}

func TestWeixinPay_MarketingResult(t *testing.T) {
	var data = `{"id":"9c0641c1-a71e-5198-955d-594b74238e8d","create_time":"2021-02-01T16:08:06+08:00","resource_type":"encrypt-resource","event_type":"COUPON.USE","summary":"代金券核销通知","resource":{"original_type":"coupon","algorithm":"AEAD_AES_256_GCM","ciphertext":"6+qemsL3NYZ9KMAe1QCBHfAPdwuIdpIRtbdFJNdapn0oAh3b1Ac70shyMeTCErTS0SZXfoIIXR8mKCaSJx/RnVzbzYGINRLz3J62ERxcnxN/KiqciNJCU2oGcaT/qcdsXC0PNhtDZ4ihGXJZgaxvWgGd9qLfD7QLz0ojK2St3uwpM8+7PyPTRtl6eyfMrC9gHaC+j+VAOOs6uQO7tprAQVdj1UCjKVA9gBLUzRl7856oP6W99rkqrItRD8lL5JhYMeM+rhRRw1GNjWwHmMyR4Gq+InMGxefd+0b/bVAw73oewnR/jaXuP4xeAzp7cx88ZQPfilUfScYq+4nvHoToGUfRCo3Tljgwo4SwQGRx2Ho+tVQBNqIo1LASMKX2l5gdHsR3shxIoh9Y3mtyAuqGFJCNW6aOsND0QHcgdL61LmFDWT38jGzby3RMfMJ5x6Ixei5pf57VtImykcM6nIh2QjKATpuWsS0gN0qBlC3WSHyurdtQEHbusIXd6kkSbHre3dLGDsGL1VMULpeMLMl7xwaPpid1R8mks+1nCvSmm2j8mCH4/lVBOOGfx6eQiIn+Qj+uJWhKx8oZYfLg8cxEUbJjAAVpLKmcI952HnXkYwx+QJLznzUVB3SmWfOCrOOSUTEf0Bk+cB0SZ3p8iEKTg6ECQnIIStt9BMQca/WqdtxKkaNfKx5ZRE5NBNKO8rfeRsrWsEey3N9MqQxzToF/xQdBJH5A2uwLKtou/dgicOnrWBY3XuidI0cShmkm26zphNx7vz333Yft9WV8lw==","associated_data":"coupon","nonce":"cJSYOsVyoek2"}}`

	cipherResult := &CipherData{}
	if err := json.Unmarshal([]byte(data), cipherResult); err != nil {
		t.Fatal(err)
	}

	result, err := weixin.MarketingResult(&cipherResult.Resource)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}

func TestWeixinPay_MarketingStocks(t *testing.T) {
	result, err := weixin.MarketingStocks(15441178)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}

func TestWeixinPay_ProfitsharingReceiversAdd(t *testing.T) {
	result, err := weixin.ProfitsharingReceiversAdd(&ProfitsharingReceiversAddReq{
		ReceiveOpenId: "oasi95rPit953LHRYfaifGnTuqgs",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}

func TestWeixinPay_ProfitsharingOrders(t *testing.T) {
	result, err := weixin.ProfitsharingOrders(&ProfitsharingOrdersReq{
		TransactionId:       "12854923295488778253",
		WeixinTransactionId: "4200001377202204289129764075",
		ReceiveOpenId:       "oasi95rPit953LHRYfaifGnTuqgs",
		ReceiveAmount:       0.30,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}
