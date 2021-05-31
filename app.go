package app

import (
    "fmt"
    "net/http"
    g "./generator"
    "encoding/json"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":9001", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
     res := g.Game()
	 res1,_  :=json.MarshalIndent(res,""," ")

    fmt.Fprintf(w, string(res1), r.URL.Path[1:])
}
