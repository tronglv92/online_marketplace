package entity

func RegisterMigrate() []any {
	return []any{
		&User{},
		&Product{},
		&Transaction{},
	}
}
