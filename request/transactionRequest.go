package request

type TransactionRequest struct {
	Transaction_id string       `json:"transaction_id"`
	Customer_info  CustomerInfo `json:"customer_info"`
	Seller_info    SellerInfo   `json:"seller_info"`
}

type CustomerInfo struct {
	Customer         string `json:"customer_name"`
	Customer_account string `json:"customer_account"`
	Ammount          int    `json:"ammount"`
}

type SellerInfo struct {
	Seller         string `json:"seller_name"`
	Seller_account string `json:"seller_account"`
}
