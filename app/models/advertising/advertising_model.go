//Package advertising 模型
package advertising

import (
	"gohub/app/models"
	"gohub/app/models/advertising_position"
	"gohub/app/models/user"
	"gohub/pkg/database"
)

type Advertising struct {
	models.BaseModel

	AdvertisingNo         uint64 `gorm:"column:advertising_no"`
	AdvertisingPositionId uint64 `gorm:"column:advertising_position_id"`
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

	User                user.User                                `json:"user" gorm:"foreignkey:id"`
	AdvertisingPosition advertising_position.AdvertisingPosition `json:"advertising_position"`

	models.CommonTimestampsField
}

func (advertising *Advertising) Create() {
	database.DB.Create(&advertising)
}

func (advertising *Advertising) Save() (rowsAffected int64) {
	result := database.DB.Save(&advertising)
	return result.RowsAffected
}

func (advertising *Advertising) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&advertising)
	return result.RowsAffected
}

func (advertising *Advertising) BatchDelete(ids []int) (rowsAffected int64) {
	result := database.DB.Delete(&advertising, ids)
	return result.RowsAffected
}
