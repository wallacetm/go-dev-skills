package vo

type Roar string

const roarSound = "Roar"

func NewRoar() Roar {
	return Roar(roarSound)
}

func (r Roar) String() string {
	return string(r)
}
