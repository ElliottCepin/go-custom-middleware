package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"strconv"
)

type Calculation struct {
	A int
	B int
	Operation string
}
var uuid int = 0

func serveCalculate(w http.ResponseWriter, r *http.Request) {
	if (r.Method != "POST") {
		fmt.Fprint(w, "There is nothing to see here")	
	} else {
		var calc Calculation
		err := json.NewDecoder(r.Body).Decode(&calc)
		if (err != nil) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if (calc.Operation == "add") {
			fmt.Fprint(w, calc.A + calc.B)
		} else if (calc.Operation == "multiply") {
			fmt.Fprint(w, calc.A * calc.B)
		} else if (calc.Operation == "subtract") {
			fmt.Fprint(w, calc.A - calc.B)
		} else {
			fmt.Fprint(w, "unrecognized operation")
		}
	}
}

func logger(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		fmt.Printf("Method: %q | Path: %q | Response Time: %v\n", r.Method, r.URL.Path, duration)
	}
}

func updateHeader(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Request-ID", strconv.Itoa(uuid))
		next.ServeHTTP(w, r)
	}
}

func main() {
	http.Handle("/calculate", logger(updateHeader(http.HandlerFunc(serveCalculate))))
	http.ListenAndServe(":8080", nil)
}
