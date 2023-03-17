package models

type Role struct {
	ID   int    `gorm:"->"`
	name string `gorm:"->"`
}
