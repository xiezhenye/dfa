package dfa

import (
	"bytes"
	"testing"
)

func TestM_Match(t *testing.T) {
	m := Char("123").AtLeast(1).Minimize()
	n, _, ok := m.Match([]byte("321321"))
	if !ok || n != 6 {
		println(n, ok)
		t.Error("match fail")
	}
}

func TestFastM_Match(t *testing.T) {
	m := Char("123").AtLeast(1).ToFast()
	n, _, ok := m.Match([]byte("321321"))
	if !ok || n != 6 {
		println(n, ok)
		t.Error("match fail")
	}
}

func TestM_MatchReader(t *testing.T) {
	m := Char("123").AtLeast(1).Minimize()
	out := make([]byte, 0, 10)
	s := "321321"
	n, _, ok, err := m.MatchReader(bytes.NewReader([]byte(s)), &out)
	if err != nil || !ok || n != 6 || string(out) != s {
		t.Error("match fail")
	}
}

func TestFastM_MatchReader(t *testing.T) {
	m := Char("123").AtLeast(1).ToFast()
	out := make([]byte, 0, 10)
	s := "321321"
	n, _, ok, err := m.MatchReader(bytes.NewReader([]byte(s)), &out)
	if err != nil || !ok || n != 6 || string(out) != s {
		println(err, ok, n)
		t.Error("match fail")
	}
}