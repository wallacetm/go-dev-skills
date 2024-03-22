package feline

// CatFeline is a adapter for a Cat to be a Feline
type CatFeline struct {
	Cat
}

func NewCatFeline(cat Cat) Feline {
	return &CatFeline{Cat: cat}
}

// Roar returns the meow of the cat
func (cf *CatFeline) Roar() string {
	return cf.Meow()
}
