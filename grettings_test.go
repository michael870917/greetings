package grettings

import (
	"regexp"
	"testing"
)

func TestHelloFromGrettings(t *testing.T) {
	name := "Michael"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := HelloFromGrettings("Michael")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Michael") = %q, %v , need match with %#q, nil`, msg, err, want)
	}
}

func TestHelloFromGrettingsEmpty(t *testing.T) {
	msg, err := HelloFromGrettings("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("Michael") = %q, %v want "", error`, msg, err)
	}
}
