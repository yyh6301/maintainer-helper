package cmdb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
	cmdbRequest "github.com/flipped-aurora/gin-vue-admin/server/model/cmdb/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
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

func (a *WorkFlowTemplateApi) CreateWorkFlowTemplate(c *gin.Context) {
	var workflowTemplate cmdb.WorkFlowTemplate
	err := c.ShouldBindJSON(&workflowTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = utils.Verify(workflowTemplate, utils.WorkFlowTemplateVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = workflowTemplateService.CreateWorkFlowTemplate(workflowTemplate)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (a *WorkFlowTemplateApi) UpdateWorkFlowTemplate(c *gin.Context) {
	var workflowTemplate cmdb.WorkFlowTemplate
	err := c.ShouldBindJSON(&workflowTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = utils.Verify(workflowTemplate, utils.WorkFlowTemplateVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = workflowTemplateService.UpdateWorkFlowTemplate(workflowTemplate)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (a *WorkFlowTemplateApi) DeleteWorkFlowTemplate(c *gin.Context) {
	var workflowTemplate cmdb.WorkFlowTemplate
	err := c.ShouldBindJSON(&workflowTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = utils.Verify(workflowTemplate, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = workflowTemplateService.DeleteWorkFlowTemplate(workflowTemplate)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (a *WorkFlowTemplateApi) GetWorkFlowTemplateById(c *gin.Context) {
	var workflowTemplate cmdb.WorkFlowTemplate
	err := c.ShouldBindJSON(&workflowTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = utils.Verify(workflowTemplate, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	workflow, err := workflowTemplateService.GetWorkFlowTemplateById(workflowTemplate)
	if err != nil {
		global.GVA_LOG.Error("获取工作流模版信息失败,请重试!", zap.Error(err))
		response.FailWithMessage("获取工作流模板失败", c)
	}
	response.OkWithDetailed(workflow, "获取工作流模板信息成功", c)
}
