// Package gigasecond implements solution to exercism.io prompt by the same name
package gigasecond

import (
	"time"
)

// AddGigasecond increases the given time by one gigasecond (10^9 sec)
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
