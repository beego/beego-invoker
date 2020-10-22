package mysql

import (
	"github.com/astaxie/beego/client/orm"
	"sync"
)

var DefaultInstances = &instancesStore{}

type instancesStore struct {
	store sync.Map
}

// Store
func (i *instancesStore) Store(key string, value orm.Ormer) {
	i.store.Store(key, value)
}

// Range
func (i *instancesStore) Range(fn func(name string, db orm.Ormer) bool) {
	i.store.Range(func(key, val interface{}) bool {
		return fn(key.(string), val.(orm.Ormer))
	})
}

// Clean
func (i *instancesStore) Clean() {
	i.Range(func(name string, db orm.Ormer) bool {
		//db.Close()
		return true
	})
}

// Stats
func (i *instancesStore) Stats() (stats map[string]interface{}) {
	stats = make(map[string]interface{})
	i.store.Range(func(key, val interface{}) bool {
		name := key.(string)
		db := val.(orm.Ormer)
		stats[name] = db.DBStats()
		return true
	})
	return
}
