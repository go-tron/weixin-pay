package weixinPay

import (
	"crypto"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	baseError "github.com/go-tron/base-error"
	"github.com/go-tron/crypto/encoding"
	"github.com/go-tron/crypto/rsaUtil"
	localTime "github.com/go-tron/local-time"
	"github.com/go-tron/logger"
	"github.com/go-tron/random"
	"github.com/tidwall/gjson"
	"net/url"
	"strconv"
)

var (
	ErrorSign          = baseError.System("53001", "微信支付签名失败")
	ErrorMethod        = baseError.SystemFactory("53002", "微信支付方式无效:{}")
	ErrorAuthorize     = baseError.System("53004", "微信授权失败")
	ErrorRequest       = baseError.System("53005", "微信服务连接失败")
	ErrorUnmarshalBody = baseError.System("53006", "微信消息解析失败")
	ErrorCode          = baseError.SystemFactory("53010")
)

func New(c *Config) *WeixinPay {

	if c == nil {
		panic("config 必须设置")
	}
	if c.AppId == "" {
		panic("AppId 必须设置")
	}
	if c.MchId == "" {
		panic("MchId 必须设置")
	}
	if c.PrivateKeyPem == "" {
		panic("PrivateKeyPem 必须设置")
	}
	if c.CertSN == "" {
		panic("CertSN 必须设置")
	}
	if c.APIV3Key == "" {
		panic("APIV3Key 必须设置")
	}
	if c.Logger == nil {
		panic("Logger 必须设置")
	}

	privateKey, err := rsaUtil.GetPrivateKeyPem([]byte(c.PrivateKeyPem))
	if err != nil {
		panic(err)
	}
	c.PrivateKey = privateKey

	publicKey, err := rsaUtil.GetPublicKeyFromCertificate([]byte(WeixinPublicCert))
	if err != nil {
		panic(err)
	}

	return &WeixinPay{
		Config:    c,
		PublicKey: publicKey,
	}
}

type WeixinPay struct {
	*Config
	PublicKey *rsa.PublicKey
}

type Config struct {
	AppId         string
	MchId         string
	CertSN        string
	PrivateKeyPem string
	PrivateKey    *rsa.PrivateKey
	APIV3Key      string
	NotifyUrl     string
	Logger        logger.Logger
}

func (wxp *WeixinPay) GetToken(method string, path string, content string) (string, error) {
	var nonceStr = random.Letter(16)
	var timestamp = strconv.FormatInt(localTime.Now().Unix(), 10)
	var message = method + "\n" + path + "\n" + timestamp + "\n" + nonceStr + "\n" + content + "\n"
	signature, err := wxp.Sign(message)
	if err != nil {
		return "", ErrorSign
	}
	return "WECHATPAY2-SHA256-RSA2048 mchid=\"" + wxp.MchId + "\"," + "nonce_str=\"" + nonceStr + "\"," + "timestamp=\"" + timestamp + "\"," + "serial_no=\"" + wxp.CertSN + "\"," + "signature=\"" + signature + "\"", nil
}

func (wxp *WeixinPay) Sign(obj string) (string, error) {
	sign, err := rsaUtil.Sign(obj, wxp.PrivateKey, crypto.SHA256, &encoding.Base64{})
	if err != nil {
		return "", ErrorSign
	}
	return sign, nil
}

func (wxp *WeixinPay) Execute(name string, data map[string]interface{}, res interface{}) (interface{}, error) {
	sdkConfig := SDKConfigMap[name]
	if sdkConfig == nil {
		return nil, ErrorMethod(name)
	}

	request, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	wxp.Logger.Info(string(request),
		wxp.Logger.Field("orderId", data["out_trade_no"]),
		wxp.Logger.Field("name", name),
		wxp.Logger.Field("type", "request"),
	)

	return wxp.Request(name, data, res, sdkConfig)
}

func (wxp *WeixinPay) Request(name string, data map[string]interface{}, res interface{}, sdkConfig *SDKConfig) (result interface{}, err error) {

	var (
		response = ""
	)
	defer func() {
		wxp.Logger.Info(response,
			wxp.Logger.Field("orderId", data["out_trade_no"]),
			wxp.Logger.Field("name", name),
			wxp.Logger.Field("type", "response"),
			wxp.Logger.Field("error", err))
	}()

	uri := ReplaceUrl(sdkConfig.Url, data)
	content := ""
	if sdkConfig.Method == "GET" {
		var tmp = ""
		for k, v := range data {
			tmp += k + "=" + fmt.Sprintf("%v", v)
		}
		if tmp != "" {
			uri += "?" + tmp
		}
	} else {
		body, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		content = string(body)
	}

	var path = ""
	URL, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	path = URL.Path
	if URL.RawQuery != "" {
		path += "?" + URL.RawQuery
	}

	token, err := wxp.GetToken(sdkConfig.Method, path, content)
	if err != nil {
		return nil, ErrorAuthorize
	}

	request := resty.New().R().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", token)

	if sdkConfig.Method == "POST" {
		request.SetBody(data)
	}
	resp, err := request.
		Execute(sdkConfig.Method, uri)
	if err != nil {
		return nil, ErrorRequest
	}

	response = string(resp.Body())
	code := gjson.Get(response, "code").String()
	message := gjson.Get(response, "message").String()

	if code != "" {
		if message == "" {
			message = name
		}
		return nil, ErrorCode(fmt.Sprintf("(%s)%s", message, code))
	}

	if err := json.Unmarshal(resp.Body(), res); err != nil {
		return nil, ErrorUnmarshalBody
	}

	return res, nil
}

func (wxp *WeixinPay) VerifySign(timestamp, nonce, body, sign string) error {
	var message = timestamp + "\n" + nonce + "\n" + body + "\n"
	return rsaUtil.Verify(message, sign, wxp.PublicKey, crypto.SHA256, &encoding.Base64{})
}

func (wxp *WeixinPay) PayResult(params *Resource) (*PayResult, error) {
	data, err := AESGCMDecrypter(wxp.APIV3Key, params.Nonce, params.Ciphertext, params.AssociatedData)
	if err != nil {
		return nil, err
	}

	result := &PayResult{}
	if err := json.Unmarshal(data, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (wxp *WeixinPay) RefundResult(params *Resource) (*RefundResult, error) {
	data, err := AESGCMDecrypter(wxp.APIV3Key, params.Nonce, params.Ciphertext, params.AssociatedData)
	if err != nil {
		return nil, err
	}
	result := &RefundResult{}
	if err := json.Unmarshal(data, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (wxp *WeixinPay) MarketingResult(params *Resource) (*MarketingResult, error) {
	data, err := AESGCMDecrypter(wxp.APIV3Key, params.Nonce, params.Ciphertext, params.AssociatedData)
	if err != nil {
		return nil, err
	}
	result := &MarketingResult{}
	if err := json.Unmarshal(data, result); err != nil {
		return nil, err
	}
	return result, nil
}
