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

func TestReversstring_Rune(t *testing.T) {
	want := "dcba  "
	if got := Reversstring_Rune("  abcd"); got != want {
		t.Errorf("Reversstring_backup() = %q, want %q", got, want)
	}
}

func TestBuiltInReverse(t *testing.T) {
	want := "dcba  "
	if got := ReverseBuiltIn("  abcd"); got != want {
		t.Errorf("Reversstring_backup() = %q, want %q", got, want)
	}
}
