package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var(
	DBConn *gorm.DB
)