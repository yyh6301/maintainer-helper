package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	CmdbServiceGroup   cmdb.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
