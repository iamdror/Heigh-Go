package puppy

import "github.com/GoesToEleven/dog"

func Barc() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Barc())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}
