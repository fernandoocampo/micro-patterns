package adventures

import "log"

type Adventure struct {
	EngineOn, Moving, CruiseControl bool
	Speed, Power                    float32
	DrivingMode                     string
}

func New() *Adventure {
	return new(Adventure)
}

func (a *Adventure) StartEngine() error {
	a.EngineOn = true

	return nil
}

func (a *Adventure) StopEngine() error {
	a.EngineOn = false

	return nil
}

func (a *Adventure) SpeedUp(increase float32) error {
	a.Moving = true
	a.Speed += increase

	return nil
}

func (a *Adventure) Stop() error {
	a.Moving = false

	return nil
}

func (a *Adventure) ActivateDrivingMode(mode string) error {
	a.DrivingMode = mode

	return nil
}

func (a *Adventure) IncreasePower(power float32) error {
	a.Power = power

	return nil
}

func (a *Adventure) ActiveCruiseControl() error {
	a.CruiseControl = true

	return nil
}

func (a *Adventure) AdjustSuspensions() error {
	log.Println("adjusting suspensions")

	return nil
}
