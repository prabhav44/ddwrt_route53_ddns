package main

import (
	"fmt"

	flags "../internal"
)

func main() {
	var operation flags.Operation = flags.GetFlags()
	switch operation.Name {
	case "set":
		fmt.Println("start set workflow")
	case "get":
		fmt.Println("start get workflow")
	}
}
