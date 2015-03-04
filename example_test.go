package singlefile_test

import (
	"log"

	"github.com/pkg/singlefile"
)

func ExampleLock() {
	const key = "xyzzy"
	unlock, err := singlefile.Lock(key)
	if err != nil {
		log.Fatalf("could not aquire singlefile lock: %v", err)
	}
	defer unlock()
}
