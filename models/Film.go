package models

import "time"

type Film struct {
	ID                 uint
	Title              string    `gorm:"not null;unique_index"`
	Genre              string    `gorm:"not null"`
	ReleaseDate        time.Time `gorm:"not null"`
	Synopsis           string    `gorm:"not null"`
	Director           string    `gorm:"not null"`
	Duration           int       `gorm:"not null"`
	FilmaffinityRating float32
	CreatedAt          time.Time `gorm:"not null"`
	UpdatedAt          time.Time `gorm:"not null"`
	CreateUserID       uint      `gorm:"not null"`
	CreateUser         User      `gorm:"foreignkey:CreateUserID"`
}
