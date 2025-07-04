// File: internal/lib/utils/format.go
// Purpose: Utility functions for formatting time, byte sizes, and standard output.

package utils

import (
	"fmt"
	"time"
)

func HumanDuration(d time.Duration) string {
	if d.Hours() >= 24 {
		days := int(d.Hours()) / 24
		return fmt.Sprintf("%dd %dh", days, int(d.Hours())%24)
	} else if d.Hours() >= 1 {
		return fmt.Sprintf("%.0fh %.0fm", d.Hours(), d.Minutes()-float64(int(d.Hours())*60))
	}
	return d.Truncate(time.Second).String()
}

func HumanBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

func YesNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}
