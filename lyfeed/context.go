package lyfeed

import (
	"database/sql"
)

// Context provides some basic application level
// information.
type Context struct {
	DB *sql.DB
}
