package main

import "github.com/vugu/vugu/js"

func consoleLog(value js.Value) {
	js.Global().Get("console").Call("log", value)
}
