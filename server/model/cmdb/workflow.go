package cmdb

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type WorkFlowTemplate struct {
	global.GVA_MODEL
	FlowName       string `json:"flowName" gorm:"comment:流程名称"`
	FlowDesc       string `json:"flowDesc" gorm:"comment:流程描述"`
	FlowDetail     string `json:"flowDetail" gorm:"type:json;comment:流程详情"`
	FlowFormDetail string `json:"flowFormDetail" gorm:"type:json;comment:流程表单详情"`
	FlowCreator    string `json:"flowCreator" gorm:"comment:流程创建者"`
	FlowModifier   string `json:"flowModifier" gorm:"comment:流程修改者"`
}

type WorkFlowOrder struct {
	global.GVA_MODEL
	TemplateID    uint   `json:"templateID" gorm:"comment:流程模板ID"`
	Title         string `json:"title" gorm:"comment:流程标题"`
	OrderDetail   string `json:"orderDetail" gorm:"type:json;comment:流程详情"`
	OrderStatus   string `json:"orderStatus" gorm:"comment:流程状态"`
	OrderCreator  string `json:"orderCreator" gorm:"comment:流程创建者"`
	OrderModifier string `json:"orderModifier" gorm:"comment:流程修改者"`
}
