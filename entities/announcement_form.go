package entities

import (
	"shared-modules/common"
	"shared-modules/utils"
)

type AnnouncementPageForm struct {
	Status common.AnnouncementStatus `json:"status"`
	utils.Page
}

type AnnouncementGetForm struct {
	Id uint64 `json:"id,string" validate:"required"`
}

type AnnouncementSaveForm struct {
	StartTime    int64                       `json:"startTime" validate:"required"`
	EndTime      int64                       `json:"endTime" validate:"required"`
	Translations AnnouncementTranslationForm `json:"translations"`
}

type AnnouncementEditForm struct {
	Id           uint64                      `json:"id,string" validate:"required"`
	Status       common.AnnouncementStatus   `json:"status"`
	StartTime    int64                       `json:"startTime" validate:"required"`
	EndTime      int64                       `json:"endTime" validate:"required"`
	Translations AnnouncementTranslationForm `json:"translations"`
}

type AnnouncementSaveAndPostForm struct {
	Id           string                      `json:"id"`
	StartTime    int64                       `json:"startTime"`
	EndTime      int64                       `json:"endTime"`
	Translations AnnouncementTranslationForm `json:"translations"`
}

type AnnouncementPublishForm struct {
	IDs []uint64 `json:"ids"`
}

type AnnouncementDeleteForm struct{}

type AnnouncementPublishMailForm struct {
	Passageway string `json:"passageway"`
	Limit      int    `json:"limit"`
}

type AnnouncementTranslationForm []*AnnouncementTranslationUnit

func (a AnnouncementTranslationForm) Get(lang common.Language) *AnnouncementTranslationUnit {
	for _, l := range a {
		if l.Language == lang {
			tl := l
			return tl
		}
	}
	return nil
}

func (a AnnouncementTranslationForm) Gets(lang common.Language) AnnouncementTranslationForm {
	var ret AnnouncementTranslationForm
	for _, l := range a {
		if l.Language == lang {
			ret = append(ret, l)
		}
	}
	return ret
}

type AnnouncementTranslationUnit struct {
	Title      string            `json:"title"`
	Content    string            `json:"content"`
	Language   common.Language   `json:"language"`
	NotifyType common.NotifyType `json:"notifyType"`
}

type AnnouncementPublishAppForm struct {
	Limit int `json:"limit"`
}

type AnnouncementMetricForm struct {
	Id uint64 `json:"id,string" validate:"required"`
}

type AnnouncementResendForm struct {
	Id uint64 `json:"id,string" validate:"required"`
}
