package kxenv

import (
	"os"

	"github.com/Keyang/golibs/kxlog"
)

var (
	dynamic  = Environment{}
	hardcode = Environment{}
)

type Environment map[string]string

func Set(key string, val string) {
	dynamic[key] = val
}

func Get(key string, def string) (val string) {
	val, ok := dynamic[key]
	if ok == false {
		val = os.Getenv(key)
		if val == "" {
			val, ok = hardcode[key]
			if ok == false {
				val = def
			}
		}
	}
	return
}

func GetWarn(key string, def string) (val string) {
	val, ok := dynamic[key]
	if ok == false {
		val = os.Getenv(key)
		if val == "" {
			val, ok = hardcode[key]
			if ok == false {
				val = def
				kxlog.W("Environment var not found for: %s, use default value: %s", key, def)
			}
		}
	}
	return
}
