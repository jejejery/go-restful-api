package web

import "time"

type EmployeeUpdateRequest struct {
	EmployeeID string 		`validate:"required" json:"employee_id"`
	Name       string 		`json:"name"`
	Role       string 		`json:"role"` // e.g., Cashier, Manager
	Email      string 		`json:"email"`
	Phone      string 		`json:"phone"`
	DateHired  time.Time	`json:"date_hired"`
}