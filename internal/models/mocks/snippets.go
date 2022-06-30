package mocks

import (
	"time"

	"richwynmorris.co.uk/internal/models"
)

var mockSnippet = &models.Snippet{
	ID:      1,
	Title:   "An old silent pond",
	Content: "An old silent pods...",
	Created: time.Now(),
	Expires: time.Now(),
}

type MockSnippetModel struct{}

func (m *MockSnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *MockSnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}
