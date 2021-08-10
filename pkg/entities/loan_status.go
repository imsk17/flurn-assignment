package entities

import "database/sql/driver"

type Status string

const (
	New       Status = "new"
	Approved  Status = "approved"
	Rejected  Status = "rejected"
	Cancelled Status = "cancelled"
)

func (e *Status) Scan(value interface{}) error {
	*e = Status(value.(string))
	return nil
}

func (e Status) Value() (driver.Value, error) {
	return string(e), nil
}

func (Status) GormDataType() string {
	return "status"
}