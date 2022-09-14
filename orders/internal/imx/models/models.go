package models

// Struct for Order is created based on the response from the IMX API here:
// https://docs.x.immutable.com/reference/#/operations/getOrder
type Order struct {
	AmountSold          string `json:"amount_sold"`
	Buy                 Buy
	ExpirationTimestamp string `json:"expiration_timestamp"`
	Fees                []Fee
	OrderId             int `json:"order_id"`
	Sell                Sell
	Status              string `json:"status"`
	Timestamp           string `json:"timestamp"`
	UpdatedTimestamp    string `json:"updated_timestamp"`
	User                string `json:"user"`
}

type Buy struct {
	Data OrderData
	Type string `json:"type"`
}

type Sell struct {
	Data OrderData
	Type string `json:"type"`
}

type Fee struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
	Token   Token
	Type    string `json:"type"`
}

type Token struct {
	Data TokenData
	Type string `json:"type"`
}

type TokenData struct {
	ContractAddress string `json:"contract_address"`
	Decimals        int64  `json:"decimals"`
}

type OrderData struct {
	Decimals         int64  `json:"decimals"`
	Id               string `json:"id"`
	Properties       Properties
	Quantity         string `json:"quantity"`
	QuantityWithFees string `json:"quantity_with_fees"`
	Symbol           string `json:"symbol"`
	TokenAddress     string `json:"token_address"`
	TokenId          string `json:"token_id"`
}

type Properties struct {
	Collection Collection
	ImageUrl   string `json:"image_url"`
	Name       string `json:"name"`
}

type Collection struct {
	IconUrl string `json:"icon_url"`
	Name    string `json:"name"`
}
