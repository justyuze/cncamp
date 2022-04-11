package main

import (
	"fmt"
)

func main() {
	arr := []string{"I", "am", "stupid", "and", "weak"}
	fmt.Println(arr)
	slice := arr[:]
	for idx, value := range slice {
		switch value {
		case "stupid":
			slice[idx] = "smart"
		case "weak":
			slice[idx] = "strong"
		}
	}
	fmt.Println(arr)
}
