package service

import (
	"fmt"
	"time"

	"github.com/wallacetm/go-dev-skills/model/feline"
)

type roarService struct {
}

func NewRoarService() RoarService {
	return &roarService{}
}

func (rs *roarService) DoRoar(feline feline.Feline) {
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%s emits a %s\n", feline.Name(), feline.Roar())
}
