package common

type AssetType int // 資產類型
const (
	ASSET_TYPE_CRYPTO  AssetType = 1
	ASSET_TYPE_FIAT    AssetType = 2
	ASSET_TYPE_PRODUCT AssetType = 3
	ASSET_TYPE_POINT   AssetType = 4
)

func GetAssetType(categoryID uint64) AssetType {
	switch true {
	case categoryID >= 1000 && categoryID < 2000:
		return ASSET_TYPE_CRYPTO
	case categoryID >= 2000 && categoryID < 3000:
		return ASSET_TYPE_FIAT
	case categoryID >= 3000 && categoryID < 10000:
		return ASSET_TYPE_PRODUCT
	case categoryID >= 10000 && categoryID < 20000:
		return ASSET_TYPE_POINT
	}
	return 0
}
