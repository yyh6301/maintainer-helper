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
	//assetsManageRouterWithoutRecord := Router.Group("cloud")
	//assetsApplyApi := v1.ApiGroupApp.CmdbApiGroup.AssetsApplyApi
	assetsQueryApi := v1.ApiGroupApp.CmdbApiGroup.AssetsQueryApi
	//assetsRecycleApi := v1.ApiGroupApp.CmdbApiGroup.AssetsRecycleApi
	{
		assetsManageRouter.GET("getAssetsList", assetsQueryApi.GetAssetsList) // 查询资产
		//assetsManageRouter.POST("apply", assetsApplyApi.ApplyAssets)       // 申请资产
		//assetsManageRouter.POST("recycle", assetsRecycleApi.RecycleAssets) // 回收资产
	}

}
