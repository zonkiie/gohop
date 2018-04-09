package main

import (
	"testing"
	"plugin"
)

func TestLoadPlugin(t *testing.T) {
	example_plugin := "plugins/example_plugin.so"
	startfuncname := "startfunc"
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
