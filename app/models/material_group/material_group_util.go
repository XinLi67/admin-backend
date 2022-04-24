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
func GetCountById(id string) (count int64) {
	database.DB.Model(MaterialGroup{}).Where(" parent_id = ?", id).Count(&count)
	return
}
//文件夹多条件查找
func GetDocumentById(c *gin.Context, perPage int, id string) (materialGroups []MaterialGroup, paging paginator.Paging) {
	var db *gorm.DB
	name := c.Query("name")
	start_time := c.Query("start_time")
	end_time := c.Query("end_time")
	db = database.DB.Model(MaterialGroup{}).Where(" parent_id = ?", id)
	if start_time != "" && end_time != "" {
		db.Where("created_at BETWEEN ? AND ? ", start_time, end_time)
	}
	if name != "" {
		db.Where("name like ? ", "%"+name+"%")
	}
	paging = paginator.Paginate(
		c,
		db,
		&materialGroups,
		app.V1URL(database.TableName(&MaterialGroup{})),
		perPage,
	)
	return
}

