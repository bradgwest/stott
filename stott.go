package stott

// TODO Blocker and Tasks should embed a more abstract type and interface

// UserID represents a user identifier
type UserID int

// User represents a user of the system
// TODO This should be an authenticated user at some point
type User struct {
	ID       UserID `json:"id"`
	Username string `json:"username"`
}

// StatusID represents a status identifier
type StatusID int

// Status represents a User's Status
type Status struct {
	ID     StatusID `json:"status_id"`
	UserID UserID   `json:"user_id"`
	Text   string   `json:"text"`
}

// LinkID represents a link identifier
type LinkID int

// Link represents a Link for a user
type Link struct {
	ID     LinkID `json:"link_id"`
	UserID UserID `json:"user_id"`
	Text   string `json:"text"`
	URL    string `json:"url"`
}

// BlockerID represents a blocker identifier
type BlockerID int

// Blocker represents a link which blocks
type Blocker struct {
	ID   BlockerID `json:"blocker_id"`
	Link LinkID    `json:"link_id"`
}

// TaskID represents a task identifier
type TaskID int

// Task represents a link which is a Task
type Task struct {
	ID   TaskID `json:"task_id"`
	Link LinkID `json:"link_id"`
}

// UserService represents a service for managing users.
type UserService interface {
	Authenticate(token string) (*User, error)
}

// StatusService represents a service for managing Statuses
type StatusService interface {
	Status(id StatusID) (*Status, error)
	CreateStatus(status *Status) error
	UpdateStatus(is StatusID, text string) error
}

// BlockerService represents a service for managing Blockers
type BlockerService interface {
	Blocker(id BlockerID) (*Blocker, error)
	CreateBlocker(blocker *Blocker) error
	UpdateBlocker(id BlockerID, blocker LinkID) error
}

// TaskService represents a service for managing Tasks
type TaskService interface {
	Task(id TaskID) (*Task, error)
	CreateTask(task *Task) error
	UpdateTask(id TaskID, text LinkID) error
}
