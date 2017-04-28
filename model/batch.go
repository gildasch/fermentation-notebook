package model

import (
	"fmt"
	"sort"
	"time"

	yaml "gopkg.in/yaml.v2"
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

func ParseBatches(input []byte, ms Methods) (bs Batches, err error) {
	err = yaml.Unmarshal(input, &bs)
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
