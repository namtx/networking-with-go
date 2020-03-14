package main

import "net/http"

func main() {
	myHandler := http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	http.ListenAndServe(":8000", myHandler)
}
