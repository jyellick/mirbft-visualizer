package main

import "syscall/js"

func consoleLog(value js.Value) {
	js.Global().Get("console").Call("log", value)
}
