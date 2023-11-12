package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	_ "github.com/imjasonh/cloudrun-slog"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.InfoContext(r.Context(), "my message",
			"mycount", 42,
			"mystring", "myvalue",
		)
		fmt.Fprintln(w, "Hello, world!")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
