package riders

import "github.com/fernandoocampo/micro-patterns/designpatterns/structural/decorators/helmets"

type DamageReport struct {
	Head float32
}

type Rider struct {
	name   string
	helmet helmets.Protector
}

func New(name string) *Rider {
	newRider := new(Rider)
	newRider.name = name
	return newRider
}

func (r *Rider) WithHelmet(helmet helmets.Protector) *Rider {
	r.helmet = helmet
	return r
}

func (r *Rider) Crash(impactForceLevel float32) DamageReport {
	return DamageReport{
		Head: r.helmet.Protect(impactForceLevel),
	}
}
