package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

var url = "sql6425053:jCb78RUR66@tcp(sql6.freesqldatabase.com:3306)/sql6425053?parseTime=true"

var err error

func DataMigration() {
	Database, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	Database.AutoMigrate(&Books{}, &Authors{}, &Users{}, &AutherBooks{})
}
