package service

import (
	"github.com/wallacetm/go-dev-skills/model/feline"
)

type catService struct {
	repository feline.CatRepository
}

func NewCatService(repository feline.CatRepository) CatService {
	return &catService{repository}
}

func (cs *catService) GetAllCatsAsFelines() []feline.Feline {
	cats := cs.repository.GetAll()
	felines := make([]feline.Feline, len(cats))
	for i := 0; i < len(cats); i++ {
		felines[i] = feline.NewCatFeline(cats[i])
	}
	return felines
}
