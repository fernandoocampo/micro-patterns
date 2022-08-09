package motorcycles

// Motorcycle contains data related to a motorcycle.
type Motorcycle struct {
	Brand string
	Name  string
	Model int
	Miles int
}

func (m *Motorcycle) Clone() Motorcycle {
	return Motorcycle{
		Brand: m.Brand,
		Name:  m.Name,
		Model: m.Model,
		Miles: m.Miles,
	}
}
