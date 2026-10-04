package main
import (
	"fmt"
	"os"
	"time"
	"github.com/404Navdeep/porty/internal/scanner"
	"github.com/404Navdeep/porty/internal/dns"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "scan":
		if len(os.Args) < 3 {
			fmt.Println("Usage porty scan <host>")
			return
		}
		host := os.Args[2]
		ports := []int{
			21,22,25,53,80,110,143,443,445,3306,5432,8080,
		}
		fmt.Printf("Scanning %s...\n\n", host)
		results := scanner.ScanPorts(host, ports, 500*time.Millisecond)

		for _, result := range results {
			if result.Open{
				fmt.Printf("	%d/tcp OPEN\n", result.Port)
			}
		}
		fmt.Println("\nScan complete.")
	case "dns":
		if len(os.Args) <3 {
			fmt.Println("Usage porty dns <host>")
			return
		}
		fmt.Printf("Scanning %s...\n", os.Args[2])
		host := os.Args[2]
		result, err := dns.Lookup(host)
		if err != nil {
			fmt.Printf("DNS lookup failed: %v\n", err)
			return
		}
		fmt.Printf("DNS lookup: %s\n\n", result.Host)

		fmt.Println("IPv4:")
		if len(result.IPv4) == 0 {
			fmt.Println("	None")
		} else {
			for _, ip := range result.IPv4 {
				fmt.Printf("	%s\n", ip)
			}
		}

		fmt.Println("IPv6:")
		if len(result.IPv6) == 0 {
			fmt.Printf("	None")
		} else {
			for _, ip := range result.IPv6 {
				fmt.Printf("	%s\n", ip)
			}
		}

	case "check":
		if len(os.Args) < 3 {
			fmt.Println("Usage porty check <host>")
			return
		}
		fmt.Printf("Scanning %s...\n", os.Args[2])
	case "help":
		printUsage()
	default:
		fmt.Printf("IDK command: %s\n", os.Args[1])
		printUsage()
	}
}
func printUsage() {
	fmt.Println("Porty, Network diagnostic toolkit")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("	porty scan <host>")
	fmt.Println("	porty dns <host>")
	fmt.Println("	porty check <hosts>")
	fmt.Println("	porty help")
}
