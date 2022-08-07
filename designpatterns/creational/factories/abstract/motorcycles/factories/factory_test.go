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
	t.Parallel()

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
		name, data := name, data

		t.Run(name, func(subtest *testing.T) {
			subtest.Parallel()
			factory, err := factories.NewUrbanFactory(data.factory)
			if data.err != err {
				subtest.Errorf("want err: %+v, but got: %+v", data.err, err)
				subtest.FailNow()
			}
			if data.err != nil {
				return
			}
			got := factory.CreateUrban()
			errStartEngine := got.StartEngine()
			if errStartEngine != nil {
				subtest.Fatalf("unexpected error: %s", errStartEngine)
			}
			errSpeedUp := got.SpeedUp(data.speedToAdd)
			if errSpeedUp != nil {
				subtest.Fatalf("unexpected error: %s", errSpeedUp)
			}

			assert.Equal(subtest, data.want, got)
		})
	}
}

func TestCreateSportBike(t *testing.T) {
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
		name, data := name, data
		t.Run(name, func(subtest *testing.T) {
			subtest.Parallel()
			factory, err := factories.NewSportFactory(data.factory)
			if data.err != err {
				subtest.Errorf("want err: %+v, but got: %+v", data.err, err)
				subtest.FailNow()
			}
			if data.err != nil {
				return
			}
			got := factory.CreateSport()
			errStartEngine := got.StartEngine()
			if errStartEngine != nil {
				subtest.Fatalf("unexpected error: %s", errStartEngine)
			}
			errSpeedUp := got.SpeedUp(data.speedToAdd)
			if errSpeedUp != nil {
				subtest.Fatalf("unexpected error: %s", errSpeedUp)
			}
			errIncreasePower := got.IncreasePower(data.power)
			if errIncreasePower != nil {
				subtest.Fatalf("unexpected error: %s", errIncreasePower)
			}
			errActivateDrivingMode := got.ActivateDrivingMode(data.driveMode)
			if errActivateDrivingMode != nil {
				subtest.Fatalf("unexpected error: %s", errActivateDrivingMode)
			}

			assert.Equal(subtest, data.want, got)
		})
	}
}
