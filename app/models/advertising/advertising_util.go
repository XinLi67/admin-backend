package advertising

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Get(idstr string) (advertising Advertising) {
	// database.DB.Where("id", idstr).First(&advertising)
	database.DB.Preload("User").Where("id", idstr).First(&advertising)
	return
}

func GetBy(field, value string) (advertising Advertising) {
	database.DB.Where("? = ?", field, value).First(&advertising)
	return
}

func All() (advertisings []Advertising) {
	// database.DB.Find(&advertisings)
	database.DB.Preload("AdvertisingPosition").Preload("User").Find(&advertisings)
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

// func Search(c *gin.Context, perPage int) (advertisings []Advertising, paging paginator.Paging) {
// 	query := database.DB.Model(Advertising{}).Where("title like ?", "%"+c.Query("title")+"%")
// 	paging = paginator.Paginate(
// 		c,
// 		query,
// 		&advertisings,
// 		app.V1URL(database.TableName(&Advertising{})),
// 		perPage,
// 	)
// 	return
// }

func Search(c *gin.Context, perPage int) (Advertisings []Advertising, paging paginator.Paging) {
	var db *gorm.DB
	types := c.Query("type")
	advertising_position_id := c.Query("advertising_position_id")
	title := c.Query("title")
	if types != "" {
		db = database.DB.Model(Advertising{}).Where("type = ? ", types)
		if advertising_position_id != "" {
			db = database.DB.Model(Advertising{}).Where("type = ? AND advertising_position_id= ?", types, advertising_position_id)
		}
		if title != "" {
			if advertising_position_id != "" {
				db = database.DB.Model(Advertising{}).Where("type = ? AND advertising_position_id= ? AND title like ?", types, advertising_position_id, "%"+title+"%")
			} else {
				db = database.DB.Model(Advertising{}).Where("type = ? AND title like ?", types, "%"+title+"%")
			}
		}
		paging = paginator.Paginate(
			c,
			db,
			&Advertisings,
			app.V1URL(database.TableName(&Advertising{})),
			perPage,
		)
		return
	} else {
		paging = paginator.Paginate(
			c,
			database.DB.Model(Advertising{}),
			&Advertisings,
			app.V1URL(database.TableName(&Advertising{})),
			perPage,
		)
		return
	}
}
