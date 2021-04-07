package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ganinugroho/belajar/handlers"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hai semua")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Ooops", http.StatusBadRequest)
			return
		}
		log.Printf("Data %s\n", d)
		fmt.Fprintf(rw, "hallo %s ", d)
	})
	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) {
		log.Println("Selamat tinggal")
	})

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	sm := http.NewServeMux()
	sm.Handle("/hello", hh)

	http.ListenAndServe(":9090", nil)
}
