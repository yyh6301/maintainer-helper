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
