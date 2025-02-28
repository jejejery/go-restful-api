package domain

import "time"

type Employee struct {
	EmployeeID string `json:"employee_id"`
	Name       string `json:"name"`
	Role       string `json:"role"` // e.g., Cashier, Manager
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	DateHired  time.Time `json:"date_hired"`
}