package motorcycles

// Brand defines the motorcycle brands.
type Brand int32

// UrbanBehavior defines a characteristic behavior of urban motorcycles.
type UrbanBehavior interface {
	StartEngine() error
	SpeedUp(increase float32) error
	Stop() error
}

// SportBehavior defines a characteristic behavior of sport motorcycles.
type SportBehavior interface {
	StartEngine() error
	SpeedUp(increase float32) error
	Stop() error
	ActivateDrivingMode(mode string) error
	IncreasePower(power float32) error
}

// AdventureBehavior defines a characteristic behavior of adventure motorcycles.
type AdventureBehavior interface {
	StartEngine() error
	SpeedUp(increase float32) error
	Stop() error
	ActiveCruiseControl() error
	AdjustSuspensions() error
	IncreasePower(power float32) error
	ActivateDrivingMode(mode string) error
}
