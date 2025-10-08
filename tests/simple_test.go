package tests

import "testing"

func TestSimple(t *testing.T) {
	// Test toujours vrai pour démarrer
	if 1+1 != 2 {
		t.Error("Math is broken!")
	}
}

func TestCompilation(t *testing.T) {
	// Test que le package peut être compilé
	t.Log("✅ Tests package compiles successfully")
}
