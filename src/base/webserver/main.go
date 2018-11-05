package main

import (
    "io"
    "net/http"
)


func main() {
    http.HandleFunc("/",firstPage)
    http.ListenAndServe(":8000",nil)
}

func firstPage(w http.ResponseWriter,r *http.Request) {
    io.WriteString(w,"<h1>hello, my first go webserver</h1>")
}

