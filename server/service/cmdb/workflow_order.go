package cmdb

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type WorkFlowOrderService struct {
}

func (w WorkFlowOrderService) CreateOrder(order *cmdb.WorkFlowOrder) error {
	if !errors.Is(global.GVA_DB.Where("title = ? ", order.Title).First(&cmdb.WorkFlowOrder{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在标题工单名称 ")
	}

	//找到对应工单模板状态的开始节点
	var status cmdb.WorkFlowStatus
	if err := global.GVA_DB.Where("template_id = ? and status_type = ?", order.TemplateID, 0).First(&status).Error; err != nil {
		return err
	}
	//创建工单
	order.OrderStatusID = status.ID
	if err := global.GVA_DB.Create(&order).Error; err != nil {
		return err
	}

	//创建工单以后，默认对工单进行了一次资源修改的提交
	err := w.HandleOrder(*order, "hook", "", "true")
	if err != nil {
		return err
	}

	return nil

}

func (w WorkFlowOrderService) GetOrderList(order cmdb.WorkFlowOrder, handler string, application string, info request.PageInfo) (templateList []cmdb.WorkFlowOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&cmdb.WorkFlowOrder{})

	if order.Title != "" {
		db = db.Where("title LIKE ?", "%"+order.Title+"%")
	}

	if handler != "" {
		db = db.Joins("LEFT JOIN work_flow_status ON work_flow_status.id = work_flow_orders.order_status_id").
			Where("work_flow_status.approval_user = ?", handler)
	}

	if application != "" {
		db = db.Where("order_creator = ?", application)
	}

	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	// 联表查询templateID，获取模板名称
	db = db.Preload("WorkFlowTemplate")
	db = db.Preload("WorkFlowStatus")
	err = db.Limit(limit).Offset(offset).Find(&templateList).Error
	if err != nil {
		return nil, 0, err
	}
	return
}

func (w WorkFlowOrderService) DeleteOrder(order cmdb.WorkFlowOrder) error {
	var entity cmdb.WorkFlowOrder
	err := global.GVA_DB.Where("id = ?", order.ID).First(&entity).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {                         // api记录不存在
		return err
	}
	err = global.GVA_DB.Delete(&entity).Error
	if err != nil {
		return err
	}

	err = global.GVA_DB.Where("order_id = ? ", order.ID).Delete(&cmdb.WorkFlowOrderLog{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (w WorkFlowOrderService) UpdateOrder(order cmdb.WorkFlowOrder) error {
	var entity cmdb.WorkFlowOrder
	err := global.GVA_DB.Where("id = ?", order.ID).First(&entity).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {                         // api记录不存在
		return err
	}
	err = global.GVA_DB.Model(&entity).Updates(&order).Error
	if err != nil {
		return err
	}
	return nil
}

func (w WorkFlowOrderService) GetOrderById(order cmdb.WorkFlowOrder) (res cmdb.WorkFlowOrder, err error) {
	err = global.GVA_DB.Preload("WorkFlowTemplate").
		Preload("WorkFlowStatus").
		Preload("WorkFlowOrderLog", func(db *gorm.DB) *gorm.DB {
			return db.Order("updated_at DESC")
		}).
		Preload("WorkFlowOrderLog.SourceStatus").
		Preload("WorkFlowOrderLog.TargetStatus").
		Where("id = ?", order.ID).First(&res).Error
	return
}

func (w WorkFlowOrderService) HandleOrder(order cmdb.WorkFlowOrder, handler, opinion, result string) error {
	var err error
	err = global.GVA_DB.Where("id = ?", order.ID).First(&order).Error
	if err != nil {
		return err
	}

	//根据orderStatusID找到当前节点
	var status cmdb.WorkFlowStatus
	if err = global.GVA_DB.Where("id = ? ", order.OrderStatusID).First(&status).Error; err != nil {
		return err
	}
	//找到当前节点的下面的流程
	var circleList []cmdb.WorkFlowCircle
	if err := global.GVA_DB.Where("template_id = ? and source_id = ?", order.TemplateID, status.ID).Find(&circleList).Error; err != nil {
		return err
	}
	if len(circleList) == 0 {
		return errors.New("当前状态没有下一步")
	}
	//找到下一个状态
	var nextStatus cmdb.WorkFlowStatus
	for _, circle := range circleList {
		if result == "true" && circle.AttributeType == 0 {
			if err := global.GVA_DB.Where("id = ?", circle.TargetID).First(&nextStatus).Error; err != nil {
				return err
			}
			break
		} else if result == "false" && circle.AttributeType == 1 {
			if err := global.GVA_DB.Where("id = ?", circle.TargetID).First(&nextStatus).Error; err != nil {
				return err
			}
			break
		}
	}

	//修改工单的状态
	err = global.GVA_DB.Model(&order).Update("order_status_id", nextStatus.ID).Error
	if err != nil {
		return err
	}

	orderLog := cmdb.WorkFlowOrderLog{
		OrderID:    order.ID,
		TemplateID: order.TemplateID,
		SourceID:   status.ID,
		TargetID:   nextStatus.ID,
		Handler:    handler,
		Opinion:    opinion,
		Status:     1,
	}

	if result != "true" {
		orderLog.Status = 2
	}

	err = global.GVA_DB.Create(&orderLog).Error
	if err != nil {
		return err
	}
	return nil

}
