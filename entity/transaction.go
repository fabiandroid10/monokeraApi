package entity

//Se crea la entidad para el caso futuro de querer almacenar informacion de las trx en la base de datos

type Transaction struct {
	Transaction_id   string `json:"transaction_id"`
	Customer         string `json:"customer"`
	Customer_account string `json:"customer_account"`
	Seller           string `json:"seller"`
	Seller_account   string `json:"seller_account"`
	Ammount          int    `json:"ammount"`
}
