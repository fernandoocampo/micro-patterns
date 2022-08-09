package motorcycles

type ScooterBuilder struct {
	scooter *Scooter
}

type EngineSpecification struct {
	CylinderCapacity int
	HorsePower       int
}

func NewScooterBuilder() *ScooterBuilder {
	newScooterBuilder := new(ScooterBuilder)
	newScooterBuilder.scooter = new(Scooter)

	return newScooterBuilder
}

func (s *ScooterBuilder) WithSerie(serie string) *ScooterBuilder {
	newScooter := new(Scooter)
	newScooter.Serie = serie
	s.scooter = newScooter

	return s
}

func (s *ScooterBuilder) WithEngine(engineSpecification EngineSpecification) *ScooterBuilder {
	newEngine := Engine{
		CC: engineSpecification.CylinderCapacity,
		HP: engineSpecification.HorsePower,
	}
	s.scooter.Engine = &newEngine

	return s
}

func (s *ScooterBuilder) WithFairing(fairingID string) *ScooterBuilder {
	// for now we are going to imagen there are a fairing catalog
	newFairing := Fairing{
		Version: fairingID,
		Model:   fairingID + "-abc-2022",
	}
	s.scooter.Fairing = &newFairing

	return s
}

func (s *ScooterBuilder) WithCentralComputer(centralComputerID string) *ScooterBuilder {
	newCC := CentralComputer{
		Version: centralComputerID,
		Model:   "abc",
	}
	s.scooter.CentralComputer = &newCC

	return s
}

func (s *ScooterBuilder) Build() *Scooter {
	return s.scooter
}
