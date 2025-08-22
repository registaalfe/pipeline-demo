package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/add", handler)

	fmt.Println("Server start on port 8080")
	http.ListenAndServe(":8080", nil)
}

func handler(responseWriter http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()
	valA := values.Get("a")
	valB := values.Get("b")

	a, err := strconv.Atoi(valA)
	if err != nil {
		http.Error(responseWriter, "a is not number", http.StatusBadRequest)
		return
	}

	b, err := strconv.Atoi(valB)
	if err != nil {
		http.Error(responseWriter, "b is not number", http.StatusBadRequest)
		return
	}

	result := a + b

	fmt.Fprintf(responseWriter, "%d", result)
}
