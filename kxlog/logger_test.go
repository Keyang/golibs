package kxlog

import "testing"

func TestLog(t *testing.T) {
	D("this is a debug msg:%d", 1234)
	I("this is a info msg:%d", 1234)
	W("this is a warn msg:%d", 1234)
	E("this is a Error msg:%d", 1234)
	C("this is a Critical msg:%d", 1234)
}
