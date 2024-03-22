package vo

type Meow string

const meowSound = "Meow"

func NewMeow() Meow {
	return Meow(meowSound)
}

func (m Meow) String() string {
	return string(m)
}
