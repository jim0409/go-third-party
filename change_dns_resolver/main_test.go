package change_dns_resolver

import (
	"context"
	"net"
	"testing"
	"time"
)

func TestNetResolver(t *testing.T) {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, "127.0.0.1:5301")
		},
	}
	ip, _ := r.LookupHost(context.Background(), "dns.jim.host")

	print(ip[0])
}
