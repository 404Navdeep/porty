package scanner
import (
	"fmt"
	"net"
	"time"
)

type Result struct {
	Port int
	Open bool
	Error error
}

func Scan(host string, port int, timeout time.Duration) Result {
	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return Result{
			Port: port,
			Open: false,
			Error: err,
		}
	}
	conn.Close()

	return Result{
		Port: port,
		Open: true,
	}
}