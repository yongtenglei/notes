package util

import "testing"

func TestRandomPort(t *testing.T) {
	port := RandomPort("localhost")
	if port == -1 {
		panic("err!")
	}
}
