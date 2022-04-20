package advertising_plan

import (
	"gohub/pkg/app"
	"gohub/pkg/cache"
	"gohub/pkg/database"
	"gohub/pkg/helpers"
	"gohub/pkg/paginator"
	"time"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (advertisingPlan AdvertisingPlan) {
	// database.DB.Where("id", idstr).First(&advertisingPlan)
	// database.DB.Preload("User").Preload("AdvertisingPosition").Where("id", idstr).First(&advertisingPlan)
	database.DB.Preload("User").Where("id", idstr).First(&advertisingPlan)
	return
}

//根据广告位获取对应广告
func GetAll(idstr string) (advertisingPlan []AdvertisingPlan) {
	// database.DB.Where("id", idstr).First(&advertising)
	database.DB.Preload("advertising_position").Joins("LEFT JOIN advertisings  on advertising_plans.advertising_id=advertisings.id").Where("advertising_plans.advertising_position_id", idstr).Find(&advertisingPlan)
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

func All2() (advertisingPlans []*AdvertisingPlan) {
	// database.DB.Find(&advertisings)
	database.DB.Find(&advertisingPlans)
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

//根据审核状态查询后分页显示
func Paginate2(c *gin.Context, perPage int, audit_status string) (advertisingPlans []AdvertisingPlan, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(AdvertisingPlan{}).Where("audit_status = ?", audit_status),
		&advertisingPlans,
		app.V1URL(database.TableName(&AdvertisingPlan{})+"/list?audit_status="+audit_status),
		perPage,
	)
	return advertisingPlans, paging
}

//缓存相关
func AllCached(id string) (advertisingPlans []AdvertisingPlan) {
	// 设置缓存 key
	cacheKey := "advertisings:all"
	// 设置过期时间
	expireTime := 120 * time.Minute
	// 取数据
	cache.GetObject(cacheKey, &advertisingPlans)

	// 如果数据为空
	if helpers.Empty(advertisingPlans) {
		// 查询数据库
		advertisingPlans = GetAll(id)
		if helpers.Empty(advertisingPlans) {
			return advertisingPlans
		}
		// 设置缓存
		cache.Set(cacheKey, advertisingPlans, expireTime)
	}
	return
}

func GetByStatus(c *gin.Context, perPage int, audit_status string) (advertisingPlans []AdvertisingPlan, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(AdvertisingPlan{}).Where("audit_status = ?", audit_status),
		&advertisingPlans,
		app.V1URL(database.TableName(&AdvertisingPlan{})+"/list?audit_status="+audit_status),
		perPage,
	)
	return advertisingPlans, paging
}