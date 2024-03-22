package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wallacetm/go-dev-skills/mocks"
	"github.com/wallacetm/go-dev-skills/model/feline"
	"github.com/wallacetm/go-dev-skills/vo"
)

func TestFelineService_RoarAll(t *testing.T) {
	// Create a mock repository
	repoMock := &mocks.FelineRepository{}
	// Create a mock roar service
	roarServiceMock := &mocks.RoarService{}
	// Create a mock cat service
	catServiceMock := &mocks.CatService{}
	// Create a slice of 1000 felines to be returned by the cat service mock
	mockFelines := make([]feline.Feline, 0)
	for i := 0; i < 1000; i++ {
		mockFelines = append(mockFelines, feline.NewFeline(vo.NewName(fmt.Sprintf("Feline %d", i))))
	}
	// Create a slice of 1000 cat felines to be returned by the cat service mock
	mockCatFelines := make([]feline.Feline, 0)
	for i := 0; i < 1000; i++ {
		mockCatFelines = append(mockCatFelines, feline.NewCatFeline(feline.NewCat(vo.NewName(fmt.Sprintf("Cat %d", i)))))
	}

	// Create a feline service instance with the mocks
	fs := NewFelineService(repoMock, roarServiceMock, catServiceMock)

	// Define the expected behavior of the mocks
	repoMock.On("GetAll").Return(mockFelines, nil).Times(1)
	catServiceMock.On("GetAllCatsAsFelines").Return(mockCatFelines, nil).Times(1)
	roarServiceMock.On("DoRoar", mock.Anything).Times(2000)

	// Call the RoarAll method
	fs.RoarAll()

	// Assert that the mocks were called as expected
	repoMock.AssertExpectations(t)
	catServiceMock.AssertExpectations(t)
	roarServiceMock.AssertExpectations(t)
}

func TestFelineService_RoarAllAsync(t *testing.T) {
	// Create a mock repository
	repoMock := &mocks.FelineRepository{}
	// Create a mock roar service
	roarServiceMock := &mocks.RoarService{}
	// Create a mock cat service
	catServiceMock := &mocks.CatService{}

	// Create a slice of 1000 felines to be returned by the cat service mock
	mockFelines := make([]feline.Feline, 0)
	for i := 0; i < 1000; i++ {
		mockFelines = append(mockFelines, feline.NewFeline(vo.NewName(fmt.Sprintf("Feline %d", i))))
	}
	// Create a slice of 1000 cat felines to be returned by the cat service mock
	mockCatFelines := make([]feline.Feline, 0)
	for i := 0; i < 1000; i++ {
		mockCatFelines = append(mockCatFelines, feline.NewCatFeline(feline.NewCat(vo.NewName(fmt.Sprintf("Cat %d", i)))))
	}

	// Create a feline service instance with the mocks
	fs := NewFelineService(repoMock, roarServiceMock, catServiceMock)

	// Define the expected behavior of the mocks
	repoMock.On("GetAll").Return(mockFelines, nil)
	catServiceMock.On("GetAllCatsAsFelines").Return(mockCatFelines, nil)
	roarServiceMock.On("DoRoar", mock.Anything).Times(2000)

	// Call the RoarAllAsync method
	roaredLen := fs.RoarAllAsync(500)

	// Assert that the number of roared felines is correct
	assert.Equal(t, 2000, roaredLen)

	// Assert that the mocks were called as expected
	repoMock.AssertExpectations(t)
	catServiceMock.AssertExpectations(t)
	roarServiceMock.AssertExpectations(t)
}
