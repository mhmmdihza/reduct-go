package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: emptyfile <filename>")
		os.Exit(1)
	}
	filename := os.Args[1]

	// Create or truncate (clear) the file
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	defer f.Close()
	_, err = f.WriteString("swagger: '2.0'")
	if err != nil {
		fmt.Println("error writing:", err)
		os.Exit(1)
	}
}
