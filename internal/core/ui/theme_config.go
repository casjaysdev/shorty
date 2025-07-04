// Source: internal/core/ui/theme_config.go
// Purpose: Provides global and user-level theme options including Dracula default

package ui

type Theme string

const (
	ThemeDark    Theme = "dark"
	ThemeLight   Theme = "light"
	ThemeDracula Theme = "dracula"
)

type ThemeConfig struct {
	GlobalDefault Theme `json:"global_default"`
	AllowOverride bool  `json:"allow_override"`
}

func DefaultThemeConfig() ThemeConfig {
	return ThemeConfig{
		AllowOverride: true,
		GlobalDefault: ThemeDracula,
	}
}
