package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := button("Ifan")
	http.Handle("/button", templ.Handler(component))

	index_page := index()
	http.Handle("/", templ.Handler(index_page))
	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
