package service

import (
	"log"
	"os"
	"github.com/gdcolmenares/GinTutorial/entity"
	"github.com/gdcolmenares/GinTutorial/request"
	"github.com/gdcolmenares/GinTutorial/response"
)

type TransactionService interface {
	MakeTransaction(request.TransactionRequest) response.TransactionResponse
}

type transactionService struct {
	transaction entity.Transaction
}

func New() TransactionService {
	return &transactionService{}
}

func (service *transactionService) MakeTransaction(request request.TransactionRequest) response.TransactionResponse {

	file, err := openLogFile("./mylog.log")
    if err != nil {
        log.Fatal(err)
    }
	log.SetOutput(file)
    log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

	log.Println("log file created")

	log.Println("-------Inicio transaccion-------")
	log.Println("Request transaccion: ")

	log.Println("transaction id:", request.Transaction_id,
		"Customer name:", request.Customer_info.Customer,
		"Customer account:", request.Customer_info.Customer_account,
		"Seller name:", request.Seller_info.Seller,
		"Seller account:", request.Seller_info.Seller_account,
		"Transaction ammount", request.Customer_info.Ammount)

	var response response.TransactionResponse

	response.Transaction_id = request.Transaction_id
	response.Customer_info.Customer_name = request.Customer_info.Customer
	response.Customer_info.Customer_account = request.Customer_info.Customer_account
	response.Customer_info.Debit_ammount = request.Customer_info.Ammount
	response.Seller_info.Seller_name = request.Seller_info.Seller
	response.Seller_info.Seller_account = request.Seller_info.Seller_account
	response.Seller_info.Credit_ammount = request.Customer_info.Ammount

	log.Println("Response transaccion: ")
	log.Println("Transaction id", response.Transaction_id,
		"Customer name:", response.Customer_info.Customer_name,
		"Customer account:", response.Customer_info.Customer_account,
		"Debit Ammount:", response.Customer_info.Debit_ammount,
		"Seller Name:", response.Seller_info.Seller_name,
		"Seller Account:", response.Seller_info.Seller_account,
		"Credit Ammount:", response.Seller_info.Credit_ammount)

	log.Println("-------Fin Transaccion-------")

	return response
}

func openLogFile(path string) (*os.File, error) {
    logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        return nil, err
    }
    return logFile, nil
}
