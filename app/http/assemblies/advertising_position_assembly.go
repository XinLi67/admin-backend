package assemblies

import (
	"gohub/app/models/advertising_position"

	"github.com/golang-module/carbon/v2"
)

type AdvertisingPositionAssembly struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	ChannelId   uint64 `json:"channel_id"`
	Code        string `json:"code"`
	Height      uint64 `json:"height"`
	Weight      uint64 `json:"weight"`
	Status      uint64 `json:"status"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`

	Channel *ChannelAssembly `json:"channel"`
}

func AdvertisingPositionAssemblyFromModel(data advertising_position.AdvertisingPosition) *AdvertisingPositionAssembly {
	var channelAssembly *ChannelAssembly
	if data.Channel != nil {
		channelAssembly = &ChannelAssembly{
			ID:        data.Channel.ID,
			Name:      data.Channel.Name,
			CreatedAt: carbon.Time2Carbon(data.Channel.CreatedAt).ToDateTimeString(),
			UpdatedAt: carbon.Time2Carbon(data.Channel.UpdatedAt).ToDateTimeString(),
		}
	}
	advertising := &AdvertisingPositionAssembly{
		ID:          data.ID,
		Name:        data.Name,
		ChannelId:   data.ChannelId,
		Code:        data.Code,
		Height:      data.Height,
		Weight:      data.Weight,
		Status:      data.Status,
		Description: data.Description,
		CreatedAt:   carbon.Time2Carbon(data.CreatedAt).ToDateTimeString(),
		UpdatedAt:   carbon.Time2Carbon(data.UpdatedAt).ToDateTimeString(),

		Channel: channelAssembly,
	}
	return advertising
}

func AdvertisingPositionAssemblyFromModelList(data []advertising_position.AdvertisingPosition, total int) interface{} {
	AdvertisingPositions := make([]AdvertisingPositionAssembly, total)
	var channelAssembly *ChannelAssembly
	for i, v := range data {
		if v.Channel != nil {
			channelAssembly = &ChannelAssembly{
				ID:        v.Channel.ID,
				Name:      v.Channel.Name,
				CreatedAt: carbon.Time2Carbon(v.Channel.CreatedAt).ToDateTimeString(),
				UpdatedAt: carbon.Time2Carbon(v.Channel.UpdatedAt).ToDateTimeString(),
			}
		}
		AdvertisingPositions[i] = AdvertisingPositionAssembly{
			ID:          v.ID,
			Name:        v.Name,
			ChannelId:   v.ChannelId,
			Code:        v.Code,
			Height:      v.Height,
			Weight:      v.Weight,
			Status:      v.Status,
			Description: v.Description,
			CreatedAt:   carbon.Time2Carbon(v.CreatedAt).ToDateTimeString(),
			UpdatedAt:   carbon.Time2Carbon(v.UpdatedAt).ToDateTimeString(),

			Channel: channelAssembly,
		}
	}

	return AdvertisingPositions
}
