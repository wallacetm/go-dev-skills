package task

import (
	"github.com/wallacetm/go-dev-skills/application/service"
	"github.com/wallacetm/go-dev-skills/infrastructure/persistence"
)

type Task struct {
	roarService   service.RoarService
	catService    service.CatService
	felineService service.FelineService
}

func NewTask() *Task {
	// Creating dependecies and injections
	catRepo := persistence.NewMemoryCatRepository()
	felineRepo := persistence.NewMemoryFelineRepository()
	catService := service.NewCatService(catRepo)
	roarService := service.NewRoarService()
	felineService := service.NewFelineService(felineRepo, roarService, catService)
	return &Task{roarService, catService, felineService}
}
