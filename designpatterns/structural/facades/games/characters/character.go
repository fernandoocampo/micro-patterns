package characters

const defaultNumberOfCharacters = 3

type Character struct {
	ID int
}

func GenerateCharacters() []*Character {
	var result []*Character

	for i := 0; i < defaultNumberOfCharacters; i++ {
		newCharacter := new(Character)
		newCharacter.ID = i + 1
		result = append(result, newCharacter)
	}

	return result
}
