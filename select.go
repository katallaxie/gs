package gs

import "context"

// Querier ...
type Querier[R any] interface {
	Query(ctx context.Context, query string, args ...any) ([]R, error)
}

// Query ...
func (db *DB[R]) Query(ctx context.Context, query string, args ...any) ([]R, error) {
	results := make([]R, 0)

	return results, nil
}
