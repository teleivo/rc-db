package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
)

func main() {
	if err := run(os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "exit due to error: %v\n", err)
		os.Exit(1)
	}
}

// TODO slog listening and requests
// TODO add tests on the HTTP or DB layer? (also that set?a= deletes the value)
// TODO encapsulate state so we can then refactor it to write to a file
// TODO should I allow setting it to nothing? or rather expect a DELETE
// TODO read up on new http pattern changes
func run(w io.Writer) error {
	dbState := make(map[string]string)
	logger := slog.New(slog.NewTextHandler(w, nil))
	mux := http.NewServeMux()
	mux.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		logger.Info("/set", query)
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

		fmt.Println(k, values)
		dbState[k] = values[0]
	})
	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		logger.Info("/get", query)
		if len(query) != 1 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "must get exactly one key using /get?somekey")
			return
		}
		k, values := getKeyValue(query)
		fmt.Println(query.Has(k))
		fmt.Println(k, values)
		// TODO len(values)=1 with empty string on get?a
		// if len(values) > 0 {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	fmt.Fprintf(w, "get must be called without a value, did you mean to set using /set?somekey=somevalue")
		// 	return
		// }

		v, ok := dbState[k]
		fmt.Println(v, ok)

		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "%s=%s", k, v)
	})
	srv := http.Server{
		Addr:    "127.0.0.1:4000",
		Handler: mux,
	}
	return srv.ListenAndServe()
}

func getKeyValue(query url.Values) (string, []string) {
	for k, v := range query {
		return k, v
	}
	return "", nil
}
