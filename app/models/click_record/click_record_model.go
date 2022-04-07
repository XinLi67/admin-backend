//Package click_record 模型
package click_record

import (
	"gohub/app/models"
	"gohub/pkg/database"
	"time"
)

type ClickRecord struct {
	models.BaseModel

	AdvertisingId uint64    `gorm:"column:advertising_id"`
	CustomerId    uint64    `gorm:"column:customer_id"`
	BrowsingTime  uint64    `gorm:"column:browsing_time"`
	StartTime     time.Time `gorm:"column:start_time"`
	EndTime       time.Time `gorm:"column:end_time"`

	models.CommonTimestampsField
}

func (clickRecord *ClickRecord) Create() {
	database.DB.Create(&clickRecord)
}

func (clickRecord *ClickRecord) Save() (rowsAffected int64) {
	result := database.DB.Save(&clickRecord)
	return result.RowsAffected
}

func (clickRecord *ClickRecord) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&clickRecord)
	return result.RowsAffected
}

func (clickRecord *ClickRecord) BatchDelete(ids []int) (rowsAffected int64) {
	result := database.DB.Delete(&clickRecord, ids)
	return result.RowsAffected
}
