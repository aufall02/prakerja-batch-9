package models

// import "gorm.io/gorm"

type User struct {
	Nama     string `json:"nama" form:"nama" validate:"required" gorm:"primaryKey;unique;size:100"`
	Email    string `json:"email" form:"email" validate:"required" gorm:"size:100;not null"`
	Password string `json:"password" form:"password" validate:"required" gorm:"size:100;not null"`
}
