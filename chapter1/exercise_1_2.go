package main

import (
	"fmt"
	"os"
)

func main() {
	for i, e := range os.Args {
		fmt.Println(i, e)
	}
}
