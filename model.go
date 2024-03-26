package weixinPay

import "time"

type CipherData struct {
	ID           string    `json:"id"`
	CreateTime   time.Time `json:"create_time"`
	ResourceType string    `json:"resource_type"`
	EventType    string    `json:"event_type"`
	Summary      string    `json:"summary"`
	Resource     Resource  `json:"resource"`
}

type Resource struct {
	OriginalType   string `json:"original_type"`
	Algorithm      string `json:"algorithm"`
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data"`
	Nonce          string `json:"nonce"`
}

type PayResult struct {
	Amount          PayAmount         `json:"amount"`
	Appid           string            `json:"appid"`
	Attach          string            `json:"attach"`
	BankType        string            `json:"bank_type"`
	Mchid           string            `json:"mchid"`
	OutTradeNo      string            `json:"out_trade_no"`
	Payer           Payer             `json:"payer"`
	PromotionDetail []PromotionDetail `json:"promotion_detail"`
	SuccessTime     time.Time         `json:"success_time"`
	TradeState      string            `json:"trade_state"`
	TradeStateDesc  string            `json:"trade_state_desc"`
	TradeType       string            `json:"trade_type"`
	TransactionID   string            `json:"transaction_id"`
}
type PayAmount struct {
	Currency      string `json:"currency"`
	PayerCurrency string `json:"payer_currency"`
	PayerTotal    int    `json:"payer_total"`
	Total         int    `json:"total"`
}
type Payer struct {
	Openid string `json:"openid"`
}

type PromotionDetail struct {
	Amount              int           `json:"amount"`
	CouponID            string        `json:"coupon_id"`
	Currency            string        `json:"currency"`
	GoodsDetail         []interface{} `json:"goods_detail"`
	MerchantContribute  int           `json:"merchant_contribute"`
	Name                string        `json:"name"`
	OtherContribute     int           `json:"other_contribute"`
	Scope               string        `json:"scope"`
	StockID             string        `json:"stock_id"`
	Type                string        `json:"type"`
	WechatpayContribute int           `json:"wechatpay_contribute"`
}

type RefundResult struct {
	Amount              RefundAmount `json:"amount"`
	Mchid               string       `json:"mchid"`
	OutRefundNo         string       `json:"out_refund_no"`
	OutTradeNo          string       `json:"out_trade_no"`
	RefundID            string       `json:"refund_id"`
	Status              string       `json:"status"`
	RefundStatus        string       `json:"refund_status"`
	SuccessTime         time.Time    `json:"success_time"`
	TransactionID       string       `json:"transaction_id"`
	UserReceivedAccount string       `json:"user_received_account"`
}

type RefundAmount struct {
	Total          int `json:"total"`
	Refund         int `json:"refund"`
	PayerTotal     int `json:"payer_total"`
	PayerRefund    int `json:"payer_refund"`
	DiscountRefund int `json:"discount_refund"`
}

type CertificateData struct {
	Data []struct {
		EffectiveTime      time.Time `json:"effective_time"`
		EncryptCertificate struct {
			Algorithm      string `json:"algorithm"`
			AssociatedData string `json:"associated_data"`
			Ciphertext     string `json:"ciphertext"`
			Nonce          string `json:"nonce"`
		} `json:"encrypt_certificate"`
		ExpireTime time.Time `json:"expire_time"`
		SerialNo   string    `json:"serial_no"`
	} `json:"data"`
}

type MarketingResult struct {
	StockCreatorMchid       string                  `json:"stock_creator_mchid"`
	StockID                 string                  `json:"stock_id"`
	CouponID                string                  `json:"coupon_id"`
	CouponName              string                  `json:"coupon_name"`
	Status                  string                  `json:"status"`
	Description             string                  `json:"description"`
	CreateTime              time.Time               `json:"create_time"`
	CouponType              string                  `json:"coupon_type"`
	NoCash                  bool                    `json:"no_cash"`
	AvailableBeginTime      time.Time               `json:"available_begin_time"`
	AvailableEndTime        time.Time               `json:"available_end_time"`
	Singleitem              bool                    `json:"singleitem"`
	NormalCouponInformation NormalCouponInformation `json:"normal_coupon_information"`
	ConsumeInformation      ConsumeInformation      `json:"consume_information"`
}
type NormalCouponInformation struct {
	CouponAmount       int `json:"coupon_amount"`
	TransactionMinimum int `json:"transaction_minimum"`
}
type ConsumeInformation struct {
	ConsumeTime   time.Time `json:"consume_time"`
	ConsumeMchid  string    `json:"consume_mchid"`
	TransactionID string    `json:"transaction_id"`
}
