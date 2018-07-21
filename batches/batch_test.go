package batches

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseBatches(t *testing.T) {
	b, err := ioutil.ReadFile("../batches.yaml")
	require.NoError(t, err)

	actual, err := ParseBatches(b)
	require.NoError(t, err)

	expected := Batches{
		Current: []*Batch{
			&Batch{
				Name: "Hydromel #1",
				Type: "Hydromel",
				Log: []Action{
					Action{
						Name: "Initial",
						Date: "2018-07-21 12:00",
						Ingredients: []Ingredient{
							Ingredient{
								Type:     "Eau",
								Quantity: "700g",
								Brand:    "Cristaline"},
							Ingredient{
								Type:     "Miel",
								Quantity: "175g",
								Brand:    "Rucher du Morvan Fleurs Sauvages"},
							Ingredient{
								Type:     "Myrtilles",
								Quantity: "80g",
								Brand:    "Biocoop"}}}}},
			&Batch{Name: "Hydromel #2",
				Type: "Hydromel",
				Log: []Action{
					Action{
						Name: "Initial",
						Date: "2018-07-21 12:02",
						Ingredients: []Ingredient{
							Ingredient{
								Type:     "Eau",
								Quantity: "700g",
								Brand:    "Cristaline"},
							Ingredient{
								Type:     "Miel",
								Quantity: "175g",
								Brand:    "Rucher du Morvan Fleurs Sauvages"}},
						Notes: "pas beaucoup de place dans le haut du bocal..."}}},
			&Batch{
				Name: "Hydromel #3",
				Type: "Hydromel",
				Log: []Action{
					Action{
						Name: "Initial",
						Date: "2018-07-21 12:10",
						Ingredients: []Ingredient{
							Ingredient{
								Type:     "Eau",
								Quantity: "1600g",
								Brand:    "Cristaline"},
							Ingredient{
								Type:     "Miel",
								Quantity: "400g",
								Brand:    "Miel de mon enfance, Fleurs de Champagne, Renois Apiculteurs"}}}}}},
		Past: []*Batch{&Batch{
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
		}}}

	assert.Equal(t, expected, actual)
}
