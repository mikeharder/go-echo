package main

import (
    "net/http"
    "io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
    contentLength := r.ContentLength

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, http.StatusText(500), 500)
    }

    bytes, err2 := w.Write(body)
    if err2 != nil {
        http.Error(w, http.StatusText(500), 500)
    }
    if int64(bytes) != contentLength {
        http.Error(w, http.StatusText(500), 500)
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
