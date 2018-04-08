package main

import (
	"plugin"
	"os"
)

var (
)

func loadPlugin(filename string, entry_func ...string) {
	if _, s_err := os.Stat(filename); !os.IsNotExist(s_err) {
		p, p_err := plugin.Open(filename)
		if p_err != nil {
			panic(p_err)
		}
		startfuncname := "startfunc"
		if len(entry_func) > 0 {
			startfuncname = entry_func[0]
		}
		startfunc, l_err := p.Lookup(startfuncname)
		if l_err != nil {
			panic(l_err)
		}
	}
}
