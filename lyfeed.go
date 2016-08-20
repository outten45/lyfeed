package lyfeed

import (
	"database/sql"
	"time"
)

// Context provides some basic application level
// information.
type Context struct {
	DB *sql.DB
}

type Item struct {
	ID              int
	GUID            string
	Feed            string
	Channel         string
	PubDate         time.Time
	Raw             string
	expectedVersion int
	changes         []interface{}
}

type ItemService interface {
	Item(id int) (*Item, error)
	SaveItem(u *Item) error
}

type User struct {
	ID   int
	Name string
}
