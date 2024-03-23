package cmdb

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type WorkFlowTemplate struct {
	global.GVA_MODEL
	FlowName       string `json:"flowName" gorm:"comment:流程名称" form:"flowName" `
	FlowDesc       string `json:"flowDesc" gorm:"comment:流程描述"`
	FlowDetail     string `json:"flowDetail" gorm:"type:json;comment:流程详情"`
	FlowFormDetail string `json:"flowFormDetail" gorm:"type:json;comment:流程表单详情"`
	FlowCreator    string `json:"flowCreator" gorm:"comment:流程创建者"`
	FlowModifier   string `json:"flowModifier" gorm:"comment:流程修改者"`

	//WorkFlowStatus []WorkFlowStatus `json:"workFlowStatus" gorm:"foreignKey:Template;references:ID"`
}

type WorkFlowStatus struct {
	global.GVA_MODEL
	StatusName   string `json:"statusName" gorm:"comment:状态名称"  form:"statusName"'`
	StatusType   int64  `json:"statusType" gorm:"comment:状态类型,0:开始,1:中间,2:结束"`
	ApprovalType int64  `json:"approvalType" gorm:"comment:审批类型,0:个人,1:hook"`
	ApprovalUser string `json:"approvalUser" gorm:"comment:审批人" form:"approvalUser" `
	OrderNumber  int64  `json:"orderNumber" gorm:"comment:状态排序"`
	TemplateID   uint   `json:"templateID" gorm:"comment:流程模板ID" form:"templateID" validate:"required"`
}

type WorkFlowCircle struct {
	global.GVA_MODEL
	CircleName    string `json:"circleName" gorm:"comment:环节名称" form:"circleName"`
	SourceID      uint   `json:"sourceID" gorm:"comment:源环节ID" form:"sourceID"`
	TargetID      uint   `json:"targetID" gorm:"comment:目标环节ID" form:"targetID"`
	AttributeType int64  `json:"attributeType" gorm:"comment:属性类型,0:同意,1:拒绝" form:"attributeType"`
	TemplateID    uint   `json:"templateID" gorm:"comment:流程模板ID" form:"templateID"`
}

//type WorkFlowOrderLog struct {
//
//}

type WorkFlowOrder struct {
	global.GVA_MODEL
	TemplateID    uint   `json:"templateID" gorm:"comment:流程模板ID"`
	Title         string `json:"title" gorm:"comment:流程标题"`
	OrderDetail   string `json:"orderDetail" gorm:"type:json;comment:流程详情"`
	OrderStatus   string `json:"orderStatus" gorm:"comment:流程状态"`
	OrderCreator  string `json:"orderCreator" gorm:"comment:流程创建者"`
	OrderModifier string `json:"orderModifier" gorm:"comment:流程修改者"`
}

func (w *WorkFlowTemplate) TableName() string {
	return "work_flow_templates"
}

func (w *WorkFlowStatus) TableName() string {
	return "work_flow_status"
}
