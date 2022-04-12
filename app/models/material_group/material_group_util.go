package material_group

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Get(idstr string) (materialGroup MaterialGroup) {
	database.DB.Where("id", idstr).First(&materialGroup)
	return
}

func GetBy(field, value string) (materialGroup MaterialGroup) {
	database.DB.Where("? = ?", field, value).First(&materialGroup)
	return
}

func All() (materialGroups []MaterialGroup) {
	database.DB.Find(&materialGroups)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(MaterialGroup{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (materialGroups []MaterialGroup, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(MaterialGroup{}),
		&materialGroups,
		app.V1URL(database.TableName(&MaterialGroup{})),
		perPage,
	)
	return
}

//素材分组多条件查询
func Search(c *gin.Context, perPage int) (materials []MaterialGroup, paging paginator.Paging) {
	var db *gorm.DB
	name := c.Query("name")
	start_time := c.Query("start_time")
	end_time := c.Query("end_time")
	db = database.DB.Model(MaterialGroup{})
	if start_time != "" && end_time != "" {
		db = database.DB.Model(MaterialGroup{}).Where("created_at BETWEEN ? AND ?", start_time, end_time)
	}
	if name != "" {
		db = database.DB.Model(MaterialGroup{}).Where("name like ? ", "%"+name+"%")
	}
	if start_time != "" && end_time != "" && name != "" {
		db = database.DB.Model(MaterialGroup{}).Where("name like ? AND created_at BETWEEN ? AND ?", "%"+name+"%", start_time, end_time)
	}
	paging = paginator.Paginate(
		c,
		db,
		&materials,
		app.V1URL(database.TableName(&MaterialGroup{})),
		perPage,
	)
	return

}
