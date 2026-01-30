package entities

import "shared-modules/common"

type AppsVO struct {
	ID          uint64              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	IconURL     string              `json:"iconUrl"`
	AppLink     string              `json:"appLink"`
	Category    common.AppsCategory `json:"category"`
	Extra       string              `json:"extra,string"`
}

type ListAppsVO struct {
	Records []*AppsVO `json:"records"`
}

type FavouriteAppsVO struct {
	ID          uint64              `json:"id"`
	AppID       uint64              `json:"appId"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	IconURL     string              `json:"iconUrl"`
	AppLink     string              `json:"appLink"`
	Category    common.AppsCategory `json:"category"`
	Extra       string              `json:"extra"`
}
type ListFavouriteAppsVO struct {
	Records []*FavouriteAppsVO `json:"records"`
}
