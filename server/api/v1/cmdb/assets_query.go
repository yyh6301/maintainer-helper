package cmdb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	assetsRequest "github.com/flipped-aurora/gin-vue-admin/server/model/cmdb/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AssetsQueryApi struct{}

func (a *AssetsQueryApi) GetAssetsList(c *gin.Context) {
	var pageInfo assetsRequest.SearchAssetsParams
	_ = c.ShouldBindQuery(&pageInfo)
	list, total, err := queryService.GetAssetsList(pageInfo.CloudAssets, pageInfo.PageInfo)
	if err != nil {
		global.GVA_LOG.Error("查询机器列表失败!", zap.Error(err))
		response.FailWithMessage("查询机器列表失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "查询机器列表成功", c)
}
