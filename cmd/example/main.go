package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	gcpslog "github.com/imjasonh/gcpslog"
)

func main() {
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

	http.HandleFunc("/critical", func(w http.ResponseWriter, r *http.Request) {
		slog.Log(r.Context(), gcpslog.LevelCritical, "I have a bad feeling about this...")
		fmt.Fprintln(w, "Check logs for critical output")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
