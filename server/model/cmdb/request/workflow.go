package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SearchWorkFlowTemplateParams struct {
	cmdb.WorkFlowTemplate
	request.PageInfo
}

type SearchTemplateStatusParams struct {
	cmdb.WorkFlowStatus
	request.PageInfo
}

type SearchOrderParams struct {
	cmdb.WorkFlowOrder
	Handler string `json:"handler" form:"handler"`
	request.PageInfo
}

type SearchCircleParams struct {
	cmdb.WorkFlowCircle
	request.PageInfo
}

type SearchWorkFlowOrderParams struct {
	cmdb.WorkFlowOrder
	request.PageInfo
}
