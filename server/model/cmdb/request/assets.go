package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SearchAssetsParams struct {
	cmdb.CloudAssets
	request.PageInfo
}

type SearchAssetsApplyParams struct {
	cmdb.CloudApply
	request.PageInfo
}

type SearchAssetsRenewParams struct {
	cmdb.CloudRenew
	request.PageInfo
}

type SearchAssetsTransferParams struct {
	cmdb.CloudTransfer
	request.PageInfo
}
