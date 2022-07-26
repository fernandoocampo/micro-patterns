package urbans

type Scooter struct {
	EngineOn, Moving bool
	Speed            float32
}

func New() *Scooter {
	return new(Scooter)
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
