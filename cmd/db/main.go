package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"

	db "github.com/teleivo/rc-pairing"
)

func main() {
	if err := run(os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "exit due to error: %v\n", err)
		os.Exit(1)
	}
}

func run(w io.Writer) error {
	database := db.New()
	logger := slog.New(slog.NewTextHandler(w, nil))
	mux := http.NewServeMux()
	mux.HandleFunc("GET /set", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		if len(query) != 1 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "must set exactly one key using /set?somekey=somevalue")
			return
		}
		k, values := getKeyValue(query)
		if len(values) > 1 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "must set key to exactly one value using /set?somekey=somevalue")
			return
		}

		database.Set(k, values[0])
	})
	mux.HandleFunc("GET /get", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		if len(query) != 1 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "must get exactly one key using /get?somekey")
			return
		}
		k, values := getKeyValue(query)
		if len(values) > 0 && values[0] != "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "get must be called without a value, did you mean to set using /set?somekey=somevalue")
			return
		}

		v, ok := database.Get(k)

		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "%s=%s", k, v)
	})
	srv := http.Server{
		Addr:    "127.0.0.1:4000",
		Handler: log(logger, mux),
	}
	logger.Info("Listening", "Addr", srv.Addr)
	return srv.ListenAndServe()
}

func getKeyValue(query url.Values) (string, []string) {
	for k, v := range query {
		return k, v
	}
	return "", nil
}

func log(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Request", slog.String("path", r.URL.Path), slog.String("method", r.Method), slog.Any("query", r.URL.Query()))
		next.ServeHTTP(w, r)
	})
}
