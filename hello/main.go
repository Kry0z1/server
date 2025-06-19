package main

import "net/http"

func main() {
	http.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello changed"))
	}))
	http.ListenAndServe(":8080", nil)
}
