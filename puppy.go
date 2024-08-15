package puppy

import dog "github.com/KamiKruse/Dog"

func Bark() string {
		return "Woof!"
}
func Barks() string {
		return "Woof!Woof!Woof!"
}

func BigBark() string {
	return dog.Grownup(Bark())
}
func BigBarks() string {
	return dog.Grownup(Barks())
}
func Fromversion110() string {
	return "This is version 1.1.0"
}
func Fromversion120() string {
	return "This is version 1.2.0"
}
