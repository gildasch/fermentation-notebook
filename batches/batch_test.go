package batches

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseBatches(t *testing.T) {
	bs, err := ParseBatches([]byte(`batches:
  - name: L1
    type: Levain solide
    log:
      - name: Création
        date: 2017-04-17 16:00
        notes: |
          300g farine, T65 à la meule de pierre, bio
          150g eau cristaline, 50% temp. ambiante, 50% 65 deg.
          15g miel, non bio

          Sources: Chritophe (Levain, Le vin), Le grand manuel du boulanger
      - name: Levage
        date: 2017-04-17 16:01
`))
	require.NoError(t, err)
	assert.Equal(t,
		Batches{
			Batches: []*Batch{&Batch{
				Name: "L1",
				Type: "Levain solide",
				Log: []Action{
					Action{
						Name: "Création",
						Date: "2017-04-17 16:00",
						Notes: `300g farine, T65 à la meule de pierre, bio
150g eau cristaline, 50% temp. ambiante, 50% 65 deg.
15g miel, non bio

Sources: Chritophe (Levain, Le vin), Le grand manuel du boulanger
`,
					},
					Action{
						Name: "Levage",
						Date: "2017-04-17 16:01",
					},
				},
			}}},
		bs)
}
