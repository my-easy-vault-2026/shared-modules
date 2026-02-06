package common

type PreviewPurpose int // 預覽目的
const (
	PREVIEW_PURPOSE_TRANSFER PreviewPurpose = 2
	PREVIEW_PURPOSE_EXCHANGE PreviewPurpose = 3
)

func (pp PreviewPurpose) String() string {
	switch pp {
	case PREVIEW_PURPOSE_TRANSFER:
		return "transfer"
	case PREVIEW_PURPOSE_EXCHANGE:
		return "exchange"

	}
	return ""
}
