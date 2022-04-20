package migrations

import (
    "database/sql"
    "gohub/app/models"
    "gohub/pkg/migrate"

    "gorm.io/gorm"
)

func init() {

    type ClickRecordGroupResult struct {
        models.BaseModel

        AdvertisingId string `gorm:"column:advertising_id"`
        CustomerId    string `gorm:"column:customer_id"`
        PositionId    string `gorm:"column:advertising_position_id"`
        Total         int64  `gorm:"column:total"`

        models.CommonTimestampsField
    }

    up := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.AutoMigrate(&ClickRecordGroupResult{})
    }

    down := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.DropTable(&ClickRecordGroupResult{})
    }

    migrate.Add("2022_04_13_160152_add_click_record_group_results_table", up, down)
}