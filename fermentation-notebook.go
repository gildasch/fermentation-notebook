package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"

	"gopkg.in/yaml.v2"
)

type Batches struct {
	Batches []Batch
}

type Batch struct {
	Name    string
	Type    string
	History []Event
}

type Event struct {
	Name    string
	Date    string
	Comment string
}

type Methods struct {
}

func main() {
	bs, err := readBatches("batches.yaml")
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(bs)

	err = serve(bs)
	fmt.Println(err)
}

func readBatches(input string) (bs Batches, err error) {
	bf, err := os.Open(input)
	if err != nil {
		return
	}
	bb, err := ioutil.ReadAll(bf)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(bb, &bs)
	return
}

func serve(bs Batches) error {
	batchesHandler := func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("batches.html").Funcs(template.FuncMap{
			"nl2br": func(s string) template.HTML {
				return template.HTML(strings.Replace(s, "\n", "<br />\n", -1))
			}}).ParseFiles("tmpl/batches.html")
		if err != nil {
			fmt.Println(err)
		}
		err = t.Execute(w, bs)
		if err != nil {
			fmt.Println(err)
		}
	}

	r := mux.NewRouter()
	r.HandleFunc("/", batchesHandler).Methods("GET")

	http.Handle("/", r)
	port := "8080"
	fmt.Printf("Listening on %s...\n", port)
	return http.ListenAndServe(":"+port, nil)
}
