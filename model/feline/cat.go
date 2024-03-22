package feline

import "github.com/wallacetm/go-dev-skills/vo"

type Cat interface {
	Meow() string
	Name() string
}

type cat struct {
	meow vo.Meow
	name vo.Name
}

func NewCat(name vo.Name) Cat {
	return &cat{meow: vo.NewMeow(), name: name}
}

func (c *cat) Meow() string {
	return c.meow.String()
}

func (c *cat) Name() string {
	return c.name.String()
}

type CatRepository interface {
	GetAll() []Cat
}
