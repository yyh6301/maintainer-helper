package system

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DBApi
	JwtApi
	BaseApi
	SystemApi
	CasbinApi
	AuthorityApi
	AuthorityMenuApi
	OperationRecordApi
	AuthorityBtnApi
}

var (
	jwtService             = service.ServiceGroupApp.SystemServiceGroup.JwtService
	menuService            = service.ServiceGroupApp.SystemServiceGroup.MenuService
	userService            = service.ServiceGroupApp.SystemServiceGroup.UserService
	casbinService          = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	baseMenuService        = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
	authorityService       = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	systemConfigService    = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
	operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
	authorityBtnService    = service.ServiceGroupApp.SystemServiceGroup.AuthorityBtnService
)
