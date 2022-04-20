package advertising

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"
)

func Get(idstr string) (advertising Advertising) {
	database.DB.Preload("AdvertisingPosition").Preload("User").Where("id", idstr).First(&advertising)
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
	database.DB.Preload("advertising_position").Preload("User").Find(&advertisings)
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
		app.V1URL(database.TableName(&Advertising{})),
		perPage,
	)
	return
}

func Paginate2(c *gin.Context, perPage int, params string) (materials []Advertising, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		//database.DB.Model(Material{}),
		database.DB.Model(Advertising{}).Where("id like ?", "%"+params+"%").
			//Or("creator_id like ?", "%"+params+"%").
			//Or("advertising_no like ?", "%"+params+"%").
			//Or("department_id like ?", "%"+params+"%").
			Or("title like ?", "%"+params+"%"),
			//Or("type like ?", "%"+params+"%").
			//Or("material_id like ?", "%"+params+"%").
			//Or("material_type like ?", "%"+params+"%").
			//Or("size like ?", "%"+params+"%").
			//Or("redirect_to like ?", "%"+params+"%").
			//Or("redirect_params like ?", "%"+params+"%").
			//Or("description like ?", "%"+params+"%"),

		&materials,
		app.V1URL(database.TableName(&Advertising{})+"/list?params="+params),
		perPage,
	)
	return materials, paging
}

