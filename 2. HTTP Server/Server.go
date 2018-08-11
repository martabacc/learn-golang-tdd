package main

import (
    "fmt"
    "net/http"
)

func returnOK(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", "dinolai")
}

/*
 * Returning ServeMux instance
 * which contains all set ServeMux
 */

func getRouters() *http.ServeMux {
    router := http.NewServeMux()
    router.HandleFunc("/", returnOK)

    return router
}

func main() {
    http.ListenAndServe("localhost:8", getRouters())
}
