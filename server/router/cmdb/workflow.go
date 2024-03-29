package cmdb

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type WorkFlowRouter struct {
}

func (a *WorkFlowRouter) InitWorkFlowRouter(Router *gin.RouterGroup) {
	workFlowRouter := Router.Group("workflow")
	workflowTemplateApi := v1.ApiGroupApp.CmdbApiGroup.WorkFlowTemplateApi
	{
		workFlowRouter.GET("getTemplateList", workflowTemplateApi.GetWorkFlowTemplateList)  // 查询工作流模板
		workFlowRouter.POST("createTemplate", workflowTemplateApi.CreateWorkFlowTemplate)   // 创建工作流模板
		workFlowRouter.POST("updateTemplate", workflowTemplateApi.UpdateWorkFlowTemplate)   // 更新工作流模板
		workFlowRouter.DELETE("deleteTemplate", workflowTemplateApi.DeleteWorkFlowTemplate) // 删除工作流模板
		workFlowRouter.POST("getTemplateById", workflowTemplateApi.GetWorkFlowTemplateById)

		// 流程状态
		workFlowRouter.GET("getStatusList", workflowTemplateApi.GetTemplateStatus)
		workFlowRouter.POST("createStatus", workflowTemplateApi.CreateTemplateStatus)
		workFlowRouter.POST("updateStatus", workflowTemplateApi.UpdateTemplateStatus)
		workFlowRouter.DELETE("deleteStatus", workflowTemplateApi.DeleteTemplateStatus)
		workFlowRouter.POST("getStatusById", workflowTemplateApi.GetTemplateStatusById)

		// 流程流转
		workFlowRouter.GET("getCircleList", workflowTemplateApi.GetCircleList)
		workFlowRouter.POST("createCircle", workflowTemplateApi.CreateCircle)
		workFlowRouter.POST("updateCircle", workflowTemplateApi.UpdateCircle)
		workFlowRouter.DELETE("deleteCircle", workflowTemplateApi.DeleteCircle)
		workFlowRouter.POST("getCircleById", workflowTemplateApi.GetCircleById)
	}

	workflowOrderApi := v1.ApiGroupApp.CmdbApiGroup.WorkflowOrderApi
	{
		workFlowRouter.POST("createOrder", workflowOrderApi.CreateOrder)
		workFlowRouter.GET("getOrderList", workflowOrderApi.GetOrderList)
		workFlowRouter.POST("updateOrder", workflowOrderApi.UpdateOrder)
		workFlowRouter.DELETE("deleteOrder", workflowOrderApi.DeleteOrder)
		workFlowRouter.POST("getOrderDetailById", workflowOrderApi.GetOrderById)
		workFlowRouter.POST("handleOrder", workflowOrderApi.HandleOrder)
	}

}
