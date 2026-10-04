package scanner
import (
	"fmt"
	"net"
	"time"
	"sort"
	"sync"
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

func ScanPorts(host string, ports []int, timeout time.Duration) []Result {
	results := make(chan Result, len(ports))

	var wg sync.WaitGroup

	for _, port := range ports {
		wg.Add(1)

		go func(port int) {
			defer wg.Done()

			results <- Scan(host, port, timeout)
		}(port)
	}

	wg.Wait()
	close(results)

	var output []Result

	for result := range results {
		output = append(output, result)
	}

	sort.Slice(output, func(i, j int) bool {
		return output[i].Port < output[j].Port
	})

	return output
}