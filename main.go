package main

import (
	"github.com/vharitonsky/iniflags"
	"github.com/zonkiie/gohop_libs"
	"os"
	"os/user"
	"fmt"
	"flag"
//	"plugin"
)

var (
	cWs = flag.Bool("Webservice", false, "Start Web Service")
	debugOutput = flag.Bool("DebugOutput", false, "Print debug values")
	pluginDir = flag.String("pluginDir", "plugins", "Plugin Directory")
	defaultConfigFile string
	err interface{}
)

func prog_start() {
	CUser, uerror := user.Current()
	if uerror != nil {
		panic(fmt.Sprintf("Could not determine User! Error: %v", uerror))
	}
	defaultConfigFile = CUser.HomeDir + "/.gohop"
	iniflags.SetAllowMissingConfigFile(true)
	/// @see https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
	if _, err := os.Stat(defaultConfigFile); !os.IsNotExist(err) {
		iniflags.SetConfigFile(defaultConfigFile)
		if *debugOutput { fmt.Fprintf(os.Stderr, "Read from Default Configfile: %s\n", defaultConfigFile) }
	} else  {
		if *debugOutput { fmt.Fprintf(os.Stderr, "No Default Configfile %s found.\n", defaultConfigFile) }
	}
	
	iniflags.Parse()
	
}	

func LoadStartPlugin() {
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

func main() {
	LoadStartPlugin()
	prog_start()
}

