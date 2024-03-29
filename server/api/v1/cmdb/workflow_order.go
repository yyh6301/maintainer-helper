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

type WorkflowOrderApi struct{}

func (a *WorkflowOrderApi) GetOrderList(c *gin.Context) {
	var pageInfo cmdbRequest.SearchOrderParams
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取查询参数失败!", zap.Any("err", err))
		response.FailWithMessage("获取查询参数失败", c)
	}
	list, total, err := workflowOrderService.GetOrderList(pageInfo.WorkFlowOrder, pageInfo.Handler, pageInfo.Application, pageInfo.PageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取所有工单列表失败!", zap.Any("err", err))
		response.FailWithMessage("获取所有列表失败", c)
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取工作流模板列表成功", c)
}

func (a *WorkflowOrderApi) CreateOrder(c *gin.Context) {
	var workflowTemplate cmdb.WorkFlowOrder
	err := c.ShouldBindJSON(&workflowTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = utils.Verify(workflowTemplate, utils.OrderVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = workflowOrderService.CreateOrder(workflowTemplate)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (a *WorkflowOrderApi) UpdateOrder(c *gin.Context) {
	var order cmdb.WorkFlowOrder
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = utils.Verify(order, utils.OrderVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = workflowOrderService.UpdateOrder(order)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (a *WorkflowOrderApi) DeleteOrder(c *gin.Context) {
	var order cmdb.WorkFlowOrder
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = utils.Verify(order, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = workflowOrderService.DeleteOrder(order)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (a *WorkflowOrderApi) GetOrderById(c *gin.Context) {
	var order cmdb.WorkFlowOrder
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	workflow, err := workflowOrderService.GetOrderById(order)
	if err != nil {
		global.GVA_LOG.Error("获取工作流模版信息失败,请重试!", zap.Error(err))
		response.FailWithMessage("获取工作流模板失败", c)
	}
	response.OkWithDetailed(workflow, "获取工作流模板信息成功", c)
}

func (a *WorkflowOrderApi) HandleOrder(c *gin.Context) {
	var order cmdbRequest.HandleOrderParams
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = workflowOrderService.HandleOrder(order.WorkFlowOrder, order.Handler, order.Opinion, order.Result)
	if err != nil {
		global.GVA_LOG.Error("处理失败!", zap.Error(err))
		response.FailWithMessage("处理失败", c)
		return
	}
	response.OkWithMessage("处理成功", c)
}
