package sports

type Sport struct {
	EngineOn, Moving bool
	Speed, Power     float32
	DrivingMode      string
}

// New creates sport  motorcycles.
func New() *Sport {
	return new(Sport)
}

func (s *Sport) StartEngine() error {
	s.EngineOn = true

	return nil
}

func (s *Sport) StopEngine() error {
	s.EngineOn = false

	return nil
}

func (s *Sport) SpeedUp(increase float32) error {
	s.Moving = true
	s.Speed += increase

	return nil
}

func (s *Sport) Stop() error {
	s.Moving = false

	return nil
}

func (s *Sport) ActivateDrivingMode(mode string) error {
	s.DrivingMode = mode

	return nil
}
func (s *Sport) IncreasePower(power float32) error {
	s.Power = power

	return nil
}
