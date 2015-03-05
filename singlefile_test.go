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

func TestLockTwice(t *testing.T) {
	const key = "somekey"
	unlock, err := Lock(key)
	if err != nil {
		t.Fatalf("could not lock key %q: %v", key, err)
	}
	u2, err := Lock(key)
	if err == nil {
		u2()
		t.Fatalf("second Lock should have failed")
	}
	if err := unlock(); err != nil {
		t.Fatalf("could not release lock: %v", err)
	}
}
