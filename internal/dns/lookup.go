package dns
import (
	"net"
	"sort"
)

type Result struct {
	Host string
	IPv4 []string
	IPv6 []string
}

func Lookup(host string) (Result, error) {
	ips, err :=  net.LookupIP(host)
	if err != nil {
		return Result{}, err
	}
	result := Result{Host: host,}
	for _, ip := range ips {
		if ip.To4() !=  nil {
			result.	IPv4 =  append(result.IPv4, ip.String())
		} else {
			result.IPv6 = append(result.IPv6, ip.String())
		}
	}
	sort.Strings(result.IPv4)
	sort.Strings(result.IPv6)

	return result, nil
}