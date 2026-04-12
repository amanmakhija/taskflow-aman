package project

import (
	"time"

	"github.com/google/uuid"
)

type Service struct{}

func (s *Service) Create(name, description, userID string) (*Project, error) {
	project := &Project{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
		OwnerID:     userID,
		CreatedAt:   time.Now(),
	}

	err := CreateProject(project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (s *Service) List(userID string) ([]Project, error) {
	return GetProjectsByUser(userID)
}
