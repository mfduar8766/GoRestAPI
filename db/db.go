package db

import (
	"github.com/jinzhu/gorm"
)

var (
	// GormInstance - Used as DB query builder
	GormInstance *gorm.DB
)
