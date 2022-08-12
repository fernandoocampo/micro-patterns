package avatars

const (
	defaultSpeed  = 80
	defaultCurves = 67
	defaultBrakes = 20
)

type Avatar struct {
	Speed, Curves, Brakes int
}

func GenerateAvatar() *Avatar {
	defaultAvatar := Avatar{
		Speed:  defaultSpeed,
		Curves: defaultCurves,
		Brakes: defaultBrakes,
	}

	return &defaultAvatar
}
