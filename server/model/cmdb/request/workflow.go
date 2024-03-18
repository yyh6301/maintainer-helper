package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SearchWorkFlowTemplateParams struct {
	cmdb.WorkFlowTemplate
	request.PageInfo
}

type SearchWorkFlowOrderParams struct {
	cmdb.WorkFlowOrder
	request.PageInfo
}
