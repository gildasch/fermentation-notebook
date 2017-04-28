package model

import (
	"io/ioutil"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"
)

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

func ReadMethods(input string) (ms Methods, err error) {
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
