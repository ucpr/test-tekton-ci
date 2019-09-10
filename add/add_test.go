package add

import (
	"testing"
)

func TestAdd(t *testing.T) {
	actual := Add(1, 1)
	expected := 2

	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
