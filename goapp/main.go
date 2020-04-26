package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/download", DownloadHandleFunc)
	http.ListenAndServe("0.0.0.0:80", nil)
}

func DownloadHandleFunc(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v", r)
	
	buf := strings.NewReader("asd")
	http.ServeContent(w, r, "blabla.txt", time.Now(), buf)
}
