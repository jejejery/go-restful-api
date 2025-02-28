package web

type ProductCreateRequest struct {
	Name        string  `validate:"required,max=200,min=1" json:"name"`
	Description string  `validate:"required,max=200,min=1" json:"description"`
	Price       float64 `validate:"required" json:"price"`
	StockQty    int     `validate:"required" json:"stock_qty"`
	Category    string  `validate:"required" json:"category"`
}
