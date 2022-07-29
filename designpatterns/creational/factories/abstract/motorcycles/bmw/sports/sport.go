package sports

type Sport struct {
	EngineOn    bool
	Speed       float32
	Moving      bool
	DrivingMode string
	Power       float32
}

// New creates sport  motorcycles.
func New() *Sport {
	newSport := Sport{}
	return &newSport
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
