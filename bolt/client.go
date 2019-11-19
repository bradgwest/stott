package bolt

import (
	"time"

	"github.com/boltdb/bolt"
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

// NewClient creates a new boltDB client
func NewClient() *Client {
	return &Client{
		Now: time.Now,
	}
}

// Open opens and initializes the BoltDB database.
func (c *Client) Open() error {
	// Open database file.
	db, err := bolt.Open(c.Path, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	c.db = db

	// Start writable transaction.
	tx, err := c.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Initialize top-level buckets.
	if _, err := tx.CreateBucketIfNotExists([]byte("Statuses")); err != nil {
		return err
	}

	// Save transaction to disk.
	return tx.Commit()
}

// Close closes then underlying BoltDB database.
func (c *Client) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}

// Connect returns a new session to the BoltDB database.
func (c *Client) Connect() *Session {
	s := newSession(c.db)
	s.authenticator = c.Authenticator
	s.now = c.Now()
	return s
}
