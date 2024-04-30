package cmdb

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type AssetsTransferService struct {
}

func (w AssetsTransferService) CreateAssetsTransfer(cloudTransfer cmdb.CloudTransfer) error {
	err := global.GVA_DB.Create(&cloudTransfer).Error
	if err != nil {
		return err
	}
	return nil
}

func (w AssetsTransferService) GetAssetsTransferList(cloudTransfer cmdb.CloudTransfer, info request.PageInfo) (TransferList []cmdb.CloudTransfer, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&cmdb.CloudTransfer{})

	if cloudTransfer.CloudType != "" {
		db = db.Where("cloud_type = ? ", cloudTransfer.CloudType)
	}

	if cloudTransfer.Owner != "" {
		db = db.Where("owner like ?", "%"+cloudTransfer.Owner+"%")
	}

	if cloudTransfer.ToOwner != "" {
		db = db.Where("instance_type like ?", "%"+cloudTransfer.ToOwner+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	db = db.Preload("WorkFlowOrder")
	db = db.Preload("WorkFlowOrder.WorkFlowStatus")
	err = db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&TransferList).Error
	if err != nil {
		return nil, 0, err
	}
	return
}

func (w AssetsTransferService) DeleteAssetsTransfer(cloudTransfer cmdb.CloudTransfer) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var entity cmdb.CloudTransfer
		err := tx.Where("id = ?", cloudTransfer.ID).First(&entity).Error // 根据id查询api记录
		if errors.Is(err, gorm.ErrRecordNotFound) {                      // api记录不存在
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
