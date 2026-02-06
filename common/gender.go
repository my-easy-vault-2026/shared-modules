package common

type Gender int // 性別
const (
	GENDER_MALE   Gender = 1
	GENDER_FEMALE Gender = 2
)

func (g Gender) String() string {
	switch g {
	case GENDER_MALE:
		return "male"
	case GENDER_FEMALE:
		return "female"
	default:
		return ""
	}
}
