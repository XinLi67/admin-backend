package channel

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"
	"gohub/pkg/helpers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Get(idstr string) (channel Channel) {
	database.DB.Where("id", idstr).First(&channel)
	return
}

func GetBy(field, value string) (channel Channel) {
	database.DB.Where("? = ?", field, value).First(&channel)
	return
}

func All() (channels []Channel) {
	database.DB.Find(&channels)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Channel{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (channels []Channel, paging paginator.Paging) {
	var db *gorm.DB
	name := c.Query("name")
	db = database.DB.Model(Channel{})
	if !helpers.Empty(name) {
		db.Where("name like ? ", "%"+name+"%")
	}
	paging = paginator.Paginate(
		c,
		db,
		&channels,
		app.V1URL(database.TableName(&Channel{})),
		perPage,
	)
	return
}

// func (channel *Channel) BatchDelete(ids []int) (rowsAffected int64) {
// 	result := database.DB.Delete(&channel, ids)
// 	return result.RowsAffected
// }
