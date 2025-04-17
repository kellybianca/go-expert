package main

import "net/http"

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World 2.0!"))
	})
	http.ListenAndServe(":8080", mux)
}
