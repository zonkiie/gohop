package main

import (
	"github.com/vharitonsky/iniflags"
	"os"
	"os/user"
	"fmt"
	"flag"
)

var (
	cWs = flag.Bool("Webservice", false, "Start Web Service")
	debugOutput = flag.Bool("DebugOutput", false, "Print debug values")
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

func main() {
	prog_start()
}

