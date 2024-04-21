package cmdb

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AssetsManageRouter struct {
}

func (a *AssetsManageRouter) InitAssetsManageRouter(Router *gin.RouterGroup) {
	assetsManageRouter := Router.Group("cmdb").Use(middleware.OperationRecord())

	assetsApplyApi := v1.ApiGroupApp.CmdbApiGroup.AssetsApplyApi
	assetsQueryApi := v1.ApiGroupApp.CmdbApiGroup.AssetsQueryApi
	assetsTransferApi := v1.ApiGroupApp.CmdbApiGroup.AssetsTransferApi
	assetsRenewApi := v1.ApiGroupApp.CmdbApiGroup.AssetsRenewApi

	{
		assetsManageRouter.GET("getAssetsList", assetsQueryApi.GetAssetsList) // 查询资产
	}

	{
		assetsManageRouter.GET("getApplyList", assetsApplyApi.GetAssetsApplyList)  // 查询申请列表
		assetsManageRouter.POST("createApply", assetsApplyApi.CreateAssetsApply)   // 创建申请
		assetsManageRouter.DELETE("deleteApply", assetsApplyApi.DeleteAssetsApply) // 删除申请
	}

	{
		assetsManageRouter.GET("getTransferList", assetsTransferApi.GetAssetsTransferList)  // 查询转让列表
		assetsManageRouter.POST("createTransfer", assetsTransferApi.CreateAssetsTransfer)   // 转让资产
		assetsManageRouter.DELETE("deleteTransfer", assetsTransferApi.DeleteAssetsTransfer) // 删除转让
	}

	{
		assetsManageRouter.GET("getRenewList", assetsRenewApi.GetAssetsRenewList)  // 查询续费列表
		assetsManageRouter.POST("createRenew", assetsRenewApi.CreateAssetsRenew)   // 续费资产
		assetsManageRouter.DELETE("deleteRenew", assetsRenewApi.DeleteAssetsRenew) // 删除续费
	}

}
