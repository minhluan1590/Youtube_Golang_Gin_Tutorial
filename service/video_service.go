package service

import "github.com/minhluan1590/Youtube_Golang_Gin_Tutorial/entity"

// VideoService defines the interface for video-related operations.
type VideoService interface {
	Save(video *entity.Video) error // Save a new video to the collection.
	FindAll() ([]entity.Video, error) // Retrieve all videos from the collection.
}

// videoService is a concrete implementation of the VideoService interface.
type videoService struct {
	videos []entity.Video // Collection of videos.
}

// NewVideoService creates a new instance of VideoService.
func NewVideoService() VideoService {
	return &videoService{
		videos: []entity.Video{}, // Initialize the videos slice as empty.
	}
}

// Save adds a new video to the collection.
func (s *videoService) Save(video *entity.Video) error {
	s.videos = append(s.videos, *video)
	return nil
}

// FindAll retrieves all videos from the collection.
func (s *videoService) FindAll() ([]entity.Video, error) {
	return s.videos, nil
}