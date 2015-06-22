package lyfeed

import (
	"database/sql"

	"github.com/gchaincl/dotsql"
)

// Context provides some basic application level
// information.
type Context struct {
	DB     *sql.DB
	DotSql *dotsql.DotSql
}
