package bolt

import (
	"time"

	"github.com/bradgwest/stott"
)

// Client represents a client to the underlying BoltDB store
type Client struct {
	// Filename to the BoltDB store
	Path string

	// Authenticator to use
	Authenticator stott.Authenticator

	Now func() time.Time

	db *bolt.DB
}

func NewClient() *Client {
	return &Client{
		Now: time.Now,
	}
}
