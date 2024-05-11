package models

type ReqAddAccBank struct {
	AccName   string `json:"account_name"`
	BankName  string `json:"bank_name"`
	AccNumber string `json:"account_number"`
	Saldo     int64  `json:"saldo"`
}

type RespDetailAccBank struct {
	ID        string `json:"id"`
	AccName   string `json:"account_name"`
	BankName  string `json:"bank_name"`
	AccNumber string `json:"account_number"`
	Saldo     int64  `json:"saldo"`
}

type RespListAccBank struct {
	ID        string `json:"id"`
	BankName  string `json:"bank_name"`
	AccNumber string `json:"account_number"`
}

type ReqUpdateAccBalance struct {
	AccNumber string `json:"account_number"`
	Amount    int64  `json:"amount"`
}
