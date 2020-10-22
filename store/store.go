package store

import (
	"github.com/astaxie/beego/core/config"
	"sync"
)

var (
	cfg   config.Configer
	store sync.Map // register module
)

func SetConfig(c config.Configer) {
	cfg = c
}

func Config() config.Configer {
	return cfg
}

func RegisterModule(module string) {
	store.Store(module, true)
}

func Range(fn func(name string, flag bool) bool) {
	store.Range(func(key, val interface{}) bool {
		return fn(key.(string), val.(bool))
	})
}
