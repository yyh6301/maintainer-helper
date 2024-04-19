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

func (w WorkFlowOrderService) CreateOrder(order cmdb.WorkFlowOrder) error {
	if !errors.Is(global.GVA_DB.Where("title = ? ", order.Title).First(&cmdb.WorkFlowOrder{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在标题工单名称 ")
	}
	//处理工单log表，找到对应工单模板状态的开始节点，找到开始节点的下一步，这里log设置状态成待处理，处理完成再修改log状态，添加新的状态到log中.

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {

		//创建工单log
		//找到对应工单模板状态的开始节点
		var status cmdb.WorkFlowStatus
		if err := tx.Where("template_id = ? and status_type = ?", order.TemplateID, 0).First(&status).Error; err != nil {
			return err
		}
		//找到开始节点的下一步
		var circleList []cmdb.WorkFlowCircle
		if err := tx.Where("template_id = ? and source_id = ?", order.TemplateID, status.ID).Find(&circleList).Error; err != nil {
			return err
		}
		if len(circleList) == 0 {
			return errors.New("开始节点没有下一步")
		}
		for _, circle := range circleList {
			//这里log设置状态成待处理
			orderLog := cmdb.WorkFlowOrderLog{
				OrderID:    order.ID,
				TemplateID: order.TemplateID,
				SourceID:   status.ID,
				TargetID:   circle.TargetID,
				Status:     0,
			}
			if err := tx.Create(&orderLog).Error; err != nil {
				return err
			}
		}
		//创建工单
		order.OrderStatusID = status.ID
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		return nil
	})
}

func (w WorkFlowOrderService) GetOrderList(order cmdb.WorkFlowOrder, handler string, application string, info request.PageInfo) (templateList []cmdb.WorkFlowOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&cmdb.WorkFlowOrder{})

	if order.Title != "" {
		db = db.Where("title LIKE ?", "%"+order.Title+"%")
	}

	if handler != "" {
		db = db.Joins("LEFT JOIN work_flow_order_logs ON work_flow_order_logs.order_id = work_flow_orders.id").
			Where("work_flow_order_logs.handler = ?", handler).
			Where("work_flow_order_logs.status = ?", 0)
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
		Preload("WorkFlowOrderLog", func(db *gorm.DB) *gorm.DB {
			return db.Order("updated_at DESC")
		}).
		Preload("WorkFlowOrderLog.SourceStatus").
		Preload("WorkFlowOrderLog.TargetStatus").
		Where("id = ?", order.ID).First(&res).Error
	return
}

func (w WorkFlowOrderService) HandleOrder(order cmdb.WorkFlowOrder, handler, opinion, result string) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		tx.Where("id = ?", order.ID).First(&order)
		//根据orderStatusID找到当前节点
		var status cmdb.WorkFlowStatus
		if err := tx.Where("id = ? ", order.OrderStatusID).First(&status).Error; err != nil {
			return err
		}
		//找到当前节点的下面的流程
		var circleList []cmdb.WorkFlowCircle
		if err := tx.Where("template_id = ? and source_id = ?", order.TemplateID, status.ID).Find(&circleList).Error; err != nil {
			return err
		}
		if len(circleList) == 0 {
			return errors.New("当前状态没有下一步")
		}
		//找到下一个状态,todo 这里后面的attributeType换成可自定义条件
		var nextStatus cmdb.WorkFlowStatus
		for _, circle := range circleList {
			if result == "true" && circle.AttributeType == 0 {
				if err := tx.Where("id = ?", circle.TargetID).First(&nextStatus).Error; err != nil {
					return err
				}
				break
			} else if result == "false" && circle.AttributeType == 1 {
				if err := tx.Where("id = ?", circle.TargetID).First(&nextStatus).Error; err != nil {
					return err
				}
				break
			}
		}
		//创建orderlog
		var s uint = 0
		if result == "true" {
			s = 1
		} else {
			s = 2
		}
		orderLog := cmdb.WorkFlowOrderLog{
			OrderID:    order.ID,
			TemplateID: order.TemplateID,
			SourceID:   status.ID,
			TargetID:   nextStatus.ID,
			Handler:    handler,
			Opinion:    opinion,
			Status:     s,
		}
		err := tx.Create(&orderLog).Error
		if err != nil {
			return err
		}
		return nil
	})
}
