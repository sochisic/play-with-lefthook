package main

import (
	"flag"
	"fmt"      // пакет для форматированного ввода вывода
	"log"      // пакет для логирования
	"net/http" // пакет для поддержки HTTP протокола
	// пакет для работы с  UTF-8 строками
)

func main() {
	port := flag.String("port", "3000", "an int")

	flag.Parse()

	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":"+*port, nil) // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func hello() string {
	return "Hello world v0.3.4"
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, hello())
}
