package adventures

import "fmt"

type Adventure struct {
	EngineOn      bool
	Speed         float32
	Moving        bool
	DrivingMode   string
	Power         float32
	CruiseControl bool
}

func New() *Adventure {
	newAdventure := Adventure{}
	return &newAdventure
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
	fmt.Println("adjusting suspensions")
	return nil
}
