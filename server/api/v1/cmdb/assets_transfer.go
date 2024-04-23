package cmdb

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
	cmdbRequest "github.com/flipped-aurora/gin-vue-admin/server/model/cmdb/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AssetsTransferApi struct {
}

func (a *AssetsTransferApi) GetAssetsTransferList(c *gin.Context) {
	var pageInfo cmdbRequest.SearchAssetsTransferParams
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取查询参数失败!", zap.Any("err", err))
		response.FailWithMessage("获取查询参数失败", c)
	}
	list, total, err := transferService.GetAssetsTransferList(pageInfo.CloudTransfer, pageInfo.PageInfo)
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

func (a *AssetsTransferApi) CreateAssetsTransfer(c *gin.Context) {
	var cloudTransfer cmdbRequest.CreateTransferParams

	err := c.ShouldBindJSON(&cloudTransfer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	//err = utils.Verify(CloudAssets, utils.AssetsTransferVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	//{"cloud_type":"test","count":5,"owner":"adf","to_owner":"zxcv","transfer_info":"zxcvsd"}

	orderDetail := fmt.Sprintf("{\"count\": %d, \"cloud_type\": \"%s\", \"transfer_info\": \"%s\", \"owner\": \"%s\", \"to_owner\": \"%s\"}",
		cloudTransfer.Count, cloudTransfer.CloudType, cloudTransfer.InstanceList, cloudTransfer.Owner, cloudTransfer.ToOwner)

	var order = &cmdb.WorkFlowOrder{
		Title:         fmt.Sprintf("%s-%s机器转让", cloudTransfer.Owner, cloudTransfer.CloudType),
		TemplateID:    cloudTransfer.WorkFlowOrder.TemplateID,
		OrderDetail:   orderDetail,
		OrderCreator:  cloudTransfer.Owner,
		OrderModifier: cloudTransfer.Owner,
	}

	err = workflowOrderService.CreateOrder(order)
	if err != nil {
		global.GVA_LOG.Error("创建工单失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	cloudTransfer.OrderID = order.ID
	cloudTransfer.WorkFlowOrder.ID = order.ID
	cloudTransfer.WorkFlowOrder.OrderDetail = orderDetail //虽然workflowOrder不会再被插入，但是其会对格式进行校验，这里避免格式校验错误

	err = transferService.CreateAssetsTransfer(cloudTransfer.CloudTransfer)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (a *AssetsTransferApi) DeleteAssetsTransfer(c *gin.Context) {
	var cloudTransfer cmdb.CloudTransfer
	err := c.ShouldBindJSON(&cloudTransfer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = transferService.DeleteAssetsTransfer(cloudTransfer)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
