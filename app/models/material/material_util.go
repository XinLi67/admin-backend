package material

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (material Material) {
	database.DB.Preload("MaterialGroup").Where("id", idstr).First(&material)
	return
}

func GetBy(field, value string) (material Material) {
	database.DB.Where("? = ?", field, value).First(&material)
	return
}

func All() (materials []Material) {
	database.DB.Preload("MaterialGroup").Find(&materials)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Material{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (materials []Material, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Material{}),
		&materials,
		app.V1URL(database.TableName(&Material{})),
		perPage,
	)
	return
}
