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
		if len(os.Args) < 3 {
			fmt.Println("Usage porty scan <host>")
			return
		}
		fmt.Printf("Scanning %s...\n", os.Args[2])
	case "dns":
		if len(os.Args) <3 {
			fmt.Println("Usage porty dns <host>")
			return
		}
		fmt.Printf("Scanning %s...\n", os.Args[2])
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
