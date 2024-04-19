package cmdb

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type WorkFlowTemplate struct {
	global.GVA_MODEL
	FlowName       string `json:"flowName" gorm:"comment:流程名称" form:"flowName" `
	FlowDesc       string `json:"flowDesc" gorm:"comment:流程描述"`
	FlowFormDetail string `json:"flowFormDetail" gorm:"type:json;comment:自定义表单信息"`
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

	//SourceStatus WorkFlowStatus `json:"sourceStatus" gorm:"foreignKey:SourceID;references:ID"`
	//TargetStatus WorkFlowStatus `json:"targetStatus" gorm:"foreignKey:TargetID;references:ID"`
}

type WorkFlowOrder struct {
	global.GVA_MODEL
	TemplateID       uint               `json:"templateID" gorm:"comment:流程模板ID" form:"templateID"`
	Title            string             `json:"title" gorm:"comment:流程标题" form:"title"`
	OrderDetail      string             `json:"orderDetail" gorm:"type:json;comment:流程详情" form:"orderDetail"`
	OrderStatusID    uint               `json:"orderStatusID" gorm:"comment:流程状态ID" form:"orderStatusID"`
	OrderCreator     string             `json:"orderCreator" gorm:"comment:流程创建者" form:"orderCreator"`
	OrderModifier    string             `json:"orderModifier" gorm:"comment:流程修改者" form:"orderModifier"`
	WorkFlowTemplate WorkFlowTemplate   `json:"workFlowTemplate" gorm:"foreignKey:TemplateID;references:ID"`
	WorkFlowOrderLog []WorkFlowOrderLog `json:"workFlowOrderLog" gorm:"foreignKey:OrderID;references:ID"`
	WorkFlowStatus   WorkFlowStatus     `json:"workFlowStatus" gorm:"foreignKey:OrderStatusID;references:ID"`
}

type WorkFlowOrderLog struct {
	global.GVA_MODEL
	OrderID    uint   `json:"orderID" gorm:"comment:工单ID" form:"orderID"`
	TemplateID uint   `json:"templateID" gorm:"comment:流程模板ID" form:"templateID"`
	SourceID   uint   `json:"sourceID" gorm:"comment:源环节ID" form:"sourceID"`
	TargetID   uint   `json:"targetID" gorm:"comment:目标环节ID" form:"targetID"`
	Handler    string `json:"handler" gorm:"comment:流程处理人" form:"handler"`
	Status     uint   `json:"status" gorm:"status:处理状态 0:未处理，1:同意，2:拒绝" form:"status"`
	Opinion    string `json:"opinion" gorm:"comment:处理意见" form:"opinion"`

	SourceStatus WorkFlowStatus `json:"sourceStatus" gorm:"foreignKey:SourceID;references:ID"`
	TargetStatus WorkFlowStatus `json:"targetStatus" gorm:"foreignKey:TargetID;references:ID"`
}

func (w *WorkFlowTemplate) TableName() string {
	return "work_flow_templates"
}

func (w *WorkFlowStatus) TableName() string {
	return "work_flow_status"
}
