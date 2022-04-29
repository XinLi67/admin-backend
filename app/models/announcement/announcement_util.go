package announcement

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/helpers"
	"gohub/pkg/paginator"
	"github.com/gin-gonic/gin"
)

func Get(idstr string) (announcement Announcement) {
	database.DB.Preload("User").Preload("AnnouncementPosition").Preload("Channel").Where("id", idstr).First(&announcement)
	return
}

func GetCreatorIdByname(name string) (result int) {
	database.DB.Table("users").Select("id").Where("name = ?", name).Find(&result)
	return
}

func GetBy(field, value string) (announcement Announcement) {
	database.DB.Where("? = ?", field, value).First(&announcement)
	return
}

func All() (announcements []Announcement) {
	// database.DB.Find(&announcements)
	database.DB.Preload("User").Preload("AnnouncementPosition").Preload("Channel").Find(&announcements)
	return
}

func All2() (announcements []*Announcement) {
	// database.DB.Find(&advertisings)
	database.DB.Find(&announcements)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Announcement{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (announcements []Announcement, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Announcement{}),
		&announcements,
		app.V1URL(database.TableName(&Announcement{})),
		perPage,
	)
	return
}

//公告列表查询
func Search(c *gin.Context, perPage int, user_id int) (announcements []Announcement, paging paginator.Paging) {

	start_date := c.Query("start_date")
	end_date := c.Query("end_date")
	title := c.Query("title")
	status := c.Query("status")
	db := database.DB.Model(Announcement{})
	if !helpers.Empty(start_date) && helpers.Empty(end_date) {
		db.Where("start_date BETWEEN ? AND ? ", start_date, end_date)
	}
	if !helpers.Empty(title) {
		db.Where("title like ? ", "%"+title+"%")
	}
	if !helpers.Empty(status) {
		db.Where("status = ?", status)
	}
	if !helpers.Empty(user_id) {
		db.Where("user_id = ?", user_id)
	}
	paging = paginator.Paginate(
		c,
		db,
		&announcements,
		app.V1URL(database.TableName(&Announcement{})),
		perPage,
	)
	return

}
