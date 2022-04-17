package entity

type User struct {
	Model
	Name  string `gorm:"size:255"`
	Age   int
	Email string `gorm:"size:255"`
}
