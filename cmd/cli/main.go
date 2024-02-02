package main

import (
	"context"
	"flag"
	"log/slog"

	"github.com/imjasonh/gcpslog"
)

// Run this CLI to see structured logs in the terminal.
// go run ./cmd/cli
// Add --verbose to see debug logs
func main() {
	verbose := flag.Bool("verbose", false, "enable verbose logging")
	flag.Parse()
	level := slog.LevelInfo
	if *verbose {
		level = slog.LevelDebug
	}

	slog.SetDefault(slog.New(gcpslog.NewHandler(level)))
	ctx := context.Background()

	slog.InfoContext(ctx, "my message",
		"mycount", 42,
		"mystring", "myvalue",
	)
	slog.InfoContext(ctx, "my message",
		"mycount", 42,
		"mystring", "myvalue",
	)
	slog.Debug("my debug message")
	slog.Log(ctx, gcpslog.LevelCritical, "I have a bad feeling about this...")

}
