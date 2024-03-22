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
	if !errors.Is(global.GVA_DB.Where("flow_name = ? ", workflow.FlowName).First(&cmdb.WorkFlowTemplate{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同流程名称 ")
	}
	return global.GVA_DB.Create(&workflow).Error
}

func (w WorkFlowTemplateService) GetWorkFlowTemplateList(workflow cmdb.WorkFlowTemplate, info request.PageInfo) (templateList []cmdb.WorkFlowTemplate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&cmdb.WorkFlowTemplate{})

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
