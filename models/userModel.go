package models

type User struct{
	ID    uint `gorm:"primaryKey;default:auto_random()"`
	Name string
	Age int
	Phone int
}