package add

import (
	"testing"
)

func TestAddInteger(t *testing.T) {
	actual := AddInteger(1, 1)
	expected := 2

	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func TestAddString(t *testing.T) {
	actual := AddString("X-", "HOGEHOGE")
	expected := "X-HOGEHOGE"

	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
