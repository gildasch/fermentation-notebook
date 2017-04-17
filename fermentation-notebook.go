package main

import (
	"fmt"
	"io/ioutil"
	"os"

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

	fmt.Println(bs)
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
