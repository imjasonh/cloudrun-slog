# Lightweight structured logging for Google Cloud Run using [`slog`](https://pkg.go.dev/log/slog)

Contrary to the [documented 'standard' approach for logging](https://cloud.google.com/logging/docs/setup/go), this doesn't use any third-party logging package for logging.

Instead, it relies on Cloud Run's support for ingesting structured logs by [simply printing JSON to standard error](https://cloud.google.com/run/docs/logging#using-json).

(Note that App Engine supports the same structured JSON output approach, so the same code can be used there)

## Usage

To use this, underscore-import this library, which will configure `slog` to use the JSON handler for all log messages:

```go
import _ "github.com/imjasonh/cloudrun-slog"
```

Then when you use `slog`, all log messages will be output in JSON format to standard error, which is automatically ingested by Cloud Logging.

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  slog.InfoContext(r.Context(), "my message",
    "mycount", 42,
    "mystring", "myvalue",
	)
})
```
