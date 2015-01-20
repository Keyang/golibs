package kxlog

import (
	"fmt"
	"time"
)

func D(format string, args ...interface{}) {
	log(format, "DEBUG", args...)
}
func I(format string, args ...interface{}) {
	log(format, "INFO", args...)
}
func W(format string, args ...interface{}) {
	log(format, "WARN", args...)
}
func E(format string, args ...interface{}) {
	log(format, "ERROR", args...)
}
func C(format string, args ...interface{}) {
	log(format, "CRITICAL", args...)
}
func B(t1, t2 time.Time, format string, args ...interface{}) {
	d := t2.UnixNano() - t1.UnixNano()
	mms := float32(d) / 1000
	a := []interface{}{}
	a = append(a, mms)
	a = append(a, args...)
	log("%.2f "+format, "BENCHMARK", a...)
}
func log(format, level string, args ...interface{}) {
	f := "%s %s " + format + "\n"
	a := []interface{}{}
	a = append(a, time.Now().Format(time.RFC3339))
	a = append(a, level)
	a = append(a, args...)
	fmt.Printf(f, a...)
}
