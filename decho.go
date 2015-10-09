package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func echo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Echo!")
	fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
}

func main() {
        fmt.Printf("Starting gregbuehler/decho using DECHO_BIND '%s'\n", os.Getenv("DECHO_BIND"))

	http.HandleFunc("/", echo)
	http.ListenAndServe(os.Getenv("DECHO_BIND")+":8765", nil)
}
