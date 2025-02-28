package domain

type Customer struct {
	CustomerID string `json:"customer_id" gorm:"primaryKey"` 
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	LoyaltyPts int    `json:"loyalty_points"`
}
