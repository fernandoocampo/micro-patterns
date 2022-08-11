package liners

// Protector defines behavior to protect the rider.
type Protector interface {
	// Protect signals that the protector must protect and
	// return the new impact force level.
	Protect(impactForceLevel float32) float32
}

// Liner contains liner layer data.
type Liner struct {
	nextProtector Protector
}

const levelOfProtectionPerImpact = 0.05

func New(protector Protector) *Liner {
	newBoot := Liner{
		nextProtector: protector,
	}

	return &newBoot
}

func NewWithoutProtector() *Liner {
	return new(Liner)
}

func (l *Liner) Protect(impactForceLevel float32) float32 {
	if impactForceLevel <= 0 {
		return 0.0
	}

	impactForceLevel -= (impactForceLevel * levelOfProtectionPerImpact)
	if l.nextProtector == nil {
		return impactForceLevel
	}

	return l.nextProtector.Protect(impactForceLevel)
}
