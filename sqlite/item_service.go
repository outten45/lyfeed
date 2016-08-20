package sqlite

import (
	"database/sql"

	// blank import for sqlite3 adapter
	_ "github.com/mattn/go-sqlite3"
	"github.com/outten45/lyfeed"
)

// ItemService represents a Sqlite implementation of lyfeed.ItemService
type ItemService struct {
	DB *sql.DB
}

// Item returns the item at id
func (is *ItemService) Item(id int) (*lyfeed.Item, error) {
	var i lyfeed.Item

	return &i, nil
}

// SaveItem
func (is *ItemService) SaveItem(item lyfeed.Item) error {

	return nil
}
