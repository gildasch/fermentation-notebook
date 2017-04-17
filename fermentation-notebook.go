package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

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
	Methods []Method
}

type Method struct {
	Name  string
	Steps []Step
}

type Step struct {
	Name     string
	Duration time.Duration
}

func main() {
	bs, err := readBatches("batches.yaml")
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(bs)

	ms, err := readMethods("methods.yaml")
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(ms)

	err = serve(bs, ms)
	fmt.Println(err)
}

func serve(bs Batches, ms Methods) error {
	batchesHandler := func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("batches.html").Funcs(template.FuncMap{
			"nl2br": func(s string) template.HTML {
				return template.HTML(strings.Replace(s, "\n", "<br />\n", -1))
			}}).ParseFiles("tmpl/batches.html")
		if err != nil {
			fmt.Println(err)
		}
		err = t.Execute(w, struct {
			Batches
			Methods
		}{bs, ms})
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

func readMethods(input string) (ms Methods, err error) {
	mf, err := os.Open(input)
	if err != nil {
		return
	}
	mb, err := ioutil.ReadAll(mf)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(mb, &ms)
	return
}
