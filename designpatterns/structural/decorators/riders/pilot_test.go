package riders_test

import (
	"testing"

	"github.com/fernandoocampo/micro-patterns/designpatterns/structural/decorators/helmets"
	"github.com/fernandoocampo/micro-patterns/designpatterns/structural/decorators/helmets/comforts"
	"github.com/fernandoocampo/micro-patterns/designpatterns/structural/decorators/helmets/liners"
	"github.com/fernandoocampo/micro-patterns/designpatterns/structural/decorators/helmets/shells"
	shields "github.com/fernandoocampo/micro-patterns/designpatterns/structural/decorators/helmets/shield"
	"github.com/fernandoocampo/micro-patterns/designpatterns/structural/decorators/riders"
	"github.com/stretchr/testify/assert"
)

func TestDressRider(t *testing.T) {
	t.Parallel()
	// Given
	expectedDamageReport := riders.DamageReport{
		Head: 299.25,
	}

	impactForceLevel := float32(2000.0)

	var helmetLayer helmets.Protector

	helmet := helmets.New()
	helmetLayer = liners.NewWithoutProtector()
	helmetLayer = comforts.New(helmetLayer)
	helmetLayer = shells.New(helmetLayer)
	helmetLayer = shields.New(helmetLayer)
	helmet.SetProtection(helmetLayer)
	newRider := riders.New("new rider").WithHelmet(helmet)
	// When
	got := newRider.Crash(impactForceLevel)
	// Then
	assert.Equal(t, expectedDamageReport, got)
}
