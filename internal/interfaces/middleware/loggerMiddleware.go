package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"
)

type responseRecorder struct {
	http.ResponseWriter
	status int
	body   bytes.Buffer
}

func (rc *responseRecorder) WriteHeader(code int) {
	rc.status = code
	rc.ResponseWriter.WriteHeader(code)
}

func (rc *responseRecorder) Write(b []byte) (int, error) {
	rc.body.Write(b) // capture the body
	return rc.ResponseWriter.Write(b)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wr := &responseRecorder{ResponseWriter: w, status: http.StatusOK}
		if r.Method == "POST" {
			if r.Body != nil {
				bodyByte, err := io.ReadAll(r.Body)
				if err != nil {
					log.Printf("Error reading request body: %v", err)
				} else {
					r.Body.Close()
					r.Body = io.NopCloser(bytes.NewBuffer(bodyByte))
					next.ServeHTTP(wr, r)
					log.Printf("%s %s took %v. Body: %s.\n Response status %v, Response body: %v", r.Method, r.URL.Path, time.Since(start), string(bodyByte), wr.status, wr.body.String())
				}

			}
		} else {
			next.ServeHTTP(wr, r)
			log.Printf("%s %s took %v.\n Response status %v, Response body: %v", r.Method, r.URL.Path, time.Since(start), wr.status, wr.body.String())
		}
	})
}
