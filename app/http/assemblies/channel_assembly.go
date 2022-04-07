package assemblies

import (
	"gohub/app/models/channel"

	"github.com/golang-module/carbon/v2"
)

type ChannelAssembly struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ChannelAssemblyFromModel(data channel.Channel) *ChannelAssembly {
	return &ChannelAssembly{
		ID:        data.ID,
		Name:      data.Name,
		CreatedAt: carbon.Time2Carbon(data.CreatedAt).ToDateTimeString(),
		UpdatedAt: carbon.Time2Carbon(data.UpdatedAt).ToDateTimeString(),
	}
}

func ChannelAssemblyFromModelList(data []channel.Channel) interface{} {
	Channels := make([]ChannelAssembly, len(data))
	for i, v := range data {
		Channels[i] = ChannelAssembly{
			ID:        v.ID,
			Name:      v.Name,
			CreatedAt: carbon.Time2Carbon(v.CreatedAt).ToDateTimeString(),
			UpdatedAt: carbon.Time2Carbon(v.UpdatedAt).ToDateTimeString(),
		}
	}

	return Channels
}
