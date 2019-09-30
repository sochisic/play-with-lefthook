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

	// // fmt.Println("Listening port ...")

	// cache := certmagic.NewCache(certmagic.CacheOptions{

	// })

	// magic := certmagic.New(certmagic.Config{

	// 	Email:          "460sochi@gmail.com",
	// 	Agreed:         true,
	// 	AltHTTPPort:    80,
	// 	AltTLSALPNPort: 443,
	// })

	// err := magic.Manage([]string{"play-with-lefthook.tk", "www.play-with-lefthook.tk"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// s := http.Server{
	// 	Addr:      "443",
	// 	Handler:   mux,
	// 	TLSConfig: magic.TLSConfig(),
	// }

	// log.Println("listening...")
	// go log.Fatal(http.ListenAndServe(":80", magic.HTTPChallengeHandler(mux)))
	// log.Fatal(s.ListenAndServeTLS("", ""))

	err := certmagic.HTTPS([]string{"play-with-lefthook.tk", "www.play-with-lefthook.tk"}, mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	// ln, err := certmagic.Listen([]string{"play-with-lefthook.tk", "www.play-with-lefthook.tk"})
	// if err != nil {
	// 	return err
	// }

	// tlsConfig, err := certmagic.TLS([]string{"play-with-lefthook.tk", "www.play-with-lefthook.tk"})
	// if err != nil {
	// 	return err
	// }
}

func hello() string {
	return "Hello world v0.9.4"
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, hello())
}
