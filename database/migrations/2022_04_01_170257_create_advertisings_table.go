package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Advertising struct {
		models.BaseModel

		AdvertisingNo         uint64 `gorm:"column:advertising_no"`
		AdvertisingPositionId uint64 `gorm:"column:advertising_position_id"`
		AdvertisingPlanId     uint64 `gorm:"column:advertising_plan_id"`
		CreatorId             uint64 `gorm:"column:creator_id"`
		DepartmentId          uint64 `gorm:"column:department_id"`
		Title                 string `gorm:"column:title"`
		Type                  uint64 `gorm:"column:type"`
		RedirectTo            uint64 `gorm:"column:redirect_to"`
		MaterialId            uint64 `gorm:"column:material_id"`
		MaterialType          uint64 `gorm:"column:material_type"`
		Size                  string `gorm:"column:size"`
		RedirectParams        string `gorm:"column:redirect_params"`
		Description           string `gorm:"column:description"`
		Status                uint64 `gorm:"column:status"`
		AuditReason           string `gorm:"column:audit_reason"`
		PushContent           string `gorm:"column:push_content"`
		PushTitle             string `gorm:"column:push_title"`
		AdvertisingCreativity string `gorm:"column:advertising_creativity"`
		StartTime             string `gorm:"type:varchar(20);column:start_time"`
		EndTime               string `gorm:"type:varchar(20);column:end_time"`
		SchedulingTime        uint64 `gorm:"column:scheduling_time"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Advertising{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Advertising{})
	}

	migrate.Add("2022_04_01_170257_create_advertisings_table", up, down)
}
