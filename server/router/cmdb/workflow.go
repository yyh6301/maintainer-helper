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
		workFlowRouter.GET("getTemplateList", workflowTemplateApi.GetWorkFlowTemplateList) // 查询工作流模板
		//workFlowRouter.POST("createTemplate")
		//workFlowRouter.POST("updateTemplate")
		//workFlowRouter.POST("deleteTemplate")
	}
}
