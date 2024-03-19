package database

import (
	"database/sql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var SqlDB *sql.DB
