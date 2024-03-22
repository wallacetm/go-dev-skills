package persistence

import (
	"github.com/wallacetm/go-dev-skills/model/feline"
	"github.com/wallacetm/go-dev-skills/vo"
)

type memoryCatRepository struct {
	cats map[string]feline.Cat
}

func NewMemoryCatRepository() feline.CatRepository {
	return &memoryCatRepository{cats: make(map[string]feline.Cat)}
}

func (m *memoryCatRepository) loadCats() {
	m.cats["Stuart"] = feline.NewCat(vo.NewName("Stuart"))
	m.cats["Loro"] = feline.NewCat(vo.NewName("Loro"))
	m.cats["Arthur"] = feline.NewCat(vo.NewName("Arthur"))
	m.cats["Missi"] = feline.NewCat(vo.NewName("Missi"))
	m.cats["Ciri"] = feline.NewCat(vo.NewName("Ciri"))
	m.cats["Whiskers"] = feline.NewCat(vo.NewName("Whiskers"))
	m.cats["Mittens"] = feline.NewCat(vo.NewName("Mittens"))
	m.cats["Shadow"] = feline.NewCat(vo.NewName("Shadow"))
	m.cats["Garfield"] = feline.NewCat(vo.NewName("Garfield"))
	m.cats["Simba"] = feline.NewCat(vo.NewName("Simba"))
	m.cats["Nala"] = feline.NewCat(vo.NewName("Nala"))
	m.cats["Milo"] = feline.NewCat(vo.NewName("Milo"))
	m.cats["Leo"] = feline.NewCat(vo.NewName("Leo"))
	m.cats["Oliver"] = feline.NewCat(vo.NewName("Oliver"))
	m.cats["Max"] = feline.NewCat(vo.NewName("Max"))
	m.cats["Charlie"] = feline.NewCat(vo.NewName("Charlie"))
	m.cats["Bella"] = feline.NewCat(vo.NewName("Bella"))
	m.cats["Lucy"] = feline.NewCat(vo.NewName("Lucy"))
	m.cats["Luna"] = feline.NewCat(vo.NewName("Luna"))
	m.cats["Mia"] = feline.NewCat(vo.NewName("Mia"))
	m.cats["Sophie"] = feline.NewCat(vo.NewName("Sophie"))
}

func (m *memoryCatRepository) GetAll() []feline.Cat {
	if len(m.cats) == 0 {
		m.loadCats()
	}
	cats := make([]feline.Cat, 0, len(m.cats))
	for _, cat := range m.cats {
		cats = append(cats, cat)
	}
	return cats
}
