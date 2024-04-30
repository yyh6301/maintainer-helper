package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
	Upload example.RouterGroup
	Cmdb   cmdb.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
