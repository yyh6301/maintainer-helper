package cmdb

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	AssetsQueryApi
	AssetsApplyApi
	AssetsRecycleApi
	WorkFlowTemplateApi
	WorkflowOrderApi
}

var (
	queryService            = service.ServiceGroupApp.CmdbServiceGroup.AssetsQueryService
	recycleService          = service.ServiceGroupApp.CmdbServiceGroup.AssetsRecycleService
	applyService            = service.ServiceGroupApp.CmdbServiceGroup.AssetsApplyService
	workflowTemplateService = service.ServiceGroupApp.CmdbServiceGroup.WorkFlowTemplateService
	workflowOrderService    = service.ServiceGroupApp.CmdbServiceGroup.WorkFlowOrderService
)
