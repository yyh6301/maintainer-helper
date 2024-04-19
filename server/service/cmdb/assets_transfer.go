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
	//1. 使用事务，先插入云资源申请表，再去调用工单申请的流程，把数据插入工单申请表
	return nil
}

func (w AssetsTransferService) GetAssetsTransferList(cloudTransfer cmdb.CloudTransfer, info request.PageInfo) (TransferList []cmdb.CloudTransfer, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&cmdb.CloudTransfer{})

	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Limit(limit).Offset(offset).Find(&TransferList).Error
	if err != nil {
		return nil, 0, err
	}
	return
}

func (w AssetsTransferService) DeleteAssetsTransfer(cloudTransfer cmdb.CloudTransfer) error {
	var entity cmdb.CloudTransfer
	err := global.GVA_DB.Where("id = ?", entity.ID).First(&entity).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {                          // api记录不存在
		return err
	}
	err = global.GVA_DB.Delete(&entity).Error
	if err != nil {
		return err
	}

	//删除工单及对应操作日志
	var order cmdb.WorkFlowOrder
	err = global.GVA_DB.Where("id = ?", entity.OrderID).First(&order).Error
	err = global.GVA_DB.Delete(&order).Error
	if err != nil {
		return err
	}
	return nil
}
