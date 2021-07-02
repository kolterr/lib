package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
	DSN         string
	MaxIdleConn int
	MaxOpenConn int
	Debug       bool
}

func NewGorm(c *Config) (db *gorm.DB) {
	var err error
	db, err = gorm.Open("mysql", c.DSN)
	if c.Debug == true {
		db.LogMode(true)
	}
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.BlockGlobalUpdate(true)
	db.DB().SetMaxIdleConns(c.MaxIdleConns)
	db.DB().SetMaxOpenConns(c.MaxOpenConns)
	return
}
