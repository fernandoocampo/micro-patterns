package helmets

// Protector defines behavior to protect the rider.
type Protector interface {
	// Protect signals that the protector must protect and
	// return the new impact force level.
	Protect(impactForceLevel float32) float32
}

// Helmet contains helmet data.
type Helmet struct {
	nextProtector Protector
}

func New() *Helmet {
	return new(Helmet)
}

func (h *Helmet) SetProtection(protector Protector) {
	h.nextProtector = protector
}

func (h *Helmet) Protect(impactForceLevel float32) float32 {
	return h.nextProtector.Protect(impactForceLevel)
}
