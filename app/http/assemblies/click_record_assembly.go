package assemblies

import (
	"gohub/app/models/click_record"

	"github.com/golang-module/carbon/v2"
)

type ClickRecordAssembly struct {
	ID            uint64 `json:"id"`
	AdvertisingId uint64 `json:"advertising_id"`
	CustomerId    uint64 `json:"customer_id"`
	BrowsingTime  uint64 `json:"browsing_time"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`

	Advertising *AdvertisingAssembly `json:"advertising"`
}

func ClickRecordAssemblyFromModel(data click_record.ClickRecord) *ClickRecordAssembly {
	var advertisingAssembly *AdvertisingAssembly
	if data.Advertising != nil {
		advertisingAssembly = &AdvertisingAssembly{
			ID:            data.Advertising.ID,
			AdvertisingNo: data.Advertising.AdvertisingNo,
			Title:         data.Advertising.Title,
			CreatedAt:     carbon.Time2Carbon(data.Advertising.CreatedAt).ToDateTimeString(),
			UpdatedAt:     carbon.Time2Carbon(data.Advertising.UpdatedAt).ToDateTimeString(),
		}
	}
	ClickRecord := &ClickRecordAssembly{
		ID:            data.ID,
		AdvertisingId: data.AdvertisingId,
		CustomerId:    data.CustomerId,
		BrowsingTime:  data.BrowsingTime,
		StartTime:     data.StartTime,
		EndTime:       data.EndTime,
		CreatedAt:     carbon.Time2Carbon(data.CreatedAt).ToDateTimeString(),
		UpdatedAt:     carbon.Time2Carbon(data.UpdatedAt).ToDateTimeString(),

		Advertising: advertisingAssembly,
	}
	return ClickRecord
}

func ClickRecordAssemblyFromModelList(data []click_record.ClickRecord) interface{} {
	ClickRecords := make([]ClickRecordAssembly, len(data))
	var advertisingAssembly *AdvertisingAssembly

	for i, v := range data {
		if v.Advertising != nil {
			advertisingAssembly = &AdvertisingAssembly{
				ID:            v.Advertising.ID,
				AdvertisingNo: v.Advertising.AdvertisingNo,
				Title:         v.Advertising.Title,
				CreatedAt:     carbon.Time2Carbon(v.Advertising.CreatedAt).ToDateTimeString(),
				UpdatedAt:     carbon.Time2Carbon(v.Advertising.UpdatedAt).ToDateTimeString(),
			}
		}
		ClickRecords[i] = ClickRecordAssembly{
			ID:            v.ID,
			AdvertisingId: v.AdvertisingId,
			CustomerId:    v.CustomerId,
			BrowsingTime:  v.BrowsingTime,
			StartTime:     v.StartTime,
			EndTime:       v.EndTime,
			CreatedAt:     carbon.Time2Carbon(v.CreatedAt).ToDateTimeString(),
			UpdatedAt:     carbon.Time2Carbon(v.UpdatedAt).ToDateTimeString(),

			Advertising: advertisingAssembly,
		}
	}

	return ClickRecords
}
