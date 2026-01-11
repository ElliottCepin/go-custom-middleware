package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Calculation struct {
	A int
	B int
	Operation string
}

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

func main() {
	http.HandleFunc("/calculate", serveCalculate)
	http.ListenAndServe(":8080", nil)
}
