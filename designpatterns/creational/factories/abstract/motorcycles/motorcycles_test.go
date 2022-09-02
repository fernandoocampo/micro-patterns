package motorcycles_test

import (
	"testing"

	motorcycles "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles"
	bmw "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/bmw"
	sportbmw "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/bmw/sports"
	urbanbmw "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/bmw/urbans"
	ducati "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/ducati"
	sportducati "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/ducati/sports"
	urbanducati "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/ducati/urbans"
	assert "github.com/stretchr/testify/assert"
)

type testSportData struct {
	want       motorcycles.SportBehavior
	factory    factory
	driveMode  string
	power      float32
	speedToAdd float32
	err        error
}

type testUrbanData struct {
	want       motorcycles.UrbanBehavior
	factory    factory
	err        error
	speedToAdd float32
}

func TestCreateUrbanMotorcycles(t *testing.T) {
	t.Parallel()

	cases := map[string]testUrbanData{
		"urban_bmw": {
			want: &urbanbmw.Scooter{
				EngineOn: true,
				Speed:    34.5,
				Moving:   true,
			},
			speedToAdd: 34.5,
			factory:    bmw.NewFactory(),
			err:        nil,
		},
		"urban_ducati": {
			want: &urbanducati.Scooter{
				EngineOn: true,
				Speed:    55.5,
				Moving:   true,
			},
			speedToAdd: 55.5,
			factory:    ducati.NewFactory(),
			err:        nil,
		},
	}

	for name, data := range cases {
		name, data := name, data
		t.Run(name, func(subtest *testing.T) {
			subtest.Parallel()
			got := data.factory.CreateUrban()
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
	t.Parallel()

	cases := map[string]testSportData{
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
			factory:    bmw.NewFactory(),
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
			factory:    ducati.NewFactory(),
			err:        nil,
		},
	}
	for name, data := range cases {
		name, data := name, data
		t.Run(name, func(subtest *testing.T) {
			subtest.Parallel()

			got := createAndTestSportBike(t, data)

			assert.Equal(subtest, data.want, got)
		})
	}
}

func createAndTestSportBike(t *testing.T, data testSportData) motorcycles.SportBehavior {
	got := data.factory.CreateSport()

	errStartEngine := got.StartEngine()
	if errStartEngine != nil {
		t.Fatalf("unexpected error: %s", errStartEngine)
	}

	errSpeedUp := got.SpeedUp(data.speedToAdd)
	if errSpeedUp != nil {
		t.Fatalf("unexpected error: %s", errSpeedUp)
	}

	errIncreasePower := got.IncreasePower(data.power)
	if errIncreasePower != nil {
		t.Fatalf("unexpected error: %s", errIncreasePower)
	}

	errActivateDrivingMode := got.ActivateDrivingMode(data.driveMode)
	if errActivateDrivingMode != nil {
		t.Fatalf("unexpected error: %s", errActivateDrivingMode)
	}

	return got
}

// factory defines behavior to create motorcycle products.
type factory interface {
	CreateUrban() motorcycles.UrbanBehavior
	CreateSport() motorcycles.SportBehavior
	CreateAdventure() motorcycles.AdventureBehavior
}
