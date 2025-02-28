package web

type CustomerCreateRequest struct {
	CustomerID string `validate:"required" json:"customer_id"` 
	Name       string `validate:"required,min=1,max=100" json:"name"`
	Email      string `validate:"required" json:"email"`
	Phone      string `validate:"required" json:"phone"`
	Address    string `validate:"required" json:"address"`
}
