package main

import (
	"fmt"
	"os"
)

// access_control_manager - Manage access control
func access_control_manager(path string) {
	fmt.Println("========================================")
	fmt.Println("  Access-Control-Manager")
	fmt.Println("  Manage access control")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	access_control_manager(path)
}
