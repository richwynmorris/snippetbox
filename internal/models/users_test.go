package models

import (
	"testing"

	"richwynmorris.co.uk/internal/assert"
)

func TestUserModelExists(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration tests")
	}

	tests := []struct {
		name   string
		userID int
		want   bool
	}{
		{
			name:   "ValidId",
			userID: 1,
			want:   true,
		},
		{
			name:   "ZeroID",
			userID: 0,
			want:   false,
		},
		{
			name:   "non-existent ID",
			userID: 2,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDB(t)

			m := UserModel{db}

			exists, err := m.Exists(tt.userID)

			assert.Equal(t, exists, tt.want)
			assert.NilError(t, err)
		})
	}
}
