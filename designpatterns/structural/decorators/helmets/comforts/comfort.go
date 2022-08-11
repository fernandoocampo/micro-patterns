package comforts

// Protector defines behavior to protect the rider.
type Protector interface {
	// Protect signals that the protector must protect and
	// return the new impact force level.
	Protect(impactForceLevel float32) float32
}

// Comfort contains comfort layer data.
type Comfort struct {
	nextProtector Protector
}

const levelOfProtectionPerImpact = 0.1

func New(protector Protector) *Comfort {
	newBoot := Comfort{
		nextProtector: protector,
	}

	return &newBoot
}

func (c *Comfort) Protect(impactForceLevel float32) float32 {
	if impactForceLevel <= 0 {
		return 0.0
	}

	impactForceLevel -= (impactForceLevel * levelOfProtectionPerImpact)
	if c.nextProtector == nil {
		return impactForceLevel
	}

	return c.nextProtector.Protect(impactForceLevel)
}
