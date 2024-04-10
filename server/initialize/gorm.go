package initialize

import (
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	case "sqlite":
		return GormSqlite()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(
	//cmdb.CloudAssets{},
	//cmdb.WorkFlowOrder{},
	//cmdb.WorkFlowTemplate{},
	//cmdb.WorkFlowStatus{},
	//cmdb.WorkFlowCircle{},
	//cmdb.WorkFlowOrderLog{},
	//system.SysApi{},
	//system.SysUser{},
	//system.SysBaseMenu{},
	//system.JwtBlacklist{},
	//system.SysAuthority{},
	//system.SysDictionary{},
	//system.SysOperationRecord{},
	//system.SysAutoCodeHistory{},
	//system.SysDictionaryDetail{},
	//system.SysBaseMenuParameter{},
	//system.SysBaseMenuBtn{},
	//system.SysAuthorityBtn{},
	//system.SysAutoCode{},
	//system.SysExportTemplate{},
	//
	//example.ExaFile{},
	//example.ExaCustomer{},
	//example.ExaFileChunk{},
	//example.ExaFileUploadAndDownload{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
