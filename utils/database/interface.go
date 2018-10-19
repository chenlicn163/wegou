package database

import "github.com/jinzhu/gorm"

type Database interface {
	Open() *gorm.DB
}
