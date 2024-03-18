package cmdb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type WorkFlowTemplateService struct {
}

func (w WorkFlowTemplateService) CreateWorkFlowTemplate() {

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
