// Source: internal/middleware/reverse_proxy.go
// Purpose: Handles trusted proxy IPs, forwarded headers, and logs relevant request metadata

package middleware

import (
	"net"
	"net/http"
	"strings"
)

var defaultTrustedRanges = []string{
	"127.0.0.0/8",
	"10.0.0.0/8",
	"172.16.0.0/12",
	"192.168.0.0/16",
	"::1/128",
	"fc00::/7",
	"fe80::/10",
}

// TrustProxyConfig allows extending trusted IPs dynamically
type TrustProxyConfig struct {
	ExtraTrusted []string
}

func ProxyHandler(config TrustProxyConfig) func(http.Handler) http.Handler {
	trusted := append(defaultTrustedRanges, config.ExtraTrusted...)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := realIP(r)
			if isTrustedIP(ip, trusted) {
				r.RemoteAddr = ip
			}
			next.ServeHTTP(w, r)
		})
	}
}

func realIP(r *http.Request) string {
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		parts := strings.Split(xff, ",")
		return strings.TrimSpace(parts[0])
	}
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	return host
}

func isTrustedIP(ip string, trustedCIDRs []string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	for _, cidr := range trustedCIDRs {
		_, network, err := net.ParseCIDR(cidr)
		if err == nil && network.Contains(parsedIP) {
			return true
		}
	}
	return false
}
