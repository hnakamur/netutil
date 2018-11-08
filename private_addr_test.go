package netutil_test

import (
	"net"
	"testing"

	"github.com/hnakamur/netutil"
)

func TestIsUniqueLocal(t *testing.T) {
	testCases := []struct {
		ip   string
		want bool
	}{
		{ip: "fbff:ffff:ffff:ffff:ffff:ffff:ffff:ffff", want: false},
		{ip: "fc00::", want: true},
		{ip: "fd00::1", want: true},
		{ip: "fdff:ffff:ffff:ffff:ffff:ffff:ffff:ffff", want: true},
		{ip: "fe00::", want: false},
		{ip: "9.255.255.255", want: false},
		{ip: "10.0.0.0", want: true},
		{ip: "10.255.255.255", want: true},
		{ip: "11.0.0.0", want: false},
		{ip: "172.15.255.255", want: false},
		{ip: "172.16.0.0", want: true},
		{ip: "172.31.255.255", want: true},
		{ip: "172.32.0.0", want: false},
		{ip: "192.167.255.255", want: false},
		{ip: "192.168.0.0", want: true},
		{ip: "192.168.255.255", want: true},
		{ip: "192.169.0.0", want: false},
	}
	for _, c := range testCases {
		got := netutil.IsUniqueLocal(net.ParseIP(c.ip))
		if got != c.want {
			t.Errorf("unexpected result from IsUniqueLocal for ip=%s, got=%v, want=%v", c.ip, got, c.want)
		}
	}
}
