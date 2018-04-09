package main

import (
	"testing"
	"plugin"
	"fmt"
)

func TestLoadPlugin(t *testing.T) {
	example_plugin := "plugins/example_plugin.so"
	startfuncname := "Startfunc"
	var p *plugin.Plugin
	var e error
	p, e = LoadPluginDirect(example_plugin)
	if e != nil {
		t.Errorf("Failed to load plugin %s\n", example_plugin)
	}
	startfunc, err := p.Lookup(startfuncname)
	if err != nil {
		t.Errorf("Failed to find function %s\n", startfuncname)
	}
	startfunc.(func())()
	
	t.Logf("Test finished.")
}

/**
 * this function works
 */
func TestLoadStartPlugin() {
	p, p_err := LoadPluginDirect("plugins/example_plugin.so")
	if p_err != nil {
		panic(fmt.Sprintf("Failed to load plugin. Error is: %s\n", p_err.Error()))
	}
	fmt.Printf("Plugin: %T %v\n", p, p)
	startfuncname := "Startfunc"
	startfuncsymbol, err := p.Lookup(startfuncname)
	if err != nil {
		panic(fmt.Sprintf("Failed to find function %s. Error is: %s\n", startfuncname, err.Error()))
	}
	startfunc, ok := startfuncsymbol.(func())
	if !ok {
		panic("Could not map startfunc\n")
	}
	fmt.Print("Executing startfunc\n")
	startfunc()
}
