package announcement_position

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (announcementPosition AnnouncementPosition) {
	database.DB.Where("id", idstr).First(&announcementPosition)
	// database.DB.Preload("Channel").Where("id", idstr).First(&announcementPosition)
	return
}

func GetBy(field, value string) (announcementPosition AnnouncementPosition) {
	database.DB.Where("? = ?", field, value).First(&announcementPosition)
	return
}

func All() (announcementPositions []AnnouncementPosition) {
	// database.DB.Find(&announcementPositions)
	database.DB.Preload("Channel").Find(&announcementPositions)
	return
}

func All2() (announcementPositions []*AnnouncementPosition) {
	// database.DB.Find(&advertisings)
	database.DB.Find(&announcementPositions)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(AnnouncementPosition{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (announcementPositions []AnnouncementPosition, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(AnnouncementPosition{}),
		&announcementPositions,
		app.V1URL(database.TableName(&AnnouncementPosition{})),
		perPage,
	)
	return
}
