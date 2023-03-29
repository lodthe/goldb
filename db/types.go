package db

import "github.com/lodthe/goldb/db/dbclient"

// Triplet represents a triplet from the database.
type Triplet struct {
	Key     string
	Value   string
	Version Version
}

func tripletFromInternal(t dbclient.Triplet) Triplet {
	return Triplet{
		Key:     t.Key,
		Value:   t.Value,
		Version: NewVersion(t.Lseq),
	}
}

// Version represents a unique sequence number of a key-value pair.
//
// The only recommended way to get Version is to fetch data.
// Inner implementation is hidden from user to provide compatibility
// of the external interface.
type Version struct {
	lseq string
}

func (v Version) String() string {
	return v.lseq
}

func NewVersion(lseq string) Version {
	return Version{
		lseq: lseq,
	}
}

var zeroVersion Version = NewVersion("#000000000000000000000000")
