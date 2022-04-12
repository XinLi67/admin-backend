package material

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

//素材多条件查询
func Search(c *gin.Context, perPage int) (materials []Material, paging paginator.Paging) {
	var db *gorm.DB
	title := c.Query("title")
	start_time := c.Query("start_time")
	end_time := c.Query("end_time")
	material_group_id := c.Query("material_group_id")
	db = database.DB.Model(Material{}).Where(" material_group_id = ?", material_group_id)
	if start_time != "" && end_time != "" {
		db = database.DB.Model(Material{}).Where("title like ? AND material_group_id = ?", "%"+title+"%", material_group_id)
	}
	if title != "" {
		db = database.DB.Model(Material{}).Where("created_at BETWEEN ? AND ? AND material_group_id = ?", start_time, end_time, material_group_id)
	}
	if start_time != "" && end_time != "" && title != "" {
		db = database.DB.Model(Material{}).Where("title like ? AND created_at BETWEEN ? AND ? AND material_group_id = ?", "%"+title+"%", start_time, end_time, material_group_id)
	}
	paging = paginator.Paginate(
		c,
		db,
		&materials,
		app.V1URL(database.TableName(&Material{})),
		perPage,
	)
	return

}
