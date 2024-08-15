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
