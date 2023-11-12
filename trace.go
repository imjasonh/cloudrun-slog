package crslog

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/compute/metadata"
)

// Middleware that adds the Cloud Trace ID to the context
// This is used to correlate the structured logs with the Cloud Run
// request log.
func WithCloudTraceContext(h http.Handler) http.Handler {
	// Get the project ID from the environment if specified
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		var err error
		// Get from metadata server
		// You can avoid this dependency by using environment variables, or by connecting
		// to the metadata endpoint directly using an `http.Client`
		// See https://cloud.google.com/compute/docs/metadata/overview
		projectID, err = metadata.ProjectID()
		if err != nil {
			panic(err)
		}
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var trace string
		traceHeader := r.Header.Get("X-Cloud-Trace-Context")
		traceParts := strings.Split(traceHeader, "/")
		if len(traceParts) > 0 && len(traceParts[0]) > 0 {
			trace = fmt.Sprintf("projects/%s/traces/%s", projectID, traceParts[0])
		}
		h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "trace", trace)))
	})
}

func traceFromContext(ctx context.Context) string {
	trace := ctx.Value("trace")
	if trace == nil {
		return ""
	}
	return trace.(string)
}
