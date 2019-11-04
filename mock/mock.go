package mock

import (
	"github.com/bradgwest/stott"
)

// StatusService mocks a stott.StatusService
type StatusService struct {
	StatusFn      func(id stott.StatusID) (*stott.Status, error)
	StatusInvoked bool

	CreateStatusFn      func(status *stott.Status) error
	CreateStatusInvoked bool

	UpdateStatusFn      func(id stott.StatusID, text string) error
	UpdateStatusInvoked bool
}

// Status mocks st0tt.Status
func (s *StatusService) Status(id stott.StatusID) (*stott.Status, error) {
	s.StatusInvoked = true
	return s.StatusFn(id)
}

// CreateStatus mocks stott.CreateStatus
func (s *StatusService) CreateStatus(status *stott.Status) error {
	s.CreateStatusInvoked = true
	return s.CreateStatusFn(status)
}

// UpdateStatus mocks stott stott.UpdateStatus
func (s *StatusService) UpdateStatus(id stott.StatusID, text string) error {
	s.UpdateStatusInvoked = true
	return s.UpdateStatusFn(id, text)
}

// UserService represents a service for managing user authentication.
type UserService struct {
	AuthenticateFn      func(token string) (*stott.User, error)
	AuthenticateInvoked bool
}

// Authenticate mocks UserService.Authenticate
func (s *UserService) Authenticate(token string) (*stott.User, error) {
	s.AuthenticateInvoked = true
	return s.AuthenticateFn(token)
}
