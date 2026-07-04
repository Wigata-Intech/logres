package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/wigata-intech/logres/cmd"
	"github.com/wigata-intech/logres/internal/shared/cli"
)

func main() {
	err := cmd.Execute(context.Background())
	if err == nil {
		return
	}

	if errors.Is(err, cli.ErrUsage) {
		os.Exit(2)
	}
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}
