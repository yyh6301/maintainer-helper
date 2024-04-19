package cmdb

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	AssetsQueryApi
	AssetsApplyApi
	AssetsTransferApi
	AssetsRenewApi
	WorkFlowTemplateApi
	WorkflowOrderApi
}

var (
	queryService            = service.ServiceGroupApp.CmdbServiceGroup.AssetsQueryService
	renewService            = service.ServiceGroupApp.CmdbServiceGroup.AssetsRenewService
	transferService         = service.ServiceGroupApp.CmdbServiceGroup.AssetsTransferService
	applyService            = service.ServiceGroupApp.CmdbServiceGroup.AssetsApplyService
	workflowTemplateService = service.ServiceGroupApp.CmdbServiceGroup.WorkFlowTemplateService
	workflowOrderService    = service.ServiceGroupApp.CmdbServiceGroup.WorkFlowOrderService
)
