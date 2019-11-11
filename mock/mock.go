package mock

import (
	"github.com/bradgwest/stott"
)

// Authenticator represents a service for authenticating users.
type Authenticator struct {
	AuthenticateFn      func(token string) (*stott.User, error)
	AuthenticateInvoked bool
}

func (a *Authenticator) Authenticate(token string) (*stott.User, error) {
	a.AuthenticateInvoked = true
	return a.AuthenticateFn(token)
}

// DefaultUser is the user authenticated by DefaultAuthenticator.
var DefaultUser = &stott.User{ID: 100}

// DefaultAuthenticator returns an authenticator that returns the default user
func DefaultAuthenticator() Authenticator {
	return Authenticator{
		AuthenticateFn: func(token string) (*stott.User, error) { return DefaultUser, nil },
	}
}

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
