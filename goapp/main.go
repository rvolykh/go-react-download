package main

import (
	"encoding/csv"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	http.HandleFunc("/download", DownloadHandleFunc)

	log.Fatalln(http.ListenAndServe("0.0.0.0:80", nil))
}

func DownloadHandleFunc(w http.ResponseWriter, r *http.Request) {
	pr, pw := io.Pipe()

	go func() {
		w := csv.NewWriter(pw)

		defer safeClose(pw.Close)
		defer w.Flush()

		generateCSV(w, rand.Intn(500*1000))
	}()
	defer safeClose(pr.Close)

	w.Header().Set("Content-Disposition", "attachment; filename=result.csv")
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Transfer-Encoding", "chunked")

	if _, err := io.Copy(w, pr); err != nil {
		log.Printf("ERROR: copy: %v", err)
	}
}

func generateCSV(w *csv.Writer, rowNumber int) {
	if err := w.Write([]string{"header1", "header2", "header3", "header4", "header5"}); err != nil {
		log.Printf("ERROR: header: %v", err)
		return
	}

	for i := 0; i < rowNumber; i++ {
		row := []string{
			randomString(),
			randomString(),
			randomString(),
			randomString(),
			randomString(),
		}

		if err := w.Write(row); err != nil {
			log.Printf("ERROR: row: %v", err)
			return
		}
	}

	if err := w.Write([]string{"footer1", "footer2", "footer3", "footer4", "footer5"}); err != nil {
		log.Printf("ERROR: footer: %v", err)
		return
	}
}

func safeClose(f func() error) {
	if err := f(); err != nil {
		log.Printf("ERROR: close: %v", err)
	}
}

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString() string {
	b := make([]rune, rand.Intn(50)+1)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
