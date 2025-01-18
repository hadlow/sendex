package helpers

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreateArgsmap(t *testing.T) {
	args := []string{"id=5", "user=qwerty"}
	expectedOutput := map[string]string{"id": "5", "user": "qwerty"}

	output, err := CreateArgsmap(args)

	if err != nil {
		t.Fatalf("Error creating argsmap")
	}

	if !cmp.Equal(output, expectedOutput) {
		t.Fatalf("Argsmap output not expected")
	}
}
