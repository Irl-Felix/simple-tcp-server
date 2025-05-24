package app

import (
	"net/http"
	"strings"
)

// Helper to get real client IP
func getClientIP(req *http.Request) string {
	// Check common proxy headers
	forwarded := req.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// May contain multiple IPs, first one is the client
		parts := strings.Split(forwarded, ",")
		return strings.TrimSpace(parts[0])
	}
	realIP := req.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}
	// Fallback to remote address (includes port)
	ipPort := strings.Split(req.RemoteAddr, ":")
	if len(ipPort) > 0 {
		return ipPort[0]
	}
	return req.RemoteAddr
}
