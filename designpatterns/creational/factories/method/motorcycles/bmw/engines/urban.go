package engines

// Urban defines an urban engine.
type Urban struct {
	Power     float32
	Type      EngineType
	Cylinders byte
	Torque    float32
}

// UrbanCreator defines an urban engine creator
type UrbanCreator struct{}

func NewUrbanCreator() *UrbanCreator {
	newCreator := UrbanCreator{}
	return &newCreator
}

func (u *UrbanCreator) Create(data BuildData) EngineBehavior {
	engine := Urban{
		Power:     data.Power,
		Type:      data.Type,
		Cylinders: data.Cylinders,
		Torque:    data.Torque,
	}
	return &engine
}

func (u *Urban) IncreasePower() error {
	u.Power += 5.0
	return nil
}

func (u *Urban) ReducePower() error {
	u.Power -= 5.0
	return nil
}
