package cmdb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AssetsQueryService struct {
}

func (s *AssetsQueryService) GetAssetsList(assets cmdb.CloudAssets, info request.PageInfo) (assetsList []cmdb.CloudAssets, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&cmdb.CloudAssets{})

	if assets.InstanceName != "" {
		db = db.Where("instance_name like ?", "%"+assets.InstanceName+"%")
	}

	if assets.PrivateIP != "" {
		db = db.Where("private_ip like ?", "%"+assets.PrivateIP+"%")
	}

	if assets.PublicIP != "" {
		db = db.Where("public_ip like ?", "%"+assets.PublicIP+"%")
	}

	if assets.CloudType != "" {
		db = db.Where("cloud_type = ?", assets.CloudType)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&assetsList).Error
	if err != nil {
		return
	}

	return
}
