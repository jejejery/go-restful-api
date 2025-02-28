package web

type ProductUpdateRequest struct {
	Id          int     `validate:"required" json:"product_id"`
	Name        string  `validate:"required,max=200,min=1" json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	StockQty    int     `json:"stock_qty"`
	Category    string  `json:"category"`
}