package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	chi "github.com/go-chi/chi/v5"
)

//go:embed all:uibuild
var svelteStatic embed.FS

func main() {

	s, err := fs.Sub(svelteStatic, "uibuild")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Handle("/*", http.FileServer(http.FS(s)))
	// r.Mount("/", http.FileServer(http.FS(svelteStatic)))

	fmt.Println("Running on port: 8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}
