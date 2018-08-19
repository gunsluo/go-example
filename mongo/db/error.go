package db

import "github.com/pkg/errors"

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

// IsErrNoDocuments return true
func IsErrNoDocuments(err error) bool {
	type errNoDocuments interface {
		NoDocuments() bool
	}

	err = errors.Cause(err)
	if e, ok := err.(errNoDocuments); ok {
		return e.NoDocuments()
	}

	return false
}

// IsErrDuplicateKey return true
func IsErrDuplicateKey(err error) bool {
	type errDuplicateKey interface {
		DuplicateKey() bool
	}

	err = errors.Cause(err)
	if e, ok := err.(errDuplicateKey); ok {
		return e.DuplicateKey()
	}

	return false
}
