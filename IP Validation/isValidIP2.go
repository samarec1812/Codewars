package kata

import(
 "net"
)

func Is_valid_ip(ip string) bool {
	validIp := net.ParseIP(ip)
	if validIp != nil {
		return true
	}
	return false
}
