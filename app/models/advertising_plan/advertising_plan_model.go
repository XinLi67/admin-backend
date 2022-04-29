//Package advertising_plan 模型
package advertising_plan

import (
	"gohub/app/models"
	"gohub/app/models/advertising_position"
	"gohub/app/models/user"
	"gohub/pkg/database"
)

type AdvertisingPlan struct {
	models.BaseModel

	Name                  string `gorm:"column:name"`
	CreatorId             uint64 `gorm:"column:creator_id"`
	AdvertisingId         uint64 `gorm:"column:advertising_id"`
	AdvertisingType       uint64 `gorm:"column:advertising_type"`
	AdvertisingPositionId uint64 `gorm:"column:advertising_position_id"`
	Order                 uint64 `gorm:"column:order"`
	SchedulingDate        uint64 `gorm:"column:scheduling_date"`
	SchedulingTime        uint64 `gorm:"column:scheduling_time"`
	StartDate             string `gorm:"column:start_date"`
	EndDate               string `gorm:"column:end_date"`
	StartTime             string `gorm:"column:start_time"`
	EndTime               string `gorm:"column:end_time"`
	AuditStatus           uint64 `gorm:"column:audit_status"`
	PresentStatus         uint64 `gorm:"column:present_status"`

	User                *user.User                                `json:"user" gorm:"foreignkey:id;references:CreatorId"`
	AdvertisingPosition *advertising_position.AdvertisingPosition `json:"advertising_position"`
	//Advertising         advertising.Advertising                `json:"advertising"`

	models.CommonTimestampsField
}

func (advertisingPlan *AdvertisingPlan) Create() {
	database.DB.Create(&advertisingPlan)
}

func (advertisingPlan *AdvertisingPlan) Save() (rowsAffected int64) {
	result := database.DB.Save(&advertisingPlan)
	return result.RowsAffected
}

func (advertisingPlan *AdvertisingPlan) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&advertisingPlan)
	return result.RowsAffected
}

func (advertisingPlan *AdvertisingPlan) BatchDelete(ids []int) (rowsAffected int64) {
	result := database.DB.Delete(&advertisingPlan, ids)
	return result.RowsAffected
}

func (advertisingPlan *AdvertisingPlan) BatchStore(adPlans []AdvertisingPlan) (rowsAffected int64) {
	result := database.DB.Create(&adPlans)
	return result.RowsAffected
}
