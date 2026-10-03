package main
import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "scan":
		fmt.Println("Scan command")
	case "dns":
		fmt.Println("DNS command")
	case "check":
		fmt.Println("check command")
	case "help":
		printUsage()
	default:
		fmt.Println("IDK command: %s\n", os.Args[1])
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
