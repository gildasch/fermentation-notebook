package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/GildasCh/fermentation-notebook/model"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:", os.Args[0], "path/to/batches.yaml", "path/to/methods.yaml")
		return
	}

	ms, err := model.ReadMethods(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	bs, err := model.ReadBatches(os.Args[1], ms)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = serve(bs, ms)
	fmt.Println(err)
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
				ret := time.Until(t)
				if ret > 2*time.Hour {
					ret -= ret % time.Hour
				} else if ret > 2*time.Minute {
					ret -= ret % time.Minute
				} else {
					ret -= ret % time.Second
				}
				return strings.Replace(strings.Replace(ret.String(), "0s", "", 1), "0m", "", 1)
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
