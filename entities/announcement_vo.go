package entities

import (
	"shared-modules/common"
	"shared-modules/utils"
)

type AnnouncementPageVO struct {
	Id         uint64                    `json:"id,string"`
	Status     common.AnnouncementStatus `json:"status"`
	StartTime  *int64                    `json:"startTime"`
	EndTime    *int64                    `json:"endTime"`
	Title      string                    `json:"title"`
	CreateUser uint64                    `json:"createUser,string"`
	CreatedAt  int64                     `json:"createdAt"`
	UpdateUser uint64                    `json:"updateUser,string"`
	UpdatedAt  int64                     `json:"updatedAt"`
}

type AnnouncementVO struct {
	Id           uint64                      `json:"id,string"`
	Status       common.AnnouncementStatus   `json:"status"`
	StartTime    *int64                      `json:"startTime"`
	EndTime      *int64                      `json:"endTime"`
	CreateUser   uint64                      `json:"createUser,string"`
	CreatedAt    int64                       `json:"createdAt"`
	UpdateUser   uint64                      `json:"updateUser,string"`
	UpdatedAt    int64                       `json:"updatedAt"`
	Translations []AnnouncementTranslationVO `json:"translations"`
}

type AnnouncementTranslationVO struct {
	Title      string            `json:"title"`
	Content    string            `json:"content"`
	Language   common.Language   `json:"language"`
	NotifyType common.NotifyType `json:"notifyType"`
}

type PageAnnouncementVO struct {
	utils.PageData[[]*AnnouncementPageVO]
}

type AnnouncementMetricVO []*AnnouncementMetricVOUnit

type AnnouncementMetricVOUnit struct {
	Id         uint64            `json:"id,string"`
	NotifyType common.NotifyType `json:"notifyType"`
	Total      uint64            `json:"total,string"`
	Success    uint64            `json:"success,string"`
	Fail       uint64            `json:"fail,string"`
	Pending    uint64            `json:"pending,string"`
}
