package main

import (
	"context"
	"fmt"
	"os"

	"ryusuke410/golang-cli-tryout/cmd/urfavetryout/dependency"
)

func main() {
	if err := dependency.New().Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
