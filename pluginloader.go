package main

import (
	"plugin"
	"os"
	"errors"
)

var (
)

/**
 * Loads a Plugin and executes a start function
 */
func LoadPlugin(filename string, entry_func ...string) (p *plugin.Plugin, err error) {
	var s_err, p_err, l_err error
	var startfunc plugin.Symbol
	if _, s_err = os.Stat(filename); !os.IsNotExist(s_err) {
		p, p_err = plugin.Open(filename)
		if p_err != nil {
			//panic(p_err)
			err = p_err
			return nil, err
		}
		startfuncname := "startfunc"
		if len(entry_func) > 0 {
			startfuncname = entry_func[0]
		}
		startfunc, l_err = p.Lookup(startfuncname)
		if l_err != nil {
			//panic(l_err)
			err = l_err
			return nil, err
		}
		startfunc.(func())()
		return
		
	} else {
		return nil,  errors.New("File not found")
	}
}

/**
 * https://medium.com/learning-the-go-programming-language/writing-modular-go-programs-with-plugins-ec46381ee1a9
 */
func LoadPluginDirect(filename string) (p *plugin.Plugin, err error) {
	var s_err, p_err error
	if _, s_err = os.Stat(filename); !os.IsNotExist(s_err) {
		p, p_err = plugin.Open(filename)
		if p_err != nil {
			//panic(p_err)
			err = p_err
			return nil, err
		}
		return
		
	} else {
		return nil,  errors.New("File not found")
	}
}
