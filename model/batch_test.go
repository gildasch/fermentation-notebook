package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseBatches(t *testing.T) {
	ms, err := ParseMethods([]byte(`methods:
  Levain solide:
    steps:
      - name: Création
      - name: Levage
        duration: 96h
      - name: Rafraichi
`))
	require.NoError(t, err)
	assert.Equal(t,
		Methods{
			Methods: map[string]*Method{
				"Levain solide": &Method{
					Steps: []Step{
						Step{
							Name: "Création",
						},
						Step{
							Name:     "Levage",
							Duration: "96h",
						},
						Step{
							Name: "Rafraichi",
						},
					},
					Durations: map[string]time.Duration{
						"Levage": 96 * time.Hour,
					},
				},
			}},
		ms)

	bs, err := ParseBatches([]byte(`batches:
  - name: L1
    type: Levain solide
    history:
      - name: Création
        date: 2017-04-17 16:00
        comment: |
          300g farine, T65 à la meule de pierre, bio
          150g eau cristaline, 50% temp. ambiante, 50% 65 deg.
          15g miel, non bio

          Sources: Chritophe (Levain, Le vin), Le grand manuel du boulanger
      - name: Levage
        date: 2017-04-17 16:01
`), ms)
	require.NoError(t, err)
	assert.Equal(t,
		Batches{
			Batches: []*Batch{&Batch{
				Name: "L1",
				Type: "Levain solide",
				History: []Event{
					Event{
						Name: "Création",
						Date: "2017-04-17 16:00",
						Comment: `300g farine, T65 à la meule de pierre, bio
150g eau cristaline, 50% temp. ambiante, 50% 65 deg.
15g miel, non bio

Sources: Chritophe (Levain, Le vin), Le grand manuel du boulanger
`,
					},
					Event{
						Name: "Levage",
						Date: "2017-04-17 16:01",
					},
				},
				NextEvent: time.Date(2017, 04, 21, 16, 1, 0, 0, time.UTC),
			}}},
		bs)
}
