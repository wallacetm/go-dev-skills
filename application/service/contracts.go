package service

import "github.com/wallacetm/go-dev-skills/model/feline"

type FelineService interface {
	RoarAll()
	RoarAllAsync(asyncLen int) int
}

type CatService interface {
	GetAllCatsAsFelines() []feline.Feline
}

type RoarService interface {
	DoRoar(feline.Feline)
}
