package announcement_plan

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (announcementPlan AnnouncementPlan) {
	// database.DB.Where("id", idstr).First(&announcementPlan)
	// database.DB.Preload("User").Preload("AnnouncementPosition").Where("id", idstr).First(&announcementPlan)
	database.DB.Preload("User").Where("id", idstr).First(&announcementPlan)
	return
}

func GetBy(field, value string) (announcementPlan AnnouncementPlan) {
	database.DB.Where("? = ?", field, value).First(&announcementPlan)
	return
}

func All() (announcementPlans []AnnouncementPlan) {
	// database.DB.Find(&announcementPlans)
	database.DB.Preload("User").Preload("AnnouncementPosition").Find(&announcementPlans)
	return
}

func All2() (announcementPlans []*AnnouncementPlan) {
	// database.DB.Find(&advertisings)
	database.DB.Find(&announcementPlans)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(AnnouncementPlan{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (announcementPlans []AnnouncementPlan, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(AnnouncementPlan{}),
		&announcementPlans,
		app.V1URL(database.TableName(&AnnouncementPlan{})),
		perPage,
	)
	return
}

func Paginate2(c *gin.Context, perPage int) (announcementPlans []AnnouncementPlan, paging paginator.Paging) {

	var db *gorm.DB
	name := c.Query("name")
	audit_status := c.Query("audit_status")

	db = database.DB.Model(AnnouncementPlan{}).Where(" id like ?", "%"+"%")

	if len(name) >0{
		db.Where("name like ? ","%"+name+"%")
	}

	if len(audit_status) >0{
		db.Where("status = ? ",audit_status)
	}

	paging = paginator.Paginate(
		c,
		database.DB.Model(AnnouncementPlan{}),
		&announcementPlans,
		app.V1URL(database.TableName(&AnnouncementPlan{})),
		perPage,
	)
	return

}
