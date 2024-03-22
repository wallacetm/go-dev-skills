package persistence

import (
	"github.com/wallacetm/go-dev-skills/model/feline"
	"github.com/wallacetm/go-dev-skills/vo"
)

type memoryFelineRepository struct {
	felines map[string]feline.Feline
}

func NewMemoryFelineRepository() feline.FelineRepository {
	return &memoryFelineRepository{felines: make(map[string]feline.Feline)}
}

func (m *memoryFelineRepository) loadFelines() {
	m.felines["Tiger"] = feline.NewFeline(vo.NewName("Tiger"))
	m.felines["Lion"] = feline.NewFeline(vo.NewName("Lion"))
	m.felines["Lioness"] = feline.NewFeline(vo.NewName("Lioness"))
	m.felines["Leopard"] = feline.NewFeline(vo.NewName("Leopard"))
	m.felines["Puma"] = feline.NewFeline(vo.NewName("Puma"))
	m.felines["Cheetah"] = feline.NewFeline(vo.NewName("Cheetah"))
	m.felines["Jaguar"] = feline.NewFeline(vo.NewName("Jaguar"))
	m.felines["Panther"] = feline.NewFeline(vo.NewName("Panther"))
	m.felines["Cougar"] = feline.NewFeline(vo.NewName("Cougar"))
	m.felines["Bobcat"] = feline.NewFeline(vo.NewName("Bobcat"))
	m.felines["Serval"] = feline.NewFeline(vo.NewName("Serval"))
	m.felines["Caracal"] = feline.NewFeline(vo.NewName("Caracal"))
	m.felines["Ocelot"] = feline.NewFeline(vo.NewName("Ocelot"))
	m.felines["Margay"] = feline.NewFeline(vo.NewName("Margay"))
}

func (m *memoryFelineRepository) GetAll() []feline.Feline {
	if len(m.felines) == 0 {
		m.loadFelines()
	}
	felines := make([]feline.Feline, 0, len(m.felines))
	for _, feline := range m.felines {
		felines = append(felines, feline)
	}
	return felines
}
