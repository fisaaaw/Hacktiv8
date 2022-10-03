package model

// import "gorm.io/gorm"

type Person struct {
	// gorm.Model
	// Post FirstName
	// @Column(type:string)
	FirstName string `gorm:"first_name"`
	// Post LasttName
	// @Column(type:string)
	LastName string `gorm:"last_name"`
}