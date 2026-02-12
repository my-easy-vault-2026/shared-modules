package common

import "time"

type RateLimit struct {
	Limit     int
	Remaining int
	Used      int
	Reset     time.Time
}
