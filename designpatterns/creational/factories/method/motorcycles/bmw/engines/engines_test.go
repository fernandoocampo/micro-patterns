package engines_test

import (
	"testing"

	"github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/method/motorcycles/bmw/engines"
	"github.com/stretchr/testify/assert"
)

func TestCreateEngines(t *testing.T) {
	t.Parallel()

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
		name, data := name, data

		t.Run(name, func(st *testing.T) {
			st.Parallel()
			got := engines.NewFactory(data.category).BuildEngine(data.input)
			assert.Equal(st, data.want, got)
		})
	}
}
