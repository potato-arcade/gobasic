package object

import (
	"math"
	"testing"
)

// Test we can create error/int/string and that they have the correct types
func TestTypes(t *testing.T) {

	v := StringObject{Value: "Steve"}
	if v.Type() != STRING {
		t.Errorf("Wrong type for String")
	}

	n := NumberObject{Value: math.Pi}
	if n.Type() != NUMBER {
		t.Errorf("Wrong type for Number")
	}

	e := ErrorObject{Value: "You fail!"}
	if e.Type() != ERROR {
		t.Errorf("Wrong type for Error")
	}
}

func TestError(t *testing.T) {

	a := Error("Test")
	b := Error("Test %d", 3)
	c := Error("Test %s", "me")

	// Test types
	if a.Type() != ERROR {
		t.Errorf("Object has the wrong type!")
	}
	if b.Type() != ERROR {
		t.Errorf("Object has the wrong type!")
	}
	if c.Type() != ERROR {
		t.Errorf("Object has the wrong type!")
	}

	// Test values
	if a.Value != "Test" {
		t.Errorf("Wrong value for error-message")
	}
	if b.Value != "Test 3" {
		t.Errorf("Wrong value for error-message")
	}
	if c.Value != "Test me" {
		t.Errorf("Wrong value for error-message")
	}
}
