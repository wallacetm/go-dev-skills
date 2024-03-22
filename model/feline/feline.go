package feline

import "github.com/wallacetm/go-dev-skills/vo"

type Feline interface {
	Roar() string
	Name() string
}

type feline struct {
	roar vo.Roar
	name vo.Name
}

func NewFeline(name vo.Name) Feline {
	return &feline{roar: vo.NewRoar(), name: name}
}

func (f *feline) Roar() string {
	return f.roar.String()
}

func (f *feline) Name() string {
	return f.name.String()
}

type FelineRepository interface {
	GetAll() []Feline
}
