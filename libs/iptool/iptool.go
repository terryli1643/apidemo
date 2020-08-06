package iptool

import (
	"log"
	"net"
	"net/http"
	"strings"
)

func GetRealIp(r *http.Request) string {

	value := r.Header.Get("X-Forwarded-For")
	if len(value) == 0 {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Printf("debug: Getting req.RemoteAddr %v", err)
			return ""
		}

		userIP := net.ParseIP(ip)
		if userIP == nil {
			log.Printf("debug: Parsing IP from Request.RemoteAddr got nothing.")
			return ""
		}
		return userIP.String()

	}

	addresses := strings.Split(value, ",")
	address := strings.TrimSpace(addresses[0])

	return address
}
