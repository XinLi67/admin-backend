package advertising_position

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

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

func PaginateByName(c *gin.Context, perPage int, params string) (advertisingPositions []AdvertisingPosition, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(AdvertisingPosition{}).Where("name like ?", "%"+params+"%"),
		&advertisingPositions,
		app.V1URL(database.TableName(&AdvertisingPosition{})+"/list?params="+params),
		perPage,
	)
	return advertisingPositions, paging
}

//根据审核状态查询后分页显示
func PaginateByStatus(c *gin.Context, perPage int, status string) (advertisingPositions []AdvertisingPosition, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(AdvertisingPosition{}).Where("status = ?", status),
		&advertisingPositions,
		app.V1URL(database.TableName(&AdvertisingPosition{})+"/list?status="+status),
		perPage,
	)
	return advertisingPositions, paging
}

//根据审核状态查询后分页显示
func PaginateByStatusAndParams(c *gin.Context, perPage int, status string, params string) (advertisingPlans []AdvertisingPosition, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(AdvertisingPosition{}).Where("status = ? and name like ?", status, "%"+params+"%"),
		&advertisingPlans,
		app.V1URL(database.TableName(&AdvertisingPosition{})+"/list?status="+status+"?params="+params),
		perPage,
	)
	return advertisingPlans, paging
}
