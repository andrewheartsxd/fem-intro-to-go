package utils

import (
	"fem-intro-to-go/05_toolkit/code/exercise_5a_solution/utils"
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 12
	actual := utils.Add(5, 7)

	if actual != expected {
		t.Errorf("Sum was incorrect! Expected: %d, Actual: %d", expected, actual)
	}
}
