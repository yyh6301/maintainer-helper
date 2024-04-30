package cmdb

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type AssetsApplyService struct {
}

func (w AssetsApplyService) CreateAssetsApply(cloudApply cmdb.CloudApply) error {
	var err error
	err = global.GVA_DB.Create(&cloudApply).Error
	if err != nil {
		return err
	}
	return nil
}

func (w AssetsApplyService) GetAssetsApplyList(cloudApply cmdb.CloudApply, info request.PageInfo) (applyList []cmdb.CloudApply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&cmdb.CloudApply{})

	if cloudApply.CloudType != "" {
		db = db.Where("cloud_type = ?", cloudApply.CloudType)
	}

	if cloudApply.ClusterName != "" {
		db = db.Where("cluster_name like ?", "%"+cloudApply.ClusterName+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	db = db.Preload("WorkFlowOrder")
	db = db.Preload("WorkFlowOrder.WorkFlowStatus")

	err = db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&applyList).Error
	if err != nil {
		return nil, 0, err
	}
	return
}

func (w AssetsApplyService) DeleteAssetsApply(cloudApply cmdb.CloudApply) error {

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var entity cmdb.CloudApply
		err := tx.Where("id = ?", cloudApply.ID).First(&entity).Error // 根据id查询api记录
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
