package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gildasch/fermentation-notebook/model"
	"github.com/gildasch/fermentation-notebook/utils/durations"
	"github.com/gorilla/mux"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:", os.Args[0], "path/to/batches.yaml", "path/to/methods.yaml")
		return
	}

	ms, err := readMethods(os.Args[2])
	if err != nil {
		fmt.Println("Error reading methods:", err)
		return
	}

	bs, err := readBatches(os.Args[1], ms)
	if err != nil {
		fmt.Println("Error reading batches:", err)
		return
	}

	err = serve(bs, ms)
	fmt.Println(err)
}

func readMethods(input string) (ms model.Methods, err error) {
	mf, err := os.Open(input)
	if err != nil {
		return
	}
	mb, err := ioutil.ReadAll(mf)
	if err != nil {
		return
	}

	ms, err = model.ParseMethods(mb)
	if err != nil {
		return
	}
	return
}

func readBatches(input string, ms model.Methods) (bs model.Batches, err error) {
	bf, err := os.Open(input)
	if err != nil {
		return
	}
	bb, err := ioutil.ReadAll(bf)
	if err != nil {
		return
	}

	bs, err = model.ParseBatches(bb, ms)
	if err != nil {
		return
	}
	return
}

func serve(bs model.Batches, ms model.Methods) error {
	batchesHandler := func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("batches.html").Funcs(template.FuncMap{
			"nl2br": func(s string) template.HTML {
				return template.HTML(strings.Replace(s, "\n", "<br />\n", -1))
			},
			"date": func(t time.Time) string {
				return t.Format("2006-01-02 15:04")
			},
			"until": func(t time.Time) string {
				return durations.DurationToString(time.Until(t))
			}}).ParseFiles("tmpl/batches.html")
		if err != nil {
			fmt.Println(err)
		}
		err = t.Execute(w, struct {
			model.Batches
			model.Methods
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
