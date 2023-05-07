package gs

import "database/sql"

// DB ...
type DB[R any] struct {
	*sql.DB
	driver string

	Querier[R]
}
