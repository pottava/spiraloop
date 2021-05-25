package lib

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// Wrap wraps HTTP request handler
func Wrap(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case eqauls(r, "/health"):
			w.WriteHeader(http.StatusOK)

		case eqauls(r, "/version"):
			if len(commit) > 0 && len(date) > 0 {
				fmt.Fprintf(w, "%s-%s (built at %s)\n", ver, commit, date)
				return
			}
			fmt.Fprintln(w, ver)

		default:
			proc := time.Now()
			addr := r.RemoteAddr
			if ip, found := header(r, "X-Forwarded-For"); found {
				addr = ip[0]
			}
			ioWriter := w.(io.Writer)
			writer := overrideWriter(w, ioWriter)
			handler.ServeHTTP(writer, r)

			if Config.AccessLog {
				log.Printf("[%s] %.3f %d %s %s",
					addr, time.Since(proc).Seconds(),
					writer.status, r.Method, r.URL)
			}
		}
	})
}

func eqauls(r *http.Request, url string) bool {
	return url == r.URL.Path
}

func header(r *http.Request, key string) (values []string, found bool) {
	if r.Header == nil {
		return
	}
	for k, v := range r.Header {
		if strings.EqualFold(k, key) && len(v) > 0 {
			return v, true
		}
	}
	return
}
