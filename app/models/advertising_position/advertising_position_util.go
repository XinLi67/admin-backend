package advertising_position

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (advertisingPosition AdvertisingPosition) {
	database.DB.Where("id", idstr).First(&advertisingPosition)
	// database.DB.Preload("Channel").Where("id", idstr).First(&advertisingPosition)
	return
}

func GetBy(field, value string) (advertisingPosition AdvertisingPosition) {
	database.DB.Where("? = ?", field, value).First(&advertisingPosition)
	return
}

func All() (advertisingPositions []AdvertisingPosition) {
	// database.DB.Find(&advertisingPositions)
	database.DB.Preload("Channel").Find(&advertisingPositions)
	return
}

func All2() (advertisingPos []*AdvertisingPosition) {
	// database.DB.Find(&advertisings)
	database.DB.Find(&advertisingPos)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(AdvertisingPosition{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (advertisingPositions []AdvertisingPosition, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(AdvertisingPosition{}),
		&advertisingPositions,
		app.V1URL(database.TableName(&AdvertisingPosition{})),
		perPage,
	)
	return
}

func Paginate2(c *gin.Context, perPage int) (advertisingPositions []AdvertisingPosition, paging paginator.Paging) {
	status := c.Query("status")
	name := c.Query("name")

	var db *gorm.DB
	db = database.DB.Model(AdvertisingPosition{}).Where(" id like ?", "%"+"%")


	if len(status) >0{
		db.Where("status like ? ","%"+status+"%")
	}

	if len(status) >0{
		db.Where("name like ? ","%"+name+"%")
	}

	paging = paginator.Paginate(
		c,
		db,
		&advertisingPositions,
		app.V1URL(database.TableName(&AdvertisingPosition{})),
		perPage,
	)
	return

}
