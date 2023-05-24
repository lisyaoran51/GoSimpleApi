package service

import (
	"regexp"
	"testing"
)

func Test_userService_Save(t *testing.T) {

	match, err := regexp.MatchString("[0-9A-Za-z]", "pppppppp")
	if err != nil {
		t.Logf("test: err: %v\n", err)
		return
	}
	if !match {
		t.Logf("test: no match")
		return
	}
	t.Logf("test: match")
}
