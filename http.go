package main

import (
	"html/template"
	"log"
	"net/http"
)

func webserver() {
	http.HandleFunc("/", handleMain)

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	err := http.ListenAndServe(*flagHTTPPort, nil) //nil의 정확한 용도를 이슈로 만들어봅시당.
	if err != nil {
		log.Fatal(err)
	}
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	i, err := template.ParseFiles(
		"assets/html/header.html",
		"assets/html/index.html",
		"assets/html/footer.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	err = i.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Fatal(err)
	}
}
