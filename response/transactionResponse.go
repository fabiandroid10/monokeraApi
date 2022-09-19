package response

type TransactionResponse struct {
	Transaction_id string       `json:"transaction_id"`
	Customer_info  CustomerInfo `json:"customer_info"`
	Seller_info    SellerInfo   `json:"seller_info"`
}

type CustomerInfo struct {
	Customer_name    string `json:"customer_name"`
	Customer_account string `json:"customer_account"`
	Debit_ammount    int    `json:"debit_ammount"`
}

type SellerInfo struct {
	Seller_name    string `json:"seller_name"`
	Seller_account string `json:"seller_account"`
	Credit_ammount int    `json:"credit_ammount"`
}
