// Package singlefile implements a host wide locking mechanism.
//
// Singlefile can be used to ensure that two scheduled processes
// on a host do not run at the same time.
package singlefile

import "errors"

// Lock takes a host wide lock on the specified key.
// Lock returns an unlock function, and an error indicating the
// success or failure of the lock operation.
// key must not be an empty string and implementations may require
// that they key is of sufficient length to ensure uniqueness.
func Lock(key string) (unlock func() error, err error) {
	if key == "" {
		return nil, errors.New("key must not be empty")
	}
	return lock(key)
}
