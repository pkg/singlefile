package singlefile

import "testing"

func TestLock(t *testing.T) {
	const key = "somekey"
	unlock, err := Lock(key)
	if err != nil {
		t.Fatalf("could not lock key %q: %v", key, err)
	}
	if err := unlock(); err != nil {
		t.Fatalf("could not release lock: %v", err)
	}
}
