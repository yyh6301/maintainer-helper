package cmdb

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type WorkFlowTemplateService struct {
}

func (w WorkFlowTemplateService) CreateWorkFlowTemplate(workflow cmdb.WorkFlowTemplate) error {
	if workflow.FlowDetail == "" {
		defaultJSON := "{}" // 设置默认的空JSON字符串
		workflow.FlowDetail = defaultJSON
	}

	if workflow.FlowFormDetail == "" {
		defaultJSON := "{}" // 设置默认的空JSON字符串
		workflow.FlowFormDetail = defaultJSON
	}

	if !errors.Is(global.GVA_DB.Where("flow_name = ? ", workflow.FlowName).First(&cmdb.WorkFlowTemplate{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同流程名称 ")
	}
	return global.GVA_DB.Create(&workflow).Error
}

func (w WorkFlowTemplateService) GetWorkFlowTemplateList(workflow cmdb.WorkFlowTemplate, info request.PageInfo) (templateList []cmdb.WorkFlowTemplate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&cmdb.WorkFlowTemplate{})

	if workflow.FlowName != "" {
		db = db.Where("flow_name LIKE ?", "%"+workflow.FlowName+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Limit(limit).Offset(offset).Find(&templateList).Error
	if err != nil {
		return nil, 0, err
	}
	return
}

func (w WorkFlowTemplateService) DeleteWorkFlowTemplate(workflow cmdb.WorkFlowTemplate) error {
	var entity cmdb.WorkFlowTemplate
	err := global.GVA_DB.Where("id = ?", workflow.ID).First(&entity).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {                            // api记录不存在
		return err
	}
	err = global.GVA_DB.Delete(&entity).Error
	if err != nil {
		return err
	}
	return nil
}

func (w WorkFlowTemplateService) UpdateWorkFlowTemplate(workflow cmdb.WorkFlowTemplate) error {
	var entity cmdb.WorkFlowTemplate
	err := global.GVA_DB.Where("id = ?", workflow.ID).First(&entity).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {                            // api记录不存在
		return err
	}
	err = global.GVA_DB.Model(&entity).Updates(&workflow).Error
	if err != nil {
		return err
	}
	return nil
}

func (w WorkFlowTemplateService) GetWorkFlowTemplateById(workflow cmdb.WorkFlowTemplate) (res cmdb.WorkFlowTemplate, err error) {
	err = global.GVA_DB.Where("id = ?", workflow.ID).First(&res).Error
	return
}

// ===============模板状态============

func (w WorkFlowTemplateService) GetTemplateStatusList(status cmdb.WorkFlowStatus, info request.PageInfo) (templateList []cmdb.WorkFlowStatus, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&cmdb.WorkFlowStatus{})

	if status.StatusName != "" {
		db = db.Where("status_name LIKE ?", "%"+status.StatusName+"%")
	}

	if status.TemplateID != 0 {
		db = db.Where("template_id = ?", status.TemplateID)
	}
	if status.ApprovalUser != "" {
		db = db.Where("approval_user LIKE ?", "%"+status.ApprovalUser+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	db.Order("order_number")
	err = db.Limit(limit).Offset(offset).Find(&templateList).Error
	if err != nil {
		return nil, 0, err
	}
	return
}

func (w WorkFlowTemplateService) CreateTemplateStatus(status cmdb.WorkFlowStatus) error {
	if !errors.Is(global.GVA_DB.Where("status_name = ? ", status.StatusName).First(&cmdb.WorkFlowStatus{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同状态名称 ")
	}
	return global.GVA_DB.Create(&status).Error
}

func (w WorkFlowTemplateService) UpdateTemplateStatus(status cmdb.WorkFlowStatus) error {
	var entity cmdb.WorkFlowStatus
	err := global.GVA_DB.Where("id = ?", status.ID).First(&entity).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {                          // api记录不存在
		return err
	}
	err = global.GVA_DB.Model(&entity).Updates(&status).Error
	if err != nil {
		return err
	}
	return nil
}

func (w WorkFlowTemplateService) DeleteTemplateStatus(status cmdb.WorkFlowStatus) error {
	var entity cmdb.WorkFlowStatus
	err := global.GVA_DB.Where("id = ?", status.ID).First(&entity).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {                          // api记录不存在
		return err
	}
	err = global.GVA_DB.Scopes().Delete(&entity).Error
	if err != nil {
		return err
	}
	return nil
}

func (w WorkFlowTemplateService) GetTemplateStatusById(status cmdb.WorkFlowStatus) (res cmdb.WorkFlowStatus, err error) {
	err = global.GVA_DB.Where("id = ?", status.ID).First(&res).Error
	return
}

// =====================流转环节====================

func (w WorkFlowTemplateService) GetCircleList(circle cmdb.WorkFlowCircle, info request.PageInfo) (circleList []cmdb.WorkFlowCircle, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&cmdb.WorkFlowCircle{})

	if circle.CircleName != "" {
		db = db.Where("circle_name LIKE ?", "%"+circle.CircleName+"%")
	}

	if circle.TemplateID != 0 {
		db = db.Where("template_id = ?", circle.TemplateID)
	}

	if circle.SourceID != 0 {
		db = db.Where("source_id = ?", circle.SourceID)
	}

	if circle.TargetID != 0 {
		db = db.Where("target_id = ?", circle.TargetID)
	}

	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Limit(limit).Offset(offset).Find(&circleList).Error
	if err != nil {
		return nil, 0, err
	}
	return
}

func (w WorkFlowTemplateService) CreateCircle(circle cmdb.WorkFlowCircle) error {
	if !errors.Is(global.GVA_DB.Where("circle_name = ? ", circle.CircleName).First(&cmdb.WorkFlowCircle{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同状态名称 ")
	}
	return global.GVA_DB.Create(&circle).Error
}

func (w WorkFlowTemplateService) UpdateCircle(circle cmdb.WorkFlowCircle) error {
	var entity cmdb.WorkFlowCircle
	err := global.GVA_DB.Where("id = ?", circle.ID).First(&entity).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {                          // api记录不存在
		return err
	}
	err = global.GVA_DB.Model(&entity).Updates(&circle).Error
	if err != nil {
		return err
	}
	return nil
}

func (w WorkFlowTemplateService) DeleteCircle(circle cmdb.WorkFlowCircle) error {
	var entity cmdb.WorkFlowCircle
	err := global.GVA_DB.Where("id = ?", circle.ID).First(&entity).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {                          // api记录不存在
		return err
	}
	err = global.GVA_DB.Scopes().Delete(&entity).Error
	if err != nil {
		return err
	}
	return nil
}

func (w WorkFlowTemplateService) GetCircleById(circle cmdb.WorkFlowCircle) (res cmdb.WorkFlowCircle, err error) {
	err = global.GVA_DB.Where("id = ?", circle.ID).First(&res).Error
	return
}
