package service

import (
	"sync"

	"github.com/wallacetm/go-dev-skills/model/feline"
)

type felineService struct {
	repository  feline.FelineRepository
	roarservice RoarService
	catService  CatService
}

func NewFelineService(repository feline.FelineRepository, roarService RoarService, catService CatService) FelineService {
	return &felineService{repository, roarService, catService}
}

func (fs *felineService) getAllFelines() []feline.Feline {
	felines := fs.repository.GetAll()
	catFelines := fs.catService.GetAllCatsAsFelines()
	felines = append(felines, catFelines...)
	return felines
}

func (fs *felineService) RoarAll() {
	felines := fs.getAllFelines()
	// Roaring all felines
	for _, f := range felines {
		fs.roarservice.DoRoar(f)
	}
}

// RoarAllAsync is a function that receives an integer that represents the number of async goroutines to roar the felines
func (fs *felineService) RoarAllAsync(asyncLen int) int {
	felines := fs.getAllFelines()
	// Creating a channel to send the felines
	felineChan := make(chan feline.Feline, len(felines))
	// Sending all felines to the channel
	go func() {
		for _, f := range felines {
			felineChan <- f
		}
		// Closing the channel to avoid deadlock
		close(felineChan)
	}()
	// Creating a slice to store the felines in the order they were roared
	roaredFelines := make([]string, 0)
	// Using a mutex to avoid race conditions when storing the feline name in the slice
	roaredFelinesMutex := sync.Mutex{}
	// Creating a wait group to wait for all async goroutines to finish
	wg := sync.WaitGroup{}
	wg.Add(asyncLen)
	// Creating async goroutines to roar the felines
	for i := 0; i < asyncLen; i++ {
		go func() {
			// Receiving the felines from the channel and roaring them
			for f := range felineChan {
				fs.roarservice.DoRoar(f)
				// Storing the feline name in the slice
				roaredFelinesMutex.Lock()
				roaredFelines = append(roaredFelines, f.Name())
				roaredFelinesMutex.Unlock()
			}
			// Decreasing the wait group counter
			wg.Done()
		}()
	}
	// Waiting for all async goroutines to finish
	wg.Wait()
	return len(roaredFelines)
}
