package motorcycles

type scooterBuilder struct {
	scooter *Scooter
}

type EngineSpecification struct {
	CylinderCapacity int
	HorsePower       int
}

func NewScooterBuilder() *scooterBuilder {
	newScooterBuilder := scooterBuilder{}
	return &newScooterBuilder
}

func (s *scooterBuilder) New(serie string) *scooterBuilder {
	newScooter := Scooter{
		Serie: serie,
	}
	s.scooter = &newScooter
	return s
}
func (s *scooterBuilder) WithEngine(engineSpecification EngineSpecification) *scooterBuilder {
	newEngine := Engine{
		CC: engineSpecification.CylinderCapacity,
		HP: engineSpecification.HorsePower,
	}
	s.scooter.Engine = &newEngine
	return s
}
func (s *scooterBuilder) WithFairing(fairingID string) *scooterBuilder {
	// TODO for now we are going to imagen there are a fairing catalog
	newFairing := Fairing{
		Version: fairingID,
		Model:   fairingID + "-abc-2022",
	}
	s.scooter.Fairing = &newFairing
	return s
}
func (s *scooterBuilder) WithCentralComputer(centralComputerID string) *scooterBuilder {
	newCC := CentralComputer{
		Version: centralComputerID,
		Model:   "abc",
	}
	s.scooter.CentralComputer = &newCC
	return s
}
func (s *scooterBuilder) Build() *Scooter {
	return s.scooter
}
