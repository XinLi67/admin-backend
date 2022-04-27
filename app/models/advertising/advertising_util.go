package advertising

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"
	"gorm.io/gorm"
)

func Get(idstr string) (advertising Advertising) {
	database.DB.Preload("AdvertisingPosition").Preload("User").Preload("Channel").Where("id", idstr).First(&advertising)
	return
}

//根据广告位获取对应广告
func GetAll(idstr string) (advertising []Advertising) {
	// database.DB.Where("id", idstr).First(&advertising)
	database.DB.Preload("advertising_position").Where("advertising_position_id", idstr).Find(&advertising)
	return
}

func GetBy(field, value string) (advertising Advertising) {
	database.DB.Where("? = ?", field, value).First(&advertising)
	return
}

func All() (advertisings []Advertising) {
	// database.DB.Find(&advertisings)
	database.DB.Preload("advertising_position").Preload("User").Preload("channel").Find(&advertisings)
	return
}

func All2() (advertisings []*Advertising) {
	// database.DB.Find(&advertisings)
	database.DB.Find(&advertisings)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Advertising{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (advertisings []Advertising, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Advertising{}),
		&advertisings,
		app.V1URL("advertising"),
		perPage,
	)
	return
}

func Paginate2(c *gin.Context, perPage int) (advertisings []Advertising, paging paginator.Paging) {

	var db *gorm.DB
	title := c.Query("title")
	adtype := c.Query("type")
	posId := c.Query("advertising_position_id")
	status := c.Query("status")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")
	creator_name := c.Query("creator_name")

	//db = database.DB.Model(Advertising{}).
	//	Joins(" left JOIN advertising_plans on advertisings.advertising_plan_id=advertising_plans.id UNION SELECT  * FROM advertisings RIGHT JOIN advertising_plans ON advertisings.advertising_plan_id = advertising_plans.id").
	//	Where(" advertisings.id like ?", "%"+"%")

	if len(start_date) >0 && len(end_date)>0{
		db = database.DB.Model(Advertising{}).
			Joins(" left JOIN advertising_plans on advertisings.advertising_plan_id=advertising_plans.id UNION SELECT  * FROM advertisings RIGHT JOIN advertising_plans ON advertisings.advertising_plan_id = advertising_plans.id").
			Where(" advertisings.id like ?", "%"+"%")
		db.Where("DATE_FORMAT(advertising_plans.start_date,'%Y-%m-%d') >=  ? ",start_date)
		db.Where("DATE_FORMAT(advertising_plans.end_date,'%Y-%m-%d') <=  ? ",end_date)
	}else{
		db = database.DB.Model(Advertising{}).Where(" id like ?", "%"+"%")
	}

	if len(title) >0{
		db.Where("title like ? ","%"+title+"%")
	}

	if len(adtype) >0{
		db.Where("type = ? ",adtype)
	}

	if len(posId) >0{
		db.Where("advertising_position_id = ? ",adtype)
	}

	if len(status) >0{
		db.Where("status = ? ",status)
	}

	if len(creator_name) >0{
		db.Where("creator_id in(SELECT id from users WHERE `name` like ?) ","%"+creator_name+"%")
	}

	paging = paginator.Paginate(
		c,
		db,
		&advertisings,
		app.V1URL("advertising"),
		perPage,
	)
	return

}