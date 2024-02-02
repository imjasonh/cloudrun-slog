package init

import (
	"log/slog"

	"github.com/imjasonh/gcpslog"
)

// Set up structured logging at Info+ level.
func init() { slog.SetDefault(slog.New(gcpslog.NewHandler(slog.LevelInfo))) }
