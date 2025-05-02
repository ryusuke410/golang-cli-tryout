package main

import (
	"os"

	"ryusuke410/golang-cli-tryout/cmd/cobratryout/dependency"
)

func main() {
	if err := dependency.New().Execute(); err != nil {
		os.Exit(1)
	}
}
