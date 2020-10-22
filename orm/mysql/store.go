package mysql

import (
	"github.com/astaxie/beego/client/orm"
	"sync"
)

var instances = sync.Map{}

func Range(fn func(name string, db orm.Ormer) bool) {
	instances.Range(func(key, val interface{}) bool {
		return fn(key.(string), val.(orm.Ormer))
	})
}