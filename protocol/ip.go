package protocol

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

const (
	IpParam   = "ip"
	IpV4Param = "ipv4"
	IpV6Param = "ipv6"
)

type ClientIP struct {
	IP   net.IP
	IpV6 bool
}

func (c ClientIP) String() string {
	return fmt.Sprintf("%s, Supports IpV6: %t", c.IP, c.IpV6)
}

func GetIP(r *http.Request) *ClientIP {

	values := r.URL.Query()

	for _, key := range [3]string{IpParam, IpV4Param, IpV6Param} {
		iprStr, found := values[key]
		if found {
			ip := net.ParseIP(iprStr[0])
			return &ClientIP{IP: ip, IpV6: key == IpV6Param}
		}
	}

	for _, header := range [2]string{"X-Real-Ip", "X-Forwarded-For"} {
		if headerVal := r.Header.Get(header); headerVal != "" {
			ip := net.ParseIP(headerVal)
			return &ClientIP{
				IP:   ip,
				IpV6: strings.Count(headerVal, ":") > 1,
			}
		}
	}

	httpAddr, _, _ := net.SplitHostPort(r.RemoteAddr)
	return &ClientIP{
		IP:   net.ParseIP(httpAddr),
		IpV6: strings.Contains(httpAddr, ":"),
	}
}
