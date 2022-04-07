package advertising

import (
	"gohub/pkg/helpers"

	"github.com/golang-module/carbon/v2"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// func (advertising *Advertising) BeforeSave(tx *gorm.DB) (err error) {}
// func (advertising *Advertising) BeforeCreate(tx *gorm.DB) (err error) {}
// func (advertising *Advertising) AfterCreate(tx *gorm.DB) (err error) {}
// func (advertising *Advertising) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (advertising *Advertising) AfterUpdate(tx *gorm.DB) (err error) {}
// func (advertising *Advertising) AfterSave(tx *gorm.DB) (err error) {}
// func (advertising *Advertising) BeforeDelete(tx *gorm.DB) (err error) {}
// func (advertising *Advertising) AfterDelete(tx *gorm.DB) (err error) {}
// func (advertising *Advertising) AfterFind(tx *gorm.DB) (err error) {}

func (advertising *Advertising) BeforeCreate(tx *gorm.DB) (err error) {
	// println(carbon.Now().Format("YmdHis") + helpers.RandomNumber(6))
	advertising.AdvertisingNo = cast.ToUint64(carbon.Now().Format("ymdHis") + helpers.RandomNumber(6))
	return
}
