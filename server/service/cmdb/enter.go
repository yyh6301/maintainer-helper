package cmdb

type ServiceGroup struct {
	AssetsQueryService
	AssetsApplyService
	AssetsRenewService
	AssetsTransferService
	WorkFlowTemplateService
	WorkFlowOrderService
}
