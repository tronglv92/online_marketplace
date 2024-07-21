package entity

type User struct {
	IdModel
	Email     string `gorm:"column:email;uniqueIndex;type:varchar(50)"`
	Password  string `gorm:"column:password;not null;type:varchar(50)"`
	Salt      string `gorm:"column:salt;not null;type:varchar(50)"`
	LastName  string `gorm:"column:last_name;type:varchar(50)"`
	FirstName string `gorm:"column:first_name;type:varchar(50)"`
	Phone     string `gorm:"column:phone;not null;type:varchar(50)"`
}

func (User) TableName() string {
	return "users"
}
