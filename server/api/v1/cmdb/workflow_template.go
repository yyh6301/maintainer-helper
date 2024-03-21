package cmdb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	cmdbRequest "github.com/flipped-aurora/gin-vue-admin/server/model/cmdb/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WorkFlowTemplateApi struct {
}

func (a *WorkFlowTemplateApi) GetWorkFlowTemplateList(c *gin.Context) {
	var pageInfo cmdbRequest.SearchWorkFlowTemplateParams
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取查询参数失败!", zap.Any("err", err))
		response.FailWithMessage("获取查询参数失败", c)
	}
	list, total, err := workflowTemplateService.GetWorkFlowTemplateList(pageInfo.WorkFlowTemplate, pageInfo.PageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取工作流模板列表失败!", zap.Any("err", err))
		response.FailWithMessage("获取工作流模板列表失败", c)
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取工作流模板列表成功", c)
}

func (a *WorkFlowTemplateApi) CreateWorkFlowTemplate() {

}

func (a *WorkFlowTemplateApi) UpdateWorkFlowTemplate() {

}

func (a *WorkFlowTemplateApi) DeleteWorkFlowTemplate() {
}
