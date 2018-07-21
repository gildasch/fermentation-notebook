package batches

import (
	yaml "gopkg.in/yaml.v2"
)

type Batches struct {
	Current []*Batch
	Past    []*Batch
}

type Batch struct {
	Name string
	Type string
	Log  []Action
}

type Action struct {
	Name        string
	Date        string
	Ingredients []Ingredient
	Notes       string
}

type Ingredient struct {
	Type     string
	Quantity string
	Brand    string
}

func ParseBatches(input []byte) (bs Batches, err error) {
	err = yaml.Unmarshal(input, &bs)
	return
}
