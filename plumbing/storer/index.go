package storer

import "github.com/Accel-Byte/go-git/v6/plumbing/format/index"

// IndexStorer generic storage of index.Index
type IndexStorer interface {
	SetIndex(*index.Index) error
	Index() (*index.Index, error)
}
