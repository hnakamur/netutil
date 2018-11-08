package netutil

import (
	"net"
)

var (
	ipv4Private24bitBlockIPNet = net.IPNet{
		IP:   net.IP{10, 0, 0, 0},
		Mask: net.CIDRMask(8, 8*net.IPv4len),
	}
	ipv4Private20bitBlockIPNet = net.IPNet{
		IP:   net.IP{172, 16, 0, 0},
		Mask: net.CIDRMask(12, 8*net.IPv4len),
	}
	ipv4Private16bitBlockIPNet = net.IPNet{
		IP:   net.IP{192, 168, 0, 0},
		Mask: net.CIDRMask(16, 8*net.IPv4len),
	}
	ipv6UniqueLocalIPNet = net.IPNet{
		IP:   net.IP{0xfc, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Mask: net.CIDRMask(7, 8*net.IPv6len),
	}
)

// IsUniqueLocal reports whether ip is a unique local address.
//
// The identification of unique local addresses uses address
// type identification as defined in "Local IPv6 Unicast Addresses"
// section of RFC 4193 and "Private Address Space" in RFC 1918.
func IsUniqueLocal(ip net.IP) bool {
	if ip4 := ip.To4(); ip4 != nil {
		return ipv4Private24bitBlockIPNet.Contains(ip4) ||
			ipv4Private20bitBlockIPNet.Contains(ip4) ||
			ipv4Private16bitBlockIPNet.Contains(ip4)
	}
	return ipv6UniqueLocalIPNet.Contains(ip)
}
