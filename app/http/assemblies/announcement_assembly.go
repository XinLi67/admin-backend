package assemblies

import (
	"gohub/app/models/announcement"

	"github.com/golang-module/carbon/v2"
)

type AnnouncementAssembly struct {
	ID                     uint64 `json:"id"`
	AnnouncementNo         uint64 `json:"announcement_no"`
	AnnouncementPositionId uint64 `json:"announcement_position_id"`
	UserId                 uint64 `json:"user_id"`
	Title                  string `json:"title"`
	Type                   uint64 `json:"type"`
	RedirectTo             uint64 `json:"redirect_to"`
	RedirectParams         string `json:"redirect_params"`
	Content                string `json:"content"`
	Status                 uint64 `json:"status"`
	AuditReason            string `json:"audit_reason"`
	SchedulingType         uint64 `json:"scheduling_type"`
	StartDate              string `json:"start_date"`
	EndDate                string `json:"end_date"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`

	User                 *UserAssembly                 `json:"user"`
	AnnouncementPosition *AnnouncementPositionAssembly `json:"announcement_position"`
	Channel              *ChannelAssembly              `json:"channel"`
}

func AnnouncementAssemblyFromModel(data announcement.Announcement) *AnnouncementAssembly {
	var userAssembly *UserAssembly
	var announcementPosition *AnnouncementPositionAssembly
	var channelAssembly *ChannelAssembly

	if data.User != nil {
		userAssembly = &UserAssembly{
			ID:        data.User.ID,
			Name:      data.User.Name,
			CreatedAt: carbon.Time2Carbon(data.AnnouncementPosition.CreatedAt).ToDateTimeString(),
			UpdatedAt: carbon.Time2Carbon(data.AnnouncementPosition.UpdatedAt).ToDateTimeString(),
		}
	}
	if data.AnnouncementPosition != nil {
		announcementPosition = &AnnouncementPositionAssembly{
			ID:        data.AnnouncementPosition.ID,
			Name:      data.AnnouncementPosition.Name,
			CreatedAt: carbon.Time2Carbon(data.AnnouncementPosition.CreatedAt).ToDateTimeString(),
			UpdatedAt: carbon.Time2Carbon(data.AnnouncementPosition.UpdatedAt).ToDateTimeString(),
		}
	}
	if data.Channel != nil {
		channelAssembly = &ChannelAssembly{
			ID:        data.Channel.ID,
			Name:      data.Channel.Name,
			CreatedAt: carbon.Time2Carbon(data.AnnouncementPosition.CreatedAt).ToDateTimeString(),
			UpdatedAt: carbon.Time2Carbon(data.AnnouncementPosition.UpdatedAt).ToDateTimeString(),
		}
	}
	announcement := &AnnouncementAssembly{
		ID:                     data.ID,
		AnnouncementNo:         data.AnnouncementNo,
		AnnouncementPositionId: data.AnnouncementPositionId,
		UserId:                 data.UserId,
		Title:                  data.Title,
		Type:                   data.Type,
		RedirectTo:             data.RedirectTo,
		RedirectParams:         data.RedirectParams,
		Content:                data.Content,
		Status:                 data.Status,
		AuditReason:            data.AuditReason,
		SchedulingType:         data.SchedulingType,
		StartDate:              data.StartDate,
		EndDate:                data.EndDate,
		CreatedAt:              carbon.Time2Carbon(data.CreatedAt).ToDateTimeString(),
		UpdatedAt:              carbon.Time2Carbon(data.UpdatedAt).ToDateTimeString(),
		User:                   userAssembly,
		AnnouncementPosition:   announcementPosition,
		Channel:                channelAssembly,
	}
	return announcement
}

func AnnouncementAssemblyFromModelList(data []announcement.Announcement, total int) interface{} {
	Announcements := make([]AnnouncementAssembly, total)
	var userAssembly *UserAssembly
	var announcementPosition *AnnouncementPositionAssembly
	var channelAssembly *ChannelAssembly
	for i, v := range data {
		if v.User != nil {
			userAssembly = &UserAssembly{
				ID:        v.User.ID,
				UserName:  v.User.Username,
				CreatedAt: carbon.Time2Carbon(v.User.CreatedAt).ToDateTimeString(),
				UpdatedAt: carbon.Time2Carbon(v.User.UpdatedAt).ToDateTimeString(),
			}
		}
		if v.AnnouncementPosition != nil {
			announcementPosition = &AnnouncementPositionAssembly{
				ID:        v.AnnouncementPosition.ID,
				Name:      v.AnnouncementPosition.Name,
				CreatedAt: carbon.Time2Carbon(v.AnnouncementPosition.CreatedAt).ToDateTimeString(),
				UpdatedAt: carbon.Time2Carbon(v.AnnouncementPosition.UpdatedAt).ToDateTimeString(),
			}
		}
		if v.Channel != nil {
			userAssembly = &UserAssembly{
				ID:        v.Channel.ID,
				UserName:  v.Channel.Name,
				CreatedAt: carbon.Time2Carbon(v.User.CreatedAt).ToDateTimeString(),
				UpdatedAt: carbon.Time2Carbon(v.User.UpdatedAt).ToDateTimeString(),
			}
		}
		Announcements[i] = AnnouncementAssembly{
			ID:                     v.ID,
			AnnouncementNo:         v.AnnouncementNo,
			AnnouncementPositionId: v.AnnouncementPositionId,
			UserId:                 v.UserId,
			Title:                  v.Title,
			Type:                   v.Type,
			RedirectTo:             v.RedirectTo,
			RedirectParams:         v.RedirectParams,
			Content:                v.Content,
			Status:                 v.Status,
			AuditReason:            v.AuditReason,
			SchedulingType:         v.SchedulingType,
			StartDate:              v.StartDate,
			EndDate:                v.EndDate,
			CreatedAt:              carbon.Time2Carbon(v.CreatedAt).ToDateTimeString(),
			UpdatedAt:              carbon.Time2Carbon(v.UpdatedAt).ToDateTimeString(),
			User:                   userAssembly,
			AnnouncementPosition:   announcementPosition,
			Channel:                channelAssembly,
		}
	}

	return Announcements
}
