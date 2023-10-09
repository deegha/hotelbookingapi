package models

type User struct {
	ID uint`json:"id"`
	Name string`json:"name"`
	Email string`json:"email" gorm:"unique"`
	Password  []byte`json:"-"`
	Type string `json:"Type"`
	CreatedAt string`json:"createdAt"`
	UpdatedAt string`json:"updatedAt"`
}
