package click_record

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (clickRecord ClickRecord) {
	database.DB.Preload("Advertising").Where("id", idstr).First(&clickRecord)
	return
}

func GetBy(field, value string) (clickRecord ClickRecord) {
	database.DB.Where("? = ?", field, value).First(&clickRecord)
	return
}

func All() (clickRecords []ClickRecord) {
	database.DB.Preload("Advertising").Find(&clickRecords)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(ClickRecord{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (clickRecords []ClickRecord, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(ClickRecord{}),
		&clickRecords,
		app.V1URL(database.TableName(&ClickRecord{})),
		perPage,
	)
	return
}
