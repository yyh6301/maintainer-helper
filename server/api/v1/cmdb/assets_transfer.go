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

type AssetsRecycleApi struct {
}

func (a *AssetsRecycleApi) GetAssetsRecycleList(c *gin.Context) {
	var pageInfo cmdbRequest.SearchAssetsRecycleParams
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取查询参数失败!", zap.Any("err", err))
		response.FailWithMessage("获取查询参数失败", c)
	}
	list, total, err := RecycleService.GetAssetsRecycleList(pageInfo.CloudRecycle, pageInfo.PageInfo)
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

func (a *AssetsRecycleApi) CreateAssetsRecycle(c *gin.Context) {
	var cloudRecycle cmdb.CloudRecycle
	err := c.ShouldBindJSON(&cloudRecycle)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	//err = utils.Verify(CloudAssets, utils.AssetsRecycleVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = RecycleService.CreateAssetsRecycle(cloudRecycle)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (a *AssetsRecycleApi) DeleteAssetsRecycle(c *gin.Context) {
	var cloudRecycle cmdb.CloudRecycle
	err := c.ShouldBindJSON(&cloudRecycle)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = utils.Verify(cloudRecycle, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = RecycleService.DeleteAssetsRecycle(cloudRecycle)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
