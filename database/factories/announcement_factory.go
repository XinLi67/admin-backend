package factories

import (
    "gohub/app/models/announcement"
    "gohub/pkg/helpers"
)

func MakeAnnouncements(count int) []announcement.Announcement {

    var objs []announcement.Announcement

    // 设置唯一性，如 Announcement 模型的某个字段需要唯一，即可取消注释
    // faker.SetGenerateUniqueValues(true)S

    for i := 0; i < count; i++ {
        announcementModel := announcement.Announcement{
            AnnouncementNo:uint64(15732549211+i),
            AnnouncementPositionId:1,
            CreatorId:1,
            DepartmentId:1,
            Title:helpers.RandomString(5),
            LongTitle:helpers.RandomString(10),
            Type:0,
            Banner:helpers.RandomString(10),
            Content:helpers.RandomString(10),
            Status:0,
        }
        objs = append(objs, announcementModel)
    }

    return objs
}