package announcement

import (
	"gohub/pkg/helpers"

	"github.com/golang-module/carbon/v2"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// func (announcement *Announcement) BeforeSave(tx *gorm.DB) (err error) {}
// func (announcement *Announcement) BeforeCreate(tx *gorm.DB) (err error) {}
// func (announcement *Announcement) AfterCreate(tx *gorm.DB) (err error) {}
// func (announcement *Announcement) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (announcement *Announcement) AfterUpdate(tx *gorm.DB) (err error) {}
// func (announcement *Announcement) AfterSave(tx *gorm.DB) (err error) {}
// func (announcement *Announcement) BeforeDelete(tx *gorm.DB) (err error) {}
// func (announcement *Announcement) AfterDelete(tx *gorm.DB) (err error) {}
// func (announcement *Announcement) AfterFind(tx *gorm.DB) (err error) {}

func (announcement *Announcement) BeforeCreate(tx *gorm.DB) (err error) {
	// println(carbon.Now().Format("YmdHis") + helpers.RandomNumber(6))
	announcement.AnnouncementNo = cast.ToUint64(carbon.Now().Format("ymdHis") + helpers.RandomNumber(6))
	return
}
