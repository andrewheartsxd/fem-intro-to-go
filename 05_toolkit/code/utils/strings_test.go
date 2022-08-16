// ### Part 2: Use TDD to Write A Function

// 1. In the utils directory navigate to your newly created `strings_test.go` file.
// 2. Copy the following code into that test file, and watch the test fail.
// 3. Implement the necessary code to make this test pass.

// ```go
// // strings_test.go
// package utils

// import "testing"

// func TestMakeExcited(t *testing.T) {
// 	expected := "OMG SO EXCITING!"
// 	actual := MakeExcited("omg so exciting")
// 	if actual != expected {
// 		t.Errorf("Average was incorrect! Expected: %s, Actual: %s", expected, actual)
// 	}
// }
// ```

package utils

import "testing"

func TestMakeExcited(t *testing.T) {
	expected := "OMG SO EXCITING!"
	actual := MakeExcited("omg so exciting")
	if actual != expected {
		t.Errorf("Average was incorrect! Expected: %s, Actual: %s", expected, actual)
	}
}
