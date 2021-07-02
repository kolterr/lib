package redis

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

type Config struct {
	MaxIdle        int
	MaxActive      int
	IdleTimeout    time.Duration
	Wait           bool
	DSN            string
	Database       int
	Password       string
	ReadTimeout    time.Duration
	ConnectTimeout time.Duration
	WriteTimeout   time.Duration
	KeepAlive      time.Duration
}

type Pool struct {
	*redis.Pool
}

func NewPoolBak(c *Config) (p *Pool) {
	pool := &redis.Pool{
		MaxIdle:     c.MaxIdle,
		MaxActive:   c.MaxActive,
		IdleTimeout: c.IdleTimeout,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", c.DSN,
				redis.DialPassword(c.Password),
				redis.DialDatabase(c.Database),
				redis.DialConnectTimeout(c.ConnectTimeout),
				redis.DialReadTimeout(c.ReadTimeout),
				redis.DialWriteTimeout(c.WriteTimeout),
				redis.DialKeepAlive(c.KeepAlive),
			)
			if err != nil {
				log.Fatalln(err)
				return nil, err
			}
			return con, nil
		},
		TestOnBorrow: func(con redis.Conn, t time.Time) error {
			_, err := con.Do("PING")
			return err
		},
	}
	return &Pool{pool}
}

func NewPool(c *Config) (p *Pool) {
	pool := &redis.Pool{
		MaxIdle:     c.MaxIdle,
		MaxActive:   c.MaxActive,
		IdleTimeout: c.IdleTimeout,
		Wait:        true,
		Dial:        func() (redis.Conn, error) { return Dial(c) },
		TestOnBorrow: func(con redis.Conn, t time.Time) error {
			_, err := con.Do("PING")
			return err
		},
	}
	return &Pool{pool}
}

func Dial(c *Config) (redis.Conn, error) {
	dialFunc := func() (redis.Conn, error) {
		return redis.Dial("tcp", c.DSN,
			redis.DialPassword(c.Password),
			redis.DialDatabase(c.Database),
			redis.DialConnectTimeout(c.ConnectTimeout),
			redis.DialReadTimeout(c.ReadTimeout),
			redis.DialWriteTimeout(c.WriteTimeout),
			redis.DialKeepAlive(c.KeepAlive),
		)
	}
	conn, err := dialFunc()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return conn, err
}
