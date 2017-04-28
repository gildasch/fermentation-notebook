package model

import (
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
	Duration string
}

func ParseMethods(input []byte) (ms Methods, err error) {
	err = yaml.Unmarshal(input, &ms)
	if err != nil {
		return
	}

	for _, m := range ms.Methods {
		m.Durations = make(map[string]time.Duration)
		for _, s := range m.Steps {
			// Only the first
			if _, ok := m.Durations[s.Name]; !ok {
				m.Durations[s.Name], _ = time.ParseDuration(s.Duration)
			}
		}
	}

	return
}
