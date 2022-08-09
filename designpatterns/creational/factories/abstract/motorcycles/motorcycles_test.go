package motorcycles_test

import (
	"errors"
	"testing"

	assert "github.com/stretchr/testify/assert"

	motorcycles "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles"
	sportbmw "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/bmw/sports"
	urbanbmw "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/bmw/urbans"
	sportducati "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/ducati/sports"
	urbanducati "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/ducati/urbans"

	bmw "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/bmw"
	ducati "github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/ducati"
)

type testSportData struct {
	want       motorcycles.SportBehavior
	factory    brand
	driveMode  string
	power      float32
	speedToAdd float32
	err        error
}

func TestCreateUrbanMotorcycles(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		want       motorcycles.UrbanBehavior
		factory    brand
		speedToAdd float32
		err        error
	}{
		"unknown_factory": {
			want:    nil,
			factory: brand(6),
			err:     errUnknownBrand,
		},
		"urban_bmw": {
			want: &urbanbmw.Scooter{
				EngineOn: true,
				Speed:    34.5,
				Moving:   true,
			},
			speedToAdd: 34.5,
			factory:    BMW,
			err:        nil,
		},
		"urban_ducati": {
			want: &urbanducati.Scooter{
				EngineOn: true,
				Speed:    55.5,
				Moving:   true,
			},
			speedToAdd: 55.5,
			factory:    Ducati,
			err:        nil,
		},
	}

	for name, data := range cases {
		name, data := name, data
		t.Run(name, func(subtest *testing.T) {
			subtest.Parallel()
			factory, err := newFactory(data.factory)
			if !errors.Is(data.err, err) {
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
			factory:    BMW,
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
			factory:    Ducati,
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
	factory, err := newFactory(data.factory)
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

// brand defines the motorcycle brand.
type brand int32

// factory defines behavior to create motorcycle products.
type factory interface {
	CreateUrban() motorcycles.UrbanBehavior
	CreateSport() motorcycles.SportBehavior
	CreateAdventure() motorcycles.AdventureBehavior
}

// Supported motorcycle brands.
const (
	BMW brand = iota
	Ducati
)

var errUnknownBrand = errors.New("unknown brand")

// newFactory creates a factory to create motorcycle products.
func newFactory(factory brand) (factory, error) {
	switch factory {
	case BMW:
		return bmw.NewFactory(), nil
	case Ducati:
		return ducati.NewFactory(), nil
	default:
		return nil, errUnknownBrand
	}
}
