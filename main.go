package main

import (
	"fmt"
	"os"
)

func main() {
	var n string = os.Getenv("TEST_SECRET")
	fmt.Print(n)
}
