package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type AuditRecord struct {
		models.BaseModel

		AuditableId   uint64 `gorm:"type:tinyint unsigned;column:auditable_id"`
		AuditableType uint64 `gorm:"type:tinyint unsigned;column:auditable_type"`
		ApplicantId   uint64 `gorm:"type:bigint unsigned;column:applicant_id"`
		AuditorId     uint64 `gorm:"type:bigint unsigned;column:auditor_id"`
		Status        uint64 `gorm:"type:tinyint unsigned;column:status"`
		Content       string `gorm:"type:text;column:content"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&AuditRecord{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&AuditRecord{})
	}

	migrate.Add("2022_04_01_170417_create_audit_records_table", up, down)
}
