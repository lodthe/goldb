package db

// Triplet represents a triplet from the database.
type Triplet struct {
	Key     string
	Value   string
	Version Version
}

// Version represents a unique sequence number of a key-value pair.
//
// The only way to get Version is to fetch data.
// Inner implementation is hidden from user to provide compatibility
// of the external interface.
type Version struct {
	lseq string
}

func (v Version) String() string {
	return v.lseq
}

func newVersion(lseq string) Version {
	return Version{
		lseq: lseq,
	}
}
