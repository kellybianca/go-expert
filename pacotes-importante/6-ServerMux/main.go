package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello, World!"))
	// })
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", &Blog{Title: "My Blog"})
	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

type Blog struct {
	Title string
}

func (b *Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.Title))
}
