package crslog

import (
	"context"
	"log/slog"
	"os"
)

// Set up structured logging
func init() { slog.SetDefault(slog.New(newHandler())) }

// Handler that outputs JSON understood by the structured log agent.
// See https://cloud.google.com/logging/docs/agent/logging/configuration#special-fields
type handler struct{ handler slog.Handler }

func newHandler() *handler {
	return &handler{handler: slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.MessageKey {
				a.Key = "message"
			} else if a.Key == slog.SourceKey {
				a.Key = "logging.googleapis.com/sourceLocation"
			} else if a.Key == slog.LevelKey {
				a.Key = "severity"
				level := a.Value.Any().(slog.Level)
				if level == slog.Level(12) { // Extra log level supported by Cloud Logging
					a.Value = slog.StringValue("CRITICAL")
				}
			}
			return a
		},
	})}
}

func (h *handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *handler) Handle(ctx context.Context, rec slog.Record) error {
	return h.handler.Handle(ctx, rec)
}

func (h *handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &handler{handler: h.handler.WithAttrs(attrs)}
}

func (h *handler) WithGroup(name string) slog.Handler {
	return &handler{handler: h.handler.WithGroup(name)}
}
