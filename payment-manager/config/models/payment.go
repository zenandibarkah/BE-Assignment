package models

type ReqTransaction struct {
	SourceBank   string `json:"source_bank"`
	SourceAccNum string `json:"source_account_number"`
	DestBank     string `json:"destination_bank"`
	DestAccNum   string `json:"destination_account_number"`
	Amount       int64  `json:"amount"`
}

type RespListTrans struct {
	ID         string `json:"id"`
	SourceBank string `json:"source_bank"`
	DestBank   string `json:"destination_bank"`
	TransType  string `json:"trans_type"`
}

type RespDetailTrans struct {
	ID           string `json:"id"`
	SourceBank   string `json:"source_bank"`
	SourceAccNum string `json:"source_account_number"`
	DestBank     string `json:"destination_bank"`
	DestAccNum   string `json:"destination_account_number"`
	Amount       int64  `json:"amount"`
	TransType    string `json:"trans_type"`
	TransDate    string `json:"trans_date"`
}

type ReqUpdateAccBalance struct {
	AccNumber string `json:"account_number"`
	Amount    int64  `json:"amount"`
}
