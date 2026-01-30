package entities

import "shared-modules/common"

type GetAppsForm struct {
	Category common.AppsCategory `json:"category"`
}

type FavouriteAppsForm struct {
	FavouriteID uint64 `json:"favouriteID" validate:"required"`
}
