package publicsites

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/sites", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}
