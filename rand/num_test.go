package rand

import (
	"testing"
)

func TestRand(t *testing.T) {
	// int num
	t.Run("random int", func(t *testing.T) {
		t.Logf("random int = %v", RandNum(6, 20))
	})
	// uint num
	t.Run("random uint", func(t *testing.T) {
		var l, u uint = 6, 20
		t.Logf("random uint = %v", RandNum(l, u))
	})
	// negative int num
	t.Run("negative int num", func(t *testing.T) {
		t.Logf("random negative int = %v", RandNum(-100, -1))
	})
	// float num
	t.Run("float num", func(t *testing.T) {
		t.Logf("random float = %v", RandNum(22.5, 500.07))
	})
}
