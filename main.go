package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "ERROR: cannot calculate color temperature")
	os.Exit(1)
}
