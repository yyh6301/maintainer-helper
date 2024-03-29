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
		order.OrderStatus = status.StatusName
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
		Preload("WorkFlowOrderLog").
		Where("id = ?", order.ID).First(&res).Error
	return
}

func (w WorkFlowOrderService) HandleOrder(order cmdb.WorkFlowOrder, handler, option, result string) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		//根据templateId及orderStatus字段,找到当前order的状态,并且找到下一个状态列表
		var status cmdb.WorkFlowStatus
		if err := tx.Where("template_id = ? and status_name = ?", order.TemplateID, order.OrderStatus).First(&status).Error; err != nil {
			return err
		}
		var circleList []cmdb.WorkFlowCircle
		if err := tx.Where("template_id = ? and source_id = ?", order.TemplateID, status.ID).Find(&circleList).Error; err != nil {
			return err
		}
		if len(circleList) == 0 {
			return errors.New("当前状态没有下一步")
		}
		//找到下一个状态
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
		//更新order状态
		order.OrderStatus = nextStatus.StatusName
		if err := tx.Model(&order).Updates(&order).Error; err != nil {
			return err
		}
		//更新orderLog状态
		var orderLog cmdb.WorkFlowOrderLog
		if err := tx.Where("order_id = ? and source_id = ? and status = ?", order.ID, status.ID, 0).First(&orderLog).Error; err != nil {
			return err
		}
		orderLog.Status = 1
		orderLog.Handler = handler
		orderLog.Opinion = option
		if err := tx.Model(&orderLog).Updates(&orderLog).Error; err != nil {
			return err
		}
		//根据nextStatus创建新的orderLog
		var nextCircleList []cmdb.WorkFlowCircle
		if err := tx.Where("template_id = ? and source_id = ?", order.TemplateID, nextStatus.ID).Find(&nextCircleList).Error; err != nil {
			return err
		}
		if len(nextCircleList) == 0 {
			return errors.New("下一个状态没有下一步")
		}
		for _, circle := range nextCircleList {
			//这里log设置状态成待处理
			orderLog := cmdb.WorkFlowOrderLog{
				OrderID:    order.ID,
				TemplateID: order.TemplateID,
				SourceID:   nextStatus.ID,
				TargetID:   circle.TargetID,
				Status:     0,
			}
			if err := tx.Create(&orderLog).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
