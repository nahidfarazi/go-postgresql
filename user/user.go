package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `json:"name"`
	Age       string `json:"age"`
	AddressID uint
	Address   *Address `gorm:"foreignKey:AddressID"`
	ContactID uint
	Contact   *Contact `gorm:"foreignKey:ContactID"`
}
type Address struct {
	gorm.Model
	Country string `json:"country"`
	City    string `json:"city"`
	State   string `json:"state"`
}
type Contact struct {
	gorm.Model
	Number string `json:"number"`
	Email  string `json:"email"`
}
