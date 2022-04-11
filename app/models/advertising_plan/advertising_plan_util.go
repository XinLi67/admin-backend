package advertising_plan

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (advertisingPlan AdvertisingPlan) {
	// database.DB.Where("id", idstr).First(&advertisingPlan)
	// database.DB.Preload("User").Preload("AdvertisingPosition").Where("id", idstr).First(&advertisingPlan)
	database.DB.Preload("User").Where("id", idstr).First(&advertisingPlan)
	return
}

func GetBy(field, value string) (advertisingPlan AdvertisingPlan) {
	database.DB.Where("? = ?", field, value).First(&advertisingPlan)
	return
}

func All() (advertisingPlans []AdvertisingPlan) {
	// database.DB.Find(&advertisingPlans)
	database.DB.Preload("User").Preload("AdvertisingPosition").Find(&advertisingPlans)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(AdvertisingPlan{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (advertisingPlans []AdvertisingPlan, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(AdvertisingPlan{}),
		&advertisingPlans,
		app.V1URL(database.TableName(&AdvertisingPlan{})),
		perPage,
	)
	return
}
