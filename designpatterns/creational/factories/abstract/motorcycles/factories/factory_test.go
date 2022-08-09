package factories_test

import (
	"errors"
	"testing"

	assert "github.com/stretchr/testify/assert"

	motorcycles "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles"
	sportbmw "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/bmw/sports"
	urbanbmw "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/bmw/urbans"
	sportducati "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/ducati/sports"
	urbanducati "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/ducati/urbans"

	factories "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/factories"
)

type testSportData struct {
	want       motorcycles.SportBehavior
	factory    motorcycles.Brand
	driveMode  string
	power      float32
	speedToAdd float32
	err        error
}

type testUrbanData struct {
	want       motorcycles.UrbanBehavior
	factory    motorcycles.Brand
	speedToAdd float32
	err        error
}

func TestCreateUrbanMotorcycles(t *testing.T) {
	t.Parallel()

	cases := map[string]testUrbanData{
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

			got := createAndTestUrbanEngine(t, data)

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

			got := createAndTestSportEngine(t, data)

			assert.Equal(subtest, data.want, got)
		})
	}
}

func createAndTestUrbanEngine(t *testing.T, data testUrbanData) motorcycles.UrbanBehavior {
	factory, err := factories.NewUrbanFactory(data.factory)
	if !errors.Is(data.err, err) {
		t.Errorf("want err: %+v, but got: %+v", data.err, err)
		t.FailNow()
	}

	if data.err != nil {
		return nil
	}

	got := factory.CreateUrban()

	errStartEngine := got.StartEngine()
	if errStartEngine != nil {
		t.Fatalf("unexpected error: %s", errStartEngine)
	}

	errSpeedUp := got.SpeedUp(data.speedToAdd)
	if errSpeedUp != nil {
		t.Fatalf("unexpected error: %s", errSpeedUp)
	}

	return got
}

func createAndTestSportEngine(t *testing.T, data testSportData) motorcycles.SportBehavior {
	t.Helper()

	factory, err := factories.NewSportFactory(data.factory)
	if !errors.Is(data.err, err) {
		t.Errorf("want err: %+v, but got: %+v", data.err, err)
		t.FailNow()
	}

	if data.err != nil {
		return nil
	}

	got := factory.CreateSport()

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
