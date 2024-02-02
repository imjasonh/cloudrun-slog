package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/imjasonh/gcpslog"
)

// Deploy this to Cloud Run to see the logs in the Cloud Logging UI.
// gcloud run deploy gcpslog --image $(ko build ./cmd/app)
func main() {
	verbose := flag.Bool("verbose", false, "enable verbose logging")
	flag.Parse()
	level := slog.LevelInfo
	if *verbose {
		level = slog.LevelDebug
	}

	slog.SetDefault(slog.New(gcpslog.NewHandler(level)))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.InfoContext(r.Context(), "my message",
			"mycount", 42,
			"mystring", "myvalue",
		)
		fmt.Fprintln(w, "Check logs for basic output")
	})

	http.Handle("/correlated", gcpslog.WithCloudTraceContext(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.InfoContext(r.Context(), "my message",
			"mycount", 42,
			"mystring", "myvalue",
		)
		fmt.Fprintln(w, "Check logs for correlated output")
	})))

	http.Handle("/debug", gcpslog.WithCloudTraceContext(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.InfoContext(r.Context(), "my message",
			"mycount", 42,
			"mystring", "myvalue",
		)
		fmt.Fprintln(w, "Check logs for debug output")
	})))

	http.HandleFunc("/critical", func(w http.ResponseWriter, r *http.Request) {
		slog.Log(r.Context(), gcpslog.LevelCritical, "I have a bad feeling about this...")
		fmt.Fprintln(w, "Check logs for critical output")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
