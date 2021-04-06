package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hai semua")
		d, _ := ioutil.ReadAll(r.Body)
		log.Printf("Data %s\n", d)
		fmt.Fprintf(rw, "hallo %s ", d)
	})
	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) {
		log.Println("Selamat tinggal")
	})
	http.ListenAndServe(":9090", nil)
}
