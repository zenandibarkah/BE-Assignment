package models

type Response struct {
	StatusCode int         `json:"statusCode"`
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	DisplayMsg string      `json:"displayMessage"`
	Response   interface{} `json:"response"`
}
