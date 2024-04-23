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

// =================流程模板===================

func (a *WorkFlowTemplateApi) GetWorkFlowTemplateList(c *gin.Context) {
	var pageInfo cmdbRequest.SearchWorkFlowTemplateParams
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取查询参数失败!", zap.Any("err", err))
		response.FailWithMessage("获取查询参数失败", c)
		return
	}
	list, total, err := workflowTemplateService.GetWorkFlowTemplateList(pageInfo.WorkFlowTemplate, pageInfo.PageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取工作流模板列表失败!", zap.Any("err", err))
		response.FailWithMessage("获取工作流模板列表失败", c)
		return
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
		return
	}
	//err = utils.Verify(workflowTemplate, utils.WorkFlowTemplateVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
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
		return
	}
	//err = utils.Verify(workflowTemplate, utils.WorkFlowTemplateVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
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
		return
	}
	err = utils.Verify(workflowTemplate, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
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
		return
	}
	err = utils.Verify(workflowTemplate, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	workflow, err := workflowTemplateService.GetWorkFlowTemplateById(workflowTemplate)
	if err != nil {
		global.GVA_LOG.Error("获取工作流模版信息失败,请重试!", zap.Error(err))
		response.FailWithMessage("获取工作流模板失败", c)
	}
	response.OkWithDetailed(workflow, "获取工作流模板信息成功", c)
}

// =================流程状态===================

func (a *WorkFlowTemplateApi) GetTemplateStatus(c *gin.Context) {
	var pageInfo cmdbRequest.SearchTemplateStatusParams
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取查询参数失败!", zap.Any("err", err))
		response.FailWithMessage("获取查询参数失败", c)
		return
	}
	list, total, err := workflowTemplateService.GetTemplateStatusList(pageInfo.WorkFlowStatus, pageInfo.PageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取工作流模板列表失败!", zap.Any("err", err))
		response.FailWithMessage("获取工作流模板列表失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取工作流模板列表成功", c)
}

func (a *WorkFlowTemplateApi) CreateTemplateStatus(c *gin.Context) {
	var templateStatus cmdb.WorkFlowStatus
	err := c.ShouldBindJSON(&templateStatus)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//err = utils.Verify(templateStatus, utils.WorkFlowStatusVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = workflowTemplateService.CreateTemplateStatus(templateStatus)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (a *WorkFlowTemplateApi) UpdateTemplateStatus(c *gin.Context) {
	var templateStatus cmdb.WorkFlowStatus
	err := c.ShouldBindJSON(&templateStatus)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(templateStatus, utils.WorkFlowStatusVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = workflowTemplateService.UpdateTemplateStatus(templateStatus)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (a *WorkFlowTemplateApi) DeleteTemplateStatus(c *gin.Context) {
	var templateStatus cmdb.WorkFlowStatus
	err := c.ShouldBindJSON(&templateStatus)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(templateStatus, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = workflowTemplateService.DeleteTemplateStatus(templateStatus)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (a *WorkFlowTemplateApi) GetTemplateStatusById(c *gin.Context) {
	var status cmdb.WorkFlowStatus
	err := c.ShouldBindJSON(&status)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(status, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	workflow, err := workflowTemplateService.GetTemplateStatusById(status)
	if err != nil {
		global.GVA_LOG.Error("获取流程状态信息失败,请重试!", zap.Error(err))
		response.FailWithMessage("获取流程状态失败", c)
		return
	}
	response.OkWithDetailed(workflow, "获取流程状态成功", c)
}

// =================流程流转===================

func (a *WorkFlowTemplateApi) GetCircleList(c *gin.Context) {
	var pageInfo cmdbRequest.SearchCircleParams
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取查询参数失败!", zap.Any("err", err))
		response.FailWithMessage("获取查询参数失败", c)
		return
	}
	list, total, err := workflowTemplateService.GetCircleList(pageInfo.WorkFlowCircle, pageInfo.PageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取工作流模板列表失败!", zap.Any("err", err))
		response.FailWithMessage("获取工作流模板列表失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取流转列表成功", c)
}

func (a *WorkFlowTemplateApi) CreateCircle(c *gin.Context) {
	var circle cmdb.WorkFlowCircle
	err := c.ShouldBindJSON(&circle)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(circle, utils.WorkFlowCircleVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = workflowTemplateService.CreateCircle(circle)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (a *WorkFlowTemplateApi) UpdateCircle(c *gin.Context) {
	var circle cmdb.WorkFlowCircle
	err := c.ShouldBindJSON(&circle)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//err = utils.Verify(circle, utils.WorkFlowStatusVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = workflowTemplateService.UpdateCircle(circle)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (a *WorkFlowTemplateApi) DeleteCircle(c *gin.Context) {
	var circle cmdb.WorkFlowCircle
	err := c.ShouldBindJSON(&circle)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(circle, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = workflowTemplateService.DeleteCircle(circle)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (a *WorkFlowTemplateApi) GetCircleById(c *gin.Context) {
	var circle cmdb.WorkFlowCircle
	err := c.ShouldBindJSON(&circle)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(circle, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	workflow, err := workflowTemplateService.GetCircleById(circle)
	if err != nil {
		global.GVA_LOG.Error("获取流程状态信息失败,请重试!", zap.Error(err))
		response.FailWithMessage("获取流程状态失败", c)
	}
	response.OkWithDetailed(workflow, "获取流程状态成功", c)
}
