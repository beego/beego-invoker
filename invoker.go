package invoker

import (
	"github.com/astaxie/beego/core/config"
	"github.com/beego/invoker/store"
)

func Init(cfg config.Configer)  {
	store.SetConfig(cfg)
}
