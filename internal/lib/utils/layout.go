// File: internal/lib/utils/layout.go
// Purpose: Layout configuration helpers for theme, padding, and alignment.

package utils

import "strings"

var AvailableThemes = []string{"light", "dark", "dracula"}
var AvailableAlignments = []string{"left", "center", "right"}
var AvailableLayoutPresets = []string{"compact", "balanced", "spacious"}

func IsValidTheme(t string) bool {
	t = strings.ToLower(t)
	for _, theme := range AvailableThemes {
		if t == theme {
			return true
		}
	}
	return false
}

func IsValidAlignment(a string) bool {
	a = strings.ToLower(a)
	for _, align := range AvailableAlignments {
		if a == align {
			return true
		}
	}
	return false
}

func IsValidLayoutPreset(p string) bool {
	p = strings.ToLower(p)
	for _, preset := range AvailableLayoutPresets {
		if p == preset {
			return true
		}
	}
	return false
}
