package common

type CardStatus int // 卡片狀態
const (
	CARD_STATUS_NOT_CREATED   CardStatus = 1
	CARD_STATUS_CREATED       CardStatus = 2
	CARD_STATUS_NOT_ACTIVATED CardStatus = 3
	CARD_STATUS_ACTIVATED     CardStatus = 4
	CARD_STATUS_DEACTIVATED   CardStatus = 5
	CARD_STATUS_DELETED       CardStatus = 6
	CARD_STATUS_BLOCKED       CardStatus = 7
)

func (s CardStatus) String() string {
	switch s {
	case CARD_STATUS_NOT_CREATED:
		return "not_created"
	case CARD_STATUS_CREATED:
		return "created"
	case CARD_STATUS_NOT_ACTIVATED:
		return "not_activated"
	case CARD_STATUS_ACTIVATED:
		return "activated"
	case CARD_STATUS_DEACTIVATED:
		return "deactivated"
	case CARD_STATUS_DELETED:
		return "deleted"
	case CARD_STATUS_BLOCKED:
		return "blocked"
	}
	return ""
}
