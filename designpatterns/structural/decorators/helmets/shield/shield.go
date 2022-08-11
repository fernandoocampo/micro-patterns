package shields

// Protector defines behavior to protect the rider.
type Protector interface {
	// Protect signals that the protector must protect and
	// return the new impact force level.
	Protect(impactForceLevel float32) float32
}

// Shield contains boot data.
type Shield struct {
	nextProtector Protector
}

const levelOfProtectionPerImpact = 0.3

func New(protector Protector) *Shield {
	newBoot := Shield{
		nextProtector: protector,
	}

	return &newBoot
}

func (s *Shield) Protect(impactForceLevel float32) float32 {
	if impactForceLevel <= 0 {
		return 0.0
	}

	impactForceLevel -= (impactForceLevel * levelOfProtectionPerImpact)
	if s.nextProtector == nil {
		return impactForceLevel
	}

	return s.nextProtector.Protect(impactForceLevel)
}
