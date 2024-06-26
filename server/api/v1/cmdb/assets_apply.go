package cmdb

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
	cmdbRequest "github.com/flipped-aurora/gin-vue-admin/server/model/cmdb/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AssetsApplyApi struct {
}

func (a *AssetsApplyApi) GetAssetsApplyList(c *gin.Context) {
	var pageInfo cmdbRequest.SearchAssetsApplyParams
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取查询参数失败!", zap.Any("err", err))
		response.FailWithMessage("获取查询参数失败", c)
		return
	}
	list, total, err := applyService.GetAssetsApplyList(pageInfo.CloudApply, pageInfo.PageInfo)
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

func (a *AssetsApplyApi) CreateAssetsApply(c *gin.Context) {
	var cloudApply cmdb.CloudApply
	err := c.ShouldBindJSON(&cloudApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	//err = utils.Verify(CloudAssets, utils.AssetsApplyVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(cloudApply.WorkFlowOrder.OrderDetail), &data)
	instanceName, ok := data["instance_name"].(string)
	if !ok {
		global.GVA_LOG.Warn("申请工单时instance_name字段不存在")
		response.FailWithMessage("instance_name字段为空", c)
		return
	}
	cloudApply.InstanceName = instanceName

	var order = &cmdb.WorkFlowOrder{
		Title:         fmt.Sprintf("%s机器申请-%s", cloudApply.CloudType, cloudApply.InstanceName),
		TemplateID:    cloudApply.WorkFlowOrder.TemplateID,
		OrderDetail:   cloudApply.WorkFlowOrder.OrderDetail,
		OrderCreator:  cloudApply.Applyer,
		OrderModifier: cloudApply.Applyer,
	}

	err = workflowOrderService.CreateOrder(order)
	if err != nil {
		global.GVA_LOG.Error("创建工单失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	cloudApply.OrderID = order.ID
	cloudApply.WorkFlowOrder.ID = order.ID

	err = applyService.CreateAssetsApply(cloudApply)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (a *AssetsApplyApi) DeleteAssetsApply(c *gin.Context) {
	var cloudApply cmdb.CloudApply
	err := c.ShouldBindJSON(&cloudApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(cloudApply, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = applyService.DeleteAssetsApply(cloudApply)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
