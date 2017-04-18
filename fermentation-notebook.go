package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"gopkg.in/yaml.v2"
)

type Batches struct {
	Batches []*Batch
}

func (b Batches) Len() int      { return len(b.Batches) }
func (b Batches) Swap(i, j int) { b.Batches[i], b.Batches[j] = b.Batches[j], b.Batches[i] }
func (b Batches) Less(i, j int) bool {
	return b.Batches[i].NextEvent.Before(b.Batches[j].NextEvent)
}

type Batch struct {
	Name      string
	Type      string
	History   []Event
	NextEvent time.Time
}

type Event struct {
	Name    string
	Date    string
	Comment string
}

type Methods struct {
	Methods map[string]*Method
}

type Method struct {
	Name      string
	Steps     []Step
	Durations map[string]time.Duration
}

type Step struct {
	Name     string
	Duration time.Duration
}

func main() {
	ms, err := readMethods("methods.yaml")
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(ms)

	bs, err := readBatches("batches.yaml", ms)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(bs)

	err = serve(bs, ms)
	fmt.Println(err)
}

func serve(bs Batches, ms Methods) error {
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

func readBatches(input string, ms Methods) (bs Batches, err error) {
	bf, err := os.Open(input)
	if err != nil {
		return
	}
	bb, err := ioutil.ReadAll(bf)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(bb, &bs)
	if err != nil || len(bs.Batches) == 0 {
		return
	}

	for _, b := range bs.Batches {
		lastEvent := b.History[len(b.History)-1]
		lastTime, err := time.Parse("2006-01-02 15:04", lastEvent.Date)
		fmt.Println(lastEvent.Date, lastTime, err)
		if err != nil {
			continue
		}
		b.NextEvent = lastTime.Add(ms.Methods[b.Type].Durations[lastEvent.Name])
		fmt.Println(lastTime, "+", ms.Methods[b.Type].Durations[lastEvent.Name], b.NextEvent)
	}

	sort.Sort(bs)

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
	if err != nil {
		return
	}

	for _, m := range ms.Methods {
		m.Durations = make(map[string]time.Duration)
		for _, s := range m.Steps {
			// Only the first
			if _, ok := m.Durations[s.Name]; !ok {
				m.Durations[s.Name] = s.Duration
			}
		}
	}

	return
}
