package cmdb

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type AssetsRenewService struct {
}

func (w AssetsRenewService) CreateAssetsRenew(cloudRenew cmdb.CloudRenew) error {
	err := global.GVA_DB.Create(&cloudRenew).Error
	if err != nil {
		return err
	}
	return nil
}

func (w AssetsRenewService) GetAssetsRenewList(cloudRenew cmdb.CloudRenew, info request.PageInfo) (RenewList []cmdb.CloudRenew, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&cmdb.CloudRenew{})

	if cloudRenew.CloudType != "" {
		db = db.Where("cloud_type = ? ", cloudRenew.CloudType)
	}

	if cloudRenew.InstanceName != "" {
		db = db.Where("instance_name like ?", "%"+cloudRenew.InstanceName+"%")
	}

	if cloudRenew.InstanceType != "" {
		db = db.Where("instance_type like ?", "%"+cloudRenew.InstanceType+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	db = db.Preload("WorkFlowOrder")
	db = db.Preload("WorkFlowOrder.WorkFlowStatus")

	err = db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&RenewList).Error
	if err != nil {
		return nil, 0, err
	}
	return
}

func (w AssetsRenewService) DeleteAssetsRenew(cloudRenew cmdb.CloudRenew) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var entity cmdb.CloudRenew
		err := tx.Where("id = ?", cloudRenew.ID).First(&entity).Error // 根据id查询api记录
		if errors.Is(err, gorm.ErrRecordNotFound) {                   // api记录不存在
			return err
		}
		err = tx.Delete(&entity).Error
		if err != nil {
			return err
		}

		//删除工单及对应操作日志
		var order cmdb.WorkFlowOrder
		err = tx.Where("id = ?", entity.OrderID).First(&order).Error
		err = tx.Delete(&order).Error
		if err != nil {
			return err
		}

		err = tx.Where("order_id = ? ", order.ID).Delete(&cmdb.WorkFlowOrderLog{}).Error
		if err != nil {
			return err
		}
		return nil
	})
}
