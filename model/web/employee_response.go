package web

type EmployeeResponse struct {
	EmployeeID string `json:"employee_id"`
	Name       string `json:"name"`
	Role       string `json:"role"` // e.g., Cashier, Manager
}