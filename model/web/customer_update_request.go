package web

// CustomerUpdateRequest struct
type CustomerUpdateRequest struct {
	CustomerID string `validate:"required" json:"customer_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	LoyaltyPts int    `json:"loyalty_points"`
}