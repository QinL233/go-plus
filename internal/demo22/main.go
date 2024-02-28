package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Name struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "got root\n")
	})
	mux.HandleFunc("GET /query", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "got path\n")
	})
	mux.HandleFunc("/path/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "handling task with id=%v\n", id)
	})
	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		var name Name
		json.NewDecoder(r.Body).Decode(&name)
		fmt.Fprintf(w, "%v", name)
	})
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		http.NewResponseController(w)
	})
	mux.HandleFunc("POST /form", func(w http.ResponseWriter, r *http.Request) {
		f, _ := r.MultipartReader()
		p, _ := f.NextPart()

		for k, v := range p.Header {
			vs := ""
			for _, s := range v {
				vs += s
			}
			fmt.Fprintf(w, "%s: %s\r\n", k, vs)
		}
	})
	http.ListenAndServe("localhost:8090", mux)

}
