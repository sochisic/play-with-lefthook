package main

import (
	"flag"
	"fmt"      // пакет для форматированного ввода вывода
	"log"      // пакет для логирования
	"net/http" // пакет для поддержки HTTP протокола

	"github.com/mholt/certmagic"
)

func main() {
	port := flag.String("port", "3000", "an int")

	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloServer)

	fmt.Printf("Listening port %v...", *port)
	// err := http.ListenAndServe(":"+*port, mux) // задаем слушать порт
	err := certmagic.HTTPS([]string{"ec2-13-48-30-65.eu-north-1.compute.amazonaws.com"}, mux) // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func hello() string {
	return "Hello world v0.9.0"
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, hello())
}
