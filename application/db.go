package application

import (
	"context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	ID        uint `gorm:"primary_key;auto_increment;not_null"`
	FirstName string
	LastName  string
}

func NewDB() *DB {
	db, _ := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})

	err := db.AutoMigrate(&User{})
	if err == nil {
		return &DB{db}
	}

	panic(err)
}

func (d *DB) FindByID(ID string, ctx context.Context) *User {
	var user User
	d.db.First(&user, ID)

	return &user
}

func (d *DB) CreateNewUser(fistName, lastName string) *User {
	user := User{FirstName: fistName, LastName: lastName}

	d.db.Save(&user)

	return &user
}
