package entities

import (
	"time"
)

type LoanRequest struct {
	ID           uint         `gorm:"primarykey"`
	CreatedAt    time.Time    `json:"-"`
	UpdatedAt    time.Time    `json:"-"` // Ignore this field in marshalling of json
	CustomerName string       `json:"customerName"`
	PhoneNo      string       `json:"phoneNo"`
	Email        string       `json:"Email"`
	LoanAmount   int32        `json:"loanAmount"`
	Status       Status       `json:"status" sql:"type:status"`
	CreditScore  int          `json:"creditScore"`
}
