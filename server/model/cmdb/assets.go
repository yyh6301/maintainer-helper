package cmdb

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type CloudAssets struct {
	global.GVA_MODEL
	UUID         string `json:"uuid" gorm:"comment:云资产唯一标识"`
	CloudType    string `json:"cloudType" gorm:"comment:云厂商字段"`
	InstanceName string `json:"instanceName" gorm:"comment:实例名称"`
	PublicIP     string `json:"publicIP" gorm:"comment:公网IP"`
	PrivateIP    string `json:"privateIP" gorm:"comment:私网IP"`
	Region       string `json:"region" gorm:"comment:区域"`
	Zone         string `json:"zone" gorm:"comment:可用区"`
	InstanceType string `json:"instanceType" gorm:"comment:实例类型"`
}

type CloudApply struct {
	global.GVA_MODEL
	CloudType    string `json:"cloudType" gorm:"comment:云厂商"`
	InstanceName string `json:"instanceName" gorm:"comment:实例名称"`
	ClusterName  string `json:"clusterName" gorm:"comment:客户名"`
	Applyer      string `json:"applyer" gorm:"comment:申请人"`
	ApplyReason  string `json:"applyReason" gorm:"comment:申请理由"`
	OrderID      uint   `json:"orderID" gorm:"comment:工单详情ID"`

	WorkFlowOrder WorkFlowOrder `json:"workFlowOrder" gorm:"foreignKey:OrderID;reference:ID"`
}

type CloudTransfer struct {
	global.GVA_MODEL
	CloudType   string `json:"cloudType" gorm:"comment:云厂商"`
	Owner       string `json:"owner" gorm:"comment:申请人"`
	ToOwner     string `json:"toOwner" gorm:"comment:接收人"`
	ApplyReason string `json:"applyReason" gorm:"comment:申请理由"`
	Count       int    `json:"count" gorm:"comment:数量"`
	OrderID     uint   `json:"orderID" gorm:"comment:工单详情ID"`

	WorkFlowOrder WorkFlowOrder `json:"workFlowOrder" gorm:"foreignKey:OrderID;reference:ID"`
}

type CloudRenew struct {
	global.GVA_MODEL
	CloudType    string `json:"cloudType" gorm:"comment:云厂商"`
	InstanceName string `json:"instanceName" gorm:"comment:实例名称"`
	InstanceType string `json:"instanceType" gorm:"comment:实例类型"`
	Applyer      string `json:"applyer" gorm:"comment:申请人"`
	RenewTime    string `json:"renewTime" gorm:"comment:续费时间"`
	OrderID      uint   `json:"orderID" gorm:"comment:工单详情ID"`

	WorkFlowOrder WorkFlowOrder `json:"workFlowOrder" gorm:"foreignKey:OrderID;reference:ID"`
}
