package helmets

// Protector defines behavior to protect the rider.
type Protector interface {
	// Protect signals that the protector must protect and
	// return the new impact force level.
	Protect(impactForceLevel float32) float32
}

// Helment contains helmet data.
type Helment struct {
	nextProtector Protector
}

func New() *Helment {
	return new(Helment)
}

func (h *Helment) SetProtection(protector Protector) {
	h.nextProtector = protector
}

func (h *Helment) Protect(impactForceLevel float32) float32 {
	return h.nextProtector.Protect(impactForceLevel)
}
