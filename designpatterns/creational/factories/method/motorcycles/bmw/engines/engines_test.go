package engines_test

import (
	"testing"

	"github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/method/motorcycles/bmw/engines"
	"github.com/stretchr/testify/assert"
)

func TestCreateEngines(t *testing.T) {
	cases := map[string]struct {
		want     *engines.Urban
		input    engines.BuildData
		category engines.Creator
	}{
		"urban_engine": {
			want: &engines.Urban{
				Power:     32,
				Type:      engines.TwoStrokes,
				Cylinders: 1,
				Torque:    10,
			},
			input: engines.BuildData{
				Power:     32,
				Type:      engines.TwoStrokes,
				Cylinders: 1,
				Torque:    10,
			},
			category: engines.NewUrbanCreator(),
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := engines.NewFactory(data.category).BuildEngine(data.input)
			assert.Equal(t, data.want, got)
		})
	}
}
