package engines

// EngineType define engine types.
type EngineType int32

// BuildData contains data to build a new engine.
type BuildData struct {
	Power     float32
	Type      EngineType
	Cylinders byte
	Torque    float32
}

// EngineBehavior defines engine behavior.
type EngineBehavior interface {
	IncreasePower() error
	ReducePower() error
}

// Creator defines logic to create engines based on category.
type Creator interface {
	// Create create a new engine
	Create(params BuildData) EngineBehavior
}

const (
	TwoStrokes EngineType = iota
	FourStrokes
	Electric
)

// EngineFactory defines an engine factory.
type EngineFactory struct {
	line Creator
}

// NewFactory creates a new factory.
func NewFactory(line Creator) *EngineFactory {
	newFactory := EngineFactory{
		line: line,
	}

	return &newFactory
}

func (e *EngineFactory) BuildEngine(data BuildData) EngineBehavior {
	// Do some stuff related to build engines
	specificEngine := e.line.Create(data)
	// Do more stuff related to build engines
	return specificEngine
}
