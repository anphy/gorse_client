package gorse

import "time"

// Item stores meta data about item.
type Item struct {
	ItemId    string
	Timestamp time.Time
	Labels    []string
	Comment   string
}

// User stores meta data about user.
type User struct {
	UserId    string
	Labels    []string
	Subscribe []string
	Comment   string
}

// Feedback stores feedback.
type Feedback struct {
	FeedbackType string
	UserId       string
	ItemId       string
	Timestamp    time.Time
	Comment      string
}
