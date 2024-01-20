package tables

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fname    string
	Lname    string
	Email    string
	Phone    string
	Password string
}

type Projects struct {
	gorm.Model
	Name        string
	Designer_id uint
	Designer    Designer `gorm:"foreignKey:Designer_id"`
}

type Designer struct {
	gorm.Model
	Name  string
	Email string
}
