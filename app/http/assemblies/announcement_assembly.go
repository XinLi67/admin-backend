package assemblies

import (
	"gohub/app/models/announcement"

	"github.com/golang-module/carbon/v2"
)

type AnnouncementAssembly struct {
	ID                     uint64 `json:"id"`
	AnnouncementNo         uint64 `json:"announcement_no"`
	AnnouncementPositionId uint64 `json:"announcement_position_id"`
	CreatorId              uint64 `json:"creator_id"`
	DepartmentId           uint64 `json:"department_id"`
	Title                  string `json:"title"`
	LongTitle              string `json:"long_title"`
	Type                   uint64 `json:"type"`
	Banner                 string `json:"banner"`
	RedirectTo             uint64 `json:"redirect_to"`
	RedirectParams         string `json:"redirect_params"`
	Content                string `json:"content"`
	Status                 uint64 `json:"status"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`

	User                 UserAssembly                 `json:"user"`
	AnnouncementPosition AnnouncementPositionAssembly `json:"announcement_position"`
}

func AnnouncementAssemblyFromModel(data announcement.Announcement) *AnnouncementAssembly {
	return &AnnouncementAssembly{
		ID:                     data.ID,
		AnnouncementNo:         data.AnnouncementNo,
		AnnouncementPositionId: data.AnnouncementPositionId,
		CreatorId:              data.CreatorId,
		DepartmentId:           data.DepartmentId,
		Title:                  data.Title,
		LongTitle:              data.LongTitle,
		Type:                   data.Type,
		Banner:                 data.Banner,
		RedirectTo:             data.RedirectTo,
		RedirectParams:         data.RedirectParams,
		Content:                data.Content,
		Status:                 data.Status,
		CreatedAt:              carbon.Time2Carbon(data.CreatedAt).ToDateTimeString(),
		UpdatedAt:              carbon.Time2Carbon(data.UpdatedAt).ToDateTimeString(),

		AnnouncementPosition: AnnouncementPositionAssembly{
			ID:        data.AnnouncementPosition.ID,
			Name:      data.AnnouncementPosition.Name,
			CreatedAt: carbon.Time2Carbon(data.AnnouncementPosition.CreatedAt).ToDateTimeString(),
			UpdatedAt: carbon.Time2Carbon(data.AnnouncementPosition.UpdatedAt).ToDateTimeString(),
		},
		User: UserAssembly{
			ID:        data.User.ID,
			UserName:  data.User.Username,
			CreatedAt: carbon.Time2Carbon(data.User.CreatedAt).ToDateTimeString(),
			UpdatedAt: carbon.Time2Carbon(data.User.UpdatedAt).ToDateTimeString(),
		},
	}
}

func AnnouncementAssemblyFromModelList(data []announcement.Announcement, total int) interface{} {
	Announcements := make([]AnnouncementAssembly, total)
	for i, v := range data {
		Announcements[i] = AnnouncementAssembly{
			ID:                     v.ID,
			AnnouncementNo:         v.AnnouncementNo,
			AnnouncementPositionId: v.AnnouncementPositionId,
			CreatorId:              v.CreatorId,
			DepartmentId:           v.DepartmentId,
			Title:                  v.Title,
			LongTitle:              v.LongTitle,
			Type:                   v.Type,
			Banner:                 v.Banner,
			RedirectTo:             v.RedirectTo,
			RedirectParams:         v.RedirectParams,
			Content:                v.Content,
			Status:                 v.Status,
			CreatedAt:              carbon.Time2Carbon(v.CreatedAt).ToDateTimeString(),
			UpdatedAt:              carbon.Time2Carbon(v.UpdatedAt).ToDateTimeString(),

			AnnouncementPosition: AnnouncementPositionAssembly{
				ID:        v.AnnouncementPosition.ID,
				Name:      v.AnnouncementPosition.Name,
				CreatedAt: carbon.Time2Carbon(v.AnnouncementPosition.CreatedAt).ToDateTimeString(),
				UpdatedAt: carbon.Time2Carbon(v.AnnouncementPosition.UpdatedAt).ToDateTimeString(),
			},
			User: UserAssembly{
				ID:        v.User.ID,
				UserName:  v.User.Username,
				CreatedAt: carbon.Time2Carbon(v.User.CreatedAt).ToDateTimeString(),
				UpdatedAt: carbon.Time2Carbon(v.User.UpdatedAt).ToDateTimeString(),
			},
		}
	}

	return Announcements
}
