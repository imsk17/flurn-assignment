package entities

import (
	"gorm.io/gorm"
)

type LoanRequest struct {
	gorm.Model
	CustomerName string `json:"customerName"`
	PhoneNo      string `json:"phoneNo"`
	Email        string `json:"Email"`
	LoanAmount   int32  `json:"loanAmount"`
	Status       Status `json:"status" sql:"type:status"`
	CreditScore  int    `json:"creditScore"`
}
