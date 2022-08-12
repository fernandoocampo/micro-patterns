package games

type Position struct {
	X, Y float32
}

type Skill struct {
	Brake, Curve, SpeedUp, Soft, Sneaky float32
}

type PlayerStatus struct {
	UserID                                    string
	Skills                                    Skill
	Position                                  Position
	Motorcycle, Jacket, Helmet, Boots, Gloves int
}
