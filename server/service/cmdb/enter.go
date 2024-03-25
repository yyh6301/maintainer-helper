package cmdb

type ServiceGroup struct {
	AssetsQueryService
	AssetsApplyService
	AssetsRecycleService
	WorkFlowTemplateService
	WorkFlowOrderService
}
