package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	CmdbApiGroup   cmdb.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
