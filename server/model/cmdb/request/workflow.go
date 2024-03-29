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
	Handler     string `json:"handler" form:"handler"`
	Application string `json:"application" form:"application"`
	request.PageInfo
}

type HandleOrderParams struct {
	cmdb.WorkFlowOrder
	Opinion string `json:"opinion" form:"opinion"`
	Handler string `json:"handler" form:"handler"`
	Result  string `json:"result" form:"result"`
}

type SearchCircleParams struct {
	cmdb.WorkFlowCircle
	request.PageInfo
}

type SearchWorkFlowOrderParams struct {
	cmdb.WorkFlowOrder
	request.PageInfo
}
