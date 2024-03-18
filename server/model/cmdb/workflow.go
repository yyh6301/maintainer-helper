package cmdb

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type WorkFlowTemplate struct {
	global.GVA_MODEL
	FlowName     string `json:"flowName" gorm:"comment:流程名称"`
	FlowType     string `json:"flowType" gorm:"comment:流程类型"`
	FlowDetail   string `json:"flowDetail" gorm:"type:json;comment:流程详情"`
	FlowVersion  string `json:"flowVersion" gorm:"comment:流程版本"`
	FlowCreator  string `json:"flowCreator" gorm:"comment:流程创建者"`
	FlowModifier string `json:"flowModifier" gorm:"comment:流程修改者"`
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
