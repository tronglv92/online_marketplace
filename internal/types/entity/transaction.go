package entity

type Transaction struct {
	IdModel
	ItemID  int32  `gorm:"column:item_id;not null"`
	BuyerID uint   `gorm:"column:buyer_id;not null"`
	Status  string `gorm:"column:status;not null;type:varchar(10)"`
}
