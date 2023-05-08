package myapp

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("hello is not a valid email")
	}

	_, err = IsEmail("hello@gmail.com")
	if err != nil {
		t.Error("hello@gmail.com is a valid email")
	}

	_, err = IsEmail("hello@xyz")
	if err == nil {
		t.Error("hello@gmail is not a valid email")
	}
}
