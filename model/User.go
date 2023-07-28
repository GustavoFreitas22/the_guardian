package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
	UserKey  string

	Salt []byte
}
