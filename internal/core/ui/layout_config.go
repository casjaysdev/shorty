// Source: internal/core/ui/layout_config.go
// Purpose: Controls global layout alignment and spacing presets site-wide

package ui

type SiteAlignment string
type LayoutPreset string

const (
	AlignLeft   SiteAlignment = "left"
	AlignCenter SiteAlignment = "center"
	AlignRight  SiteAlignment = "right"

	PresetCompact  LayoutPreset = "compact"
	PresetBalanced LayoutPreset = "balanced"
	PresetSpacious LayoutPreset = "spacious"
)

type LayoutConfig struct {
	Alignment SiteAlignment `json:"alignment"`
	Preset    LayoutPreset  `json:"preset"`
}

func DefaultLayoutConfig() LayoutConfig {
	return LayoutConfig{
		Alignment: AlignCenter,
		Preset:    PresetBalanced,
	}
}
