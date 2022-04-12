package announcement

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Get(idstr string) (announcement Announcement) {
	// database.DB.Where("id", idstr).First(&announcement)
	database.DB.Preload("User").Where("id", idstr).First(&announcement)
	return
}

func GetBy(field, value string) (announcement Announcement) {
	database.DB.Where("? = ?", field, value).First(&announcement)
	return
}

func All() (announcements []Announcement) {
	// database.DB.Find(&announcements)
	database.DB.Preload("AnnouncementPosition").Preload("User").Find(&announcements)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Announcement{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

// func Paginate(c *gin.Context, perPage int) (announcements []Announcement, paging paginator.Paging) {
// 	paging = paginator.Paginate(
// 		c,
// 		database.DB.Model(Announcement{}),
// 		&announcements,
// 		app.V1URL(database.TableName(&Announcement{})),
// 		perPage,
// 	)
// 	return
// }

func Search(c *gin.Context, perPage int) (announcements []Announcement, paging paginator.Paging) {
	var db *gorm.DB
	types := c.Query("type")
	announcement_position_id := c.Query("announcement_position_id")
	title := c.Query("title")
	status := c.Query("status")
	if types != "" {
		db = database.DB.Model(Announcement{}).Where("type = ? ", types)
		if announcement_position_id != "" {
			db = database.DB.Model(Announcement{}).Where("type = ? AND announcement_position_id= ?", types, announcement_position_id)
		}
		if title != "" {
			if announcement_position_id != "" {
				db = database.DB.Model(Announcement{}).Where("type = ? AND announcement_position_id= ? AND title like ?", types, announcement_position_id, "%"+title+"%")
			} else {
				db = database.DB.Model(Announcement{}).Where("type = ? AND title like ?", types, "%"+title+"%")
			}
		}
		paging = paginator.Paginate(
			c,
			db,
			&announcements,
			app.V1URL(database.TableName(&Announcement{})),
			perPage,
		)
		return
	} else {
		if status != "" {
			db = database.DB.Model(Announcement{}).Where("status = ? ", status)
			paging = paginator.Paginate(
				c,
				db,
				&announcements,
				app.V1URL(database.TableName(&Announcement{})),
				perPage,
			)
		} else {
			paging = paginator.Paginate(
				c,
				database.DB.Model(Announcement{}),
				&announcements,
				app.V1URL(database.TableName(&Announcement{})),
				perPage,
			)
		}
		return
	}
}
