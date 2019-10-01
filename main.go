package main

import (
	"fmt"      // пакет для форматированного ввода вывода
	"log"      // пакет для логирования
	"net/http" // пакет для поддержки HTTP протокола

	"github.com/mholt/certmagic"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloServer)

	go func() {
		if err := http.ListenAndServe(":80", http.HandlerFunc(redirectTLS)); err != nil {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	err := certmagic.HTTPS([]string{"play-with-lefthook.tk", "www.play-with-lefthook.tk"}, mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
}

func hello() string {
	return "Hello world v0.9.8"
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, hello())
}
