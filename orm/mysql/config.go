package mysql

import (
	"github.com/astaxie/beego/client/orm"
	"github.com/beego/invoker/store"
)

const module = "orm.mysql"

type Config struct {
	// common config
	Debug          bool `ini:"debug"`
	beegoUniqueKey string
	err            error
	// db config
	AliasName    string `ini:"aliasName"`
	MaxIdleConns int    `ini:"maxIdleConns"`
	MaxOpenConns int    `ini:"maxOpenConns"`
	Dsn          string `ini:"dsn"`
}

func DefaultConfig() *Config {
	return &Config{
		AliasName:    "default",
		MaxIdleConns: 10,
		MaxOpenConns: 50,
	}
}

func Invoker(key string) *Config {
	var dc = DefaultConfig()
	dc.beegoUniqueKey = key
	dc.err = store.Config().Unmarshaler(key, &dc)
	return dc
}

func (c *Config) Build() (o orm.Ormer, err error) {
	if c.err != nil {
		return
	}
	err = orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		return
	}
	err = orm.RegisterDataBase(c.AliasName,
		"mysql",
		c.Dsn,
		orm.MaxIdleConnections(c.MaxIdleConns),
		orm.MaxOpenConnections(c.MaxOpenConns),
	)
	if err != nil {
		return
	}
	o = orm.NewOrmUsingDB(c.AliasName)
	instances.Store(c.beegoUniqueKey, o)
	return
}

func init() {
	store.RegisterModule(module)
}
