package worlds

const defaultSeed = 10000000

type World struct {
	Seed int
}

func Generate() *World {
	newWorld := World{
		Seed: defaultSeed,
	}

	return &newWorld
}
