package lyfeed

import (
	"database/sql"
)

// Context provides some basic application level
// information.
type Context struct {
	DB *sql.DB
}

type Item struct {
	ID       int
	GUID     string
	FeedName string
}

type ItemService interface {
	Item(id int) (*Item, error)
	CreateItem(u *Item) error
}
