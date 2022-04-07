//Package audit_record 模型
package audit_record

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

type AuditRecord struct {
	models.BaseModel

	AuditableId   uint64 `gorm:"column:auditable_id"`
	AuditableType uint64 `gorm:"column:auditable_type"`
	ApplicantId   uint64 `gorm:"column:applicant_id"`
	AuditorId     uint64 `gorm:"column:auditor_id"`
	Status        uint64 `gorm:"column:status"`
	Content       string `gorm:"column:content"`

	models.CommonTimestampsField
}

func (auditRecord *AuditRecord) Create() {
	database.DB.Create(&auditRecord)
}

func (auditRecord *AuditRecord) Save() (rowsAffected int64) {
	result := database.DB.Save(&auditRecord)
	return result.RowsAffected
}

func (auditRecord *AuditRecord) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&auditRecord)
	return result.RowsAffected
}

func (auditRecord *AuditRecord) BatchDelete(ids []int) (rowsAffected int64) {
	result := database.DB.Delete(&auditRecord, ids)
	return result.RowsAffected
}
