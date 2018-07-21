package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gildasch/fermentation-notebook/batches"
	"github.com/gorilla/mux"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/batches.yaml")
		return
	}

	bs, err := readBatches(os.Args[1])
	if err != nil {
		fmt.Println("Error reading batches:", err)
		return
	}

	err = serve(bs)
	fmt.Println(err)
}

func readBatches(input string) (bs batches.Batches, err error) {
	bf, err := os.Open(input)
	if err != nil {
		return
	}
	bb, err := ioutil.ReadAll(bf)
	if err != nil {
		return
	}

	bs, err = batches.ParseBatches(bb)
	if err != nil {
		return
	}
	return
}

func serve(bs batches.Batches) error {
	batchesHandler := func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("batches.html").Funcs(template.FuncMap{
			"nl2br": func(s string) template.HTML {
				return template.HTML(strings.Replace(s, "\n", "<br />\n", -1))
			}}).ParseFiles("tmpl/batches.html")
		if err != nil {
			fmt.Println(err)
		}
		err = t.Execute(w, struct {
			batches.Batches
		}{bs})
		if err != nil {
			fmt.Println(err)
		}
	}

	r := mux.NewRouter()
	r.HandleFunc("/", batchesHandler).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.Handle("/", r)
	port := "8080"
	fmt.Printf("Listening on %s...\n", port)
	return http.ListenAndServe(":"+port, nil)
}
