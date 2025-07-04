// File: internal/lib/uptime/info.go
// Purpose: Provides uptime tracking and metadata like version, commit, build date.

package uptime

import (
	"os"
	"time"
)

var startTime = time.Now()

func Since(t time.Time) string {
	return time.Since(t).Truncate(time.Second).String()
}

func Start() time.Time {
	return startTime
}

// Optional: Get hostname for diagnostics
func Hostname() string {
	h, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return h
}
