package main

import (
	"context"
	"log/slog"

	"github.com/imjasonh/gcpslog"
)

// Run this CLI to see structured logs in the terminal.
// go run ./cmd/cli
func main() {
	slog.SetDefault(slog.New(gcpslog.NewHandler()))
	ctx := context.Background()

	slog.InfoContext(ctx, "my message",
		"mycount", 42,
		"mystring", "myvalue",
	)
	slog.InfoContext(ctx, "my message",
		"mycount", 42,
		"mystring", "myvalue",
	)
	slog.Log(ctx, gcpslog.LevelCritical, "I have a bad feeling about this...")
}
