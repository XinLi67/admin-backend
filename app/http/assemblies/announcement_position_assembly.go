package assemblies

import (
	"gohub/app/models/announcement_position"

	"github.com/golang-module/carbon/v2"
)

type AnnouncementPositionAssembly struct {
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

	Channel ChannelAssembly `json:"channel"`
}

func AnnouncementPositionAssemblyFromModel(data announcement_position.AnnouncementPosition) *AnnouncementPositionAssembly {
	return &AnnouncementPositionAssembly{
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

		Channel: ChannelAssembly{
			ID:        data.Channel.ID,
			Name:      data.Channel.Name,
			CreatedAt: carbon.Time2Carbon(data.Channel.CreatedAt).ToDateTimeString(),
			UpdatedAt: carbon.Time2Carbon(data.Channel.UpdatedAt).ToDateTimeString(),
		},
	}
}

func AnnouncementPositionAssemblyFromModelList(data []announcement_position.AnnouncementPosition, total int) interface{} {
	announcementPositions := make([]AnnouncementPositionAssembly, total)
	for i, v := range data {
		announcementPositions[i] = AnnouncementPositionAssembly{
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

			Channel: ChannelAssembly{
				ID:        v.Channel.ID,
				Name:      v.Channel.Name,
				CreatedAt: carbon.Time2Carbon(v.Channel.CreatedAt).ToDateTimeString(),
				UpdatedAt: carbon.Time2Carbon(v.Channel.UpdatedAt).ToDateTimeString(),
			},
		}
	}

	return announcementPositions
}
