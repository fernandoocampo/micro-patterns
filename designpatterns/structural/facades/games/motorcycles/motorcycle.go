package motorcycles

import "log"

const (
	defaultID     = 1
	defaultHP     = 12.0
	defaultTorque = 10.0
)

type Scooter struct {
	ID         int
	HP, Torque float32
}

func GenerateScooter() *Scooter {
	newScooter := Scooter{
		ID:     defaultID,
		HP:     defaultHP,
		Torque: defaultTorque,
	}

	return &newScooter
}

func (s *Scooter) SpeedUp() {
	log.Println("speeding up")
}

func (s *Scooter) Brake() {
	log.Println("braking")
}

func (s *Scooter) Lean() {
	log.Println("lean down")
}
