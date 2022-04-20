package advertising

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"
)

func Get(idstr string) (advertising Advertising) {
	// database.DB.Where("id", idstr).First(&advertising)
	database.DB.Preload("advertising_position").Preload("User").Where("id", idstr).First(&advertising)
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
		database.DB.Model(Advertising{}).Where("advertising_position_id = ?", params),
		&materials,
		app.V1URL(database.TableName(&Advertising{})+"/list?params="+params),
		perPage,
	)
	return materials, paging
}

