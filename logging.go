//go:build !use_macro
// +build !use_macro

package main

import "fmt"

// This file will be used when the macro is not enabled.

func LOG(format string, args ...interface{}) {
	if LoggingEnabled {
		fmt.Printf("LOG(func): "+format+"\n", args...)
	}
}
