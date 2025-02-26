package domain

type Category struct {
	Id   int    `gorm:"primary_key; column:id"`
	Name string `gorm:"column:name"`
}
