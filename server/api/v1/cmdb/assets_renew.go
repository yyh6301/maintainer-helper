package cmdb

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
	cmdbRequest "github.com/flipped-aurora/gin-vue-admin/server/model/cmdb/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type AssetsRenewApi struct {
}

func (a *AssetsRenewApi) GetAssetsRenewList(c *gin.Context) {
	var pageInfo cmdbRequest.SearchAssetsRenewParams
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取查询参数失败!", zap.Any("err", err))
		response.FailWithMessage("获取查询参数失败", c)
	}
	list, total, err := renewService.GetAssetsRenewList(pageInfo.CloudRenew, pageInfo.PageInfo)
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

func (a *AssetsRenewApi) CreateAssetsRenew(c *gin.Context) {
	var cloudRenew cmdb.CloudRenew
	err := c.ShouldBindJSON(&cloudRenew)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	//err = utils.Verify(CloudAssets, utils.AssetsRenewVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(cloudRenew.WorkFlowOrder.OrderDetail), &data)
	instanceName, ok := data["instance_name"].(string)
	if !ok {
		global.GVA_LOG.Warn("申请工单时instance_name字段不存在")
		response.FailWithMessage("instance_name字段为空", c)
		return
	}

	instanceType, ok := data["instance_type"].(string)
	if !ok {
		global.GVA_LOG.Warn("申请工单时instance_type字段不存在")
		response.FailWithMessage("instance_type字段为空", c)
		return
	}
	renewTime, ok := data["renew_time"].(float64)
	if !ok {
		global.GVA_LOG.Warn("申请工单时time字段不存在")
		response.FailWithMessage("time字段为空", c)
		return
	}
	cloudRenew.RenewTime = strconv.Itoa(int(renewTime))
	cloudRenew.InstanceType = instanceType
	cloudRenew.InstanceName = instanceName

	var order = &cmdb.WorkFlowOrder{
		Title:         fmt.Sprintf("%s机器续费-%s", cloudRenew.CloudType, cloudRenew.InstanceName),
		TemplateID:    cloudRenew.WorkFlowOrder.TemplateID,
		OrderDetail:   cloudRenew.WorkFlowOrder.OrderDetail,
		OrderCreator:  cloudRenew.Applyer,
		OrderModifier: cloudRenew.Applyer,
	}

	err = workflowOrderService.CreateOrder(order)
	if err != nil {
		global.GVA_LOG.Error("创建工单失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	cloudRenew.OrderID = order.ID
	cloudRenew.WorkFlowOrder.ID = order.ID

	err = renewService.CreateAssetsRenew(cloudRenew)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (a *AssetsRenewApi) DeleteAssetsRenew(c *gin.Context) {
	var cloudRenew cmdb.CloudRenew
	err := c.ShouldBindJSON(&cloudRenew)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = renewService.DeleteAssetsRenew(cloudRenew)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
