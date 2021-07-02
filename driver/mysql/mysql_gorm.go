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
	db.LogMode(c.Debug)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.BlockGlobalUpdate(true)
	db.DB().SetMaxIdleConns(c.MaxIdleConn)
	db.DB().SetMaxOpenConns(c.MaxOpenConn)
	return
}
