package entity

type Product struct {
	IdModel
	Name        string  `gorm:"column:name;not null;type:varchar(50)"`
	Description string  `gorm:"column:name;type:varchar(250)"`
	Price       float64 `gorm:"column:price"`
	Seller      string  `gorm:"column:seller;type:varchar(50)"`
}

func (*Product) TableName() string {
	return "products"
}
