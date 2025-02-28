package web

import "time"

type EmployeeCreateRequest struct {
	EmployeeID string 		`validate:"required" json:"employee_id"`
	Name       string 		`validate:"required" json:"name"`
	Role       string 		`validate:"required" json:"role"` // e.g., Cashier, Manager
	Email      string 		`validate:"required" json:"email"`
	Phone      string 		`validate:"required" json:"phone"`
	DateHired  time.Time	`validate:"required" json:"date_hired"`
}