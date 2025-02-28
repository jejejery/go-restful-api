package domain

type Product struct {
	ProductID   int  `json:"product_id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	StockQty    int     `json:"stock_qty"`
	Category    string  `json:"category"`
}
