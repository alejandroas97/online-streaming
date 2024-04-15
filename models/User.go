package models

type User struct {
	ID        uint
	Username  string `gorm:"not null"`
	Password  string `gorm:"not null"`
	IsEnabled bool   `gorm:"default:true"`
}
