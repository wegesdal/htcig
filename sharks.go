package main

import (
	"fmt"
)

func main() {
	// Define sharks variable as a slice of strings
	sharks := []string{"hammerhead", "great white", "dogfish"}

	for _, shark := range sharks {
		fmt.Println(shark)
	}
}
