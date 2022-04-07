package seeders

import (
	"fmt"
	"gohub/app/models/audit_record"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gohub/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedAuditRecordsTable", func(db *gorm.DB) {

		auditRecords := []audit_record.AuditRecord{
			{
				AuditableId:   1,
				AuditableType: 1,
				ApplicantId:   1,
				AuditorId:     2,
				Status:        1,
				Content:       "同意",
			},
			{
				AuditableId:   1,
				AuditableType: 1,
				ApplicantId:   1,
				AuditorId:     2,
				Status:        1,
				Content:       "驳回",
			},
			{
				AuditableId:   1,
				AuditableType: 2,
				ApplicantId:   1,
				AuditorId:     2,
				Status:        1,
				Content:       "同意",
			},
			{
				AuditableId:   1,
				AuditableType: 2,
				ApplicantId:   1,
				AuditorId:     2,
				Status:        1,
				Content:       "驳回",
			},
		}

		result := db.Table("audit_records").Create(&auditRecords)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
