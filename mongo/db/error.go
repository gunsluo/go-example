package db

type errNoDocuments struct {
	error
}

type errDuplicateKey struct {
	error
}

// NoDocuments return true
func (e *errNoDocuments) NoDocuments() bool {
	return true
}

// DuplicateKey return true
func (e *errDuplicateKey) DuplicateKey() bool {
	return true
}
