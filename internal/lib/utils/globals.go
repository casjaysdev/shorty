// File: internal/lib/utils/globals.go
// Purpose: Centralized global defaults for system-wide fallback behaviors.

package utils

var (
	DefaultTheme           = "dracula"
	DefaultSiteAlignment   = "center"
	DefaultLayoutPreset    = "balanced"
	DefaultSlugLength      = 6
	AllowCustomSlugs       = true
	CustomSlugCaseSensitive = false
	DefaultTrustedProxies  = []string{
		"127.0.0.0/8",
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"::1/128",
		"fc00::/7",
	}

	DefaultSMTPPort     = 587
	DefaultFromName     = "Shorty"
	DefaultFromEmail    = "no-reply@localhost"
	DefaultEmailCharset = "UTF-8"

	FreePlanFeatures = []string{
		"custom_theme:false",
		"analytics:false",
		"custom_domain:false",
		"white_label:false",
	}

	DefaultPlanPricing = map[string]float64{
		"pro":      5.00,
		"business": 10.00,
	}

	AnnualDiscountPercent = 20
	DefaultBillingCycle    = "monthly" // or "annual"
)

func FeatureEnabled(plan string, feature string) bool {
	if plan == "free" {
		for _, f := range FreePlanFeatures {
			if f == feature+":true" {
				return false
			}
		}
	}
	return true
}
