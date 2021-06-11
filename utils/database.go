package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var _db *gorm.DB

func init() {

	var err error
	dsn := "intelligent:x&$Y4M0Gzr&V^@tcp(finlogix-lab.chyx61txa0gb.eu-west-2.rds.amazonaws.com:3306)/schema_dev?charset=utf8mb4&parseTime=True&loc=Local"

	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	},
	})
	if err != nil {
		panic("connect error, error=" + err.Error())
	}

	sqlDB, _ := _db.DB()

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
}

func GetDB() *gorm.DB {
	return _db
}
