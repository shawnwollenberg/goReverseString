package Reversestring

import "testing"

func TestReversestring(t *testing.T) {
	want := "nwahs ih"
	if got := Reversestring("hi shawn"); got != want {
		t.Errorf("Reversestring() = %q, want %q", got, want)
	}
}

func TestReversestring2(t *testing.T) {
	want := "dcba  "
	if got := Reversestring("  abcd"); got != want {
		t.Errorf("Reversestring() = %q, want %q", got, want)
	}
}

func TestReversstring_backup(t *testing.T) {
	want := "dcba  "
	if got := Reversstring_backup("  abcd"); got != want {
		t.Errorf("Reversstring_backup() = %q, want %q", got, want)
	}
}
