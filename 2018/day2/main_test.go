package main

import (
	"log"
	"testing"
)

func TestExactlyOneMatch(t *testing.T) {
	if match, isMatch := MatchExactlyOne("abcde", "avcde"); !isMatch || match != "acde" {
		log.Println("Out", match, isMatch)
		t.Fail()
	}
}

func TestNotOneMatch(t *testing.T) {
	if match, isMatch := MatchExactlyOne("zzzzz", "avcde"); isMatch || match != "" {
		log.Println("Out", match, isMatch)
		t.Fail()
	}
}
