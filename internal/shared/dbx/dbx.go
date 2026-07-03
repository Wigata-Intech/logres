// Package dbx holds helpers shared by every repository.
package dbx

import "errors"

// ErrNotFound is the sentinel repositories return for a missing or soft-deleted
// row. Compare with errors.Is.
var ErrNotFound = errors.New("record not found")
