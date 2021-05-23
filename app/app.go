package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":9000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello go from Let's Learn, %s!\n", r.URL.Path[1:])
	fmt.Fprintf(w, "Hello world from Anand\n")
	fmt.Fprintf(w, "Hi Everyone, guess who\n")
}
