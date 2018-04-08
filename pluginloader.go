package main

import (
	"plugin"
	"os"
)

var (
)

func loadPlugin(filename string) {
	if _, s_err := os.Stat(filename); !os.IsNotExist(s_err) {
		p, p_err := plugin.Open(filename)
	}
}
