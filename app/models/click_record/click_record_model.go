//Package click_record 模型
package click_record

import (
	"gohub/app/models"
	"gohub/app/models/advertising"
	"gohub/pkg/database"
)

type ClickRecord struct {
	models.BaseModel

	AdvertisingId uint64 `gorm:"column:advertising_id"`
	CustomerId    uint64 `gorm:"column:customer_id"`
	PositionId    uint64 `gorm:"column:position_id"`
	BrowsingTime  uint64 `gorm:"column:browsing_time"`
	StartTime     string `gorm:"column:start_time"`
	EndTime       string `gorm:"column:end_time"`

	Advertising advertising.Advertising `json:"advertising"`

	models.CommonTimestampsField
}

type DateGroupResult struct {
	Mon  string
	Week  string
	Total int64
}

//type ClickRecordGroupResult struct {
//	AdvertisingId  string
//	CustomerId  string
//	PositionId  string  `gorm:"column:advertising_position_id"`
//	Total int64
//}

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
