package common

type OrderDirection int // 排序方向
const (
	ORDER_DIRECTION_ASC  OrderDirection = 1
	ORDER_DIRECTION_DESC OrderDirection = 2
)

func (od OrderDirection) String() string {
	switch od {
	case ORDER_DIRECTION_ASC:
		return "asc"
	case ORDER_DIRECTION_DESC:
		return "desc"
	}
	return ""
}
