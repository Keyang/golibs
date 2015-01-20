package kxenv

import "testing"

func TestEnvcanSetAndGet(t *testing.T) {
	Set("hello", "world")
	val := Get("hello", "unknown")
	if val != "world" {
		t.Fatal("Hello should be world", val)
	}
	val = Get("something", "notdefined")
	if val != "notdefined" {
		t.Fatal("something should be notdefined", val)
	}
	val = GetWarn("another", "not defined")
	if val != "not defined" {
		t.Fatal("another should be not defined", val)
	}
}
