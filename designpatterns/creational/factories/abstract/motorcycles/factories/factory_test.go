package factories_test

import (
	"testing"

	"github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles"
	sportbmw "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/bmw/sports"
	urbanbmw "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/bmw/urbans"
	sportducati "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/ducati/sports"
	urbanducati "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/ducati/urbans"
	"github.com/stretchr/testify/assert"

	"github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/factories"
)

func TestCreateUrbanMotorcycles(t *testing.T) {
	cases := map[string]struct {
		want       motorcycles.UrbanBehavior
		factory    motorcycles.Brand
		speedToAdd float32
		err        error
	}{
		"unknown_factory": {
			want:    nil,
			factory: motorcycles.Brand(6),
			err:     factories.ErrUnknowBrand,
		},
		"urban_bmw": {
			want: &urbanbmw.Scooter{
				EngineOn: true,
				Speed:    34.5,
				Moving:   true,
			},
			speedToAdd: 34.5,
			factory:    factories.BMW,
			err:        nil,
		},
		"urban_ducati": {
			want: &urbanducati.Scooter{
				EngineOn: true,
				Speed:    55.5,
				Moving:   true,
			},
			speedToAdd: 55.5,
			factory:    factories.Ducati,
			err:        nil,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			factory, err := factories.NewUrbanFactory(data.factory)
			if data.err != err {
				st.Errorf("want err: %+v, but got: %+v", data.err, err)
				st.FailNow()
			}
			if data.err != nil {
				return
			}
			got := factory.CreateUrban()
			got.StartEngine()
			got.SpeedUp(data.speedToAdd)

			assert.Equal(st, data.want, got)
		})
	}
}

func TestCreateSportMotorcycles(t *testing.T) {
	cases := map[string]struct {
		want       motorcycles.SportBehavior
		factory    motorcycles.Brand
		driveMode  string
		power      float32
		speedToAdd float32
		err        error
	}{
		"sport_bmw": {
			want: &sportbmw.Sport{
				EngineOn:    true,
				Speed:       34.5,
				DrivingMode: "rain",
				Power:       200,
				Moving:      true,
			},
			driveMode:  "rain",
			power:      200,
			speedToAdd: 34.5,
			factory:    factories.BMW,
			err:        nil,
		},
		"sport_ducati": {
			want: &sportducati.Sport{
				EngineOn:    true,
				Speed:       55.5,
				DrivingMode: "race",
				Power:       250,
				Moving:      true,
			},
			driveMode:  "race",
			power:      250,
			speedToAdd: 55.5,
			factory:    factories.Ducati,
			err:        nil,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			factory, err := factories.NewSportFactory(data.factory)
			if data.err != err {
				st.Errorf("want err: %+v, but got: %+v", data.err, err)
				st.FailNow()
			}
			if data.err != nil {
				return
			}
			got := factory.CreateSport()
			got.StartEngine()
			got.SpeedUp(data.speedToAdd)
			got.IncreasePower(data.power)
			got.ActivateDrivingMode(data.driveMode)

			assert.Equal(st, data.want, got)
		})
	}
}
