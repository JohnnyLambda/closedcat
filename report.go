package closedcat

import "time"

type Report struct {
	Created  time.Time
	Note     string
	Start    time.Time
	Duration time.Duration
}
