package urbans

type Scooter struct {
	EngineOn bool
	Speed    float32
	Moving   bool
}

func New() *Scooter {
	newScooter := Scooter{}
	return &newScooter
}

func (s *Scooter) StartEngine() error {
	s.EngineOn = true
	return nil
}

func (s *Scooter) StopEngine() error {
	s.EngineOn = false
	return nil
}

func (s *Scooter) SpeedUp(increase float32) error {
	s.Moving = true
	s.Speed += increase
	return nil
}

func (s *Scooter) Stop() error {
	s.Moving = false
	return nil
}
