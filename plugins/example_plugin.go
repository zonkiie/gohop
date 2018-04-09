// +build -buildmode=plugin
package main

import "fmt"

// Exported Symbols must always start with upper case!
func Startfunc() { fmt.Printf("Hello, This is a test.\n") }
