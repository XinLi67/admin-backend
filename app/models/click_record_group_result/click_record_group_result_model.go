//Package click_record_group_result 模型
package click_record_group_result

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

type ClickRecordGroupResult struct {
	models.BaseModel

	// Put fields in here
	AdvertisingId         string `gorm:"column:advertising_id"`
	CustomerId            string `gorm:"column:customer_id"`
	AdvertisingPositionId string `gorm:"column:advertising_position_id"`
	Total                 int64  `gorm:"column:total"`

	models.CommonTimestampsField
}

func (clickRecordGroupResult *ClickRecordGroupResult) Create() {
	database.DB.Create(&clickRecordGroupResult)
}

func (clickRecordGroupResult *ClickRecordGroupResult) Save() (rowsAffected int64) {
	result := database.DB.Save(&clickRecordGroupResult)
	return result.RowsAffected
}

func (clickRecordGroupResult *ClickRecordGroupResult) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&clickRecordGroupResult)
	return result.RowsAffected
}

func (clickRecordGroupResult *ClickRecordGroupResult) BatchDelete(ids []int) (rowsAffected int64) {
	result := database.DB.Delete(&clickRecordGroupResult, ids)
	return result.RowsAffected
}
