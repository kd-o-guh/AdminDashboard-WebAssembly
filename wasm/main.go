package main

import (
	"syscall/js"
)

func main() {
	doc := js.Global().Get("document")

	sideMenu := doc.Call("querySelector", "aside")
	menuBtn := doc.Call("getElementById", "menu-btn")
	CloseBtn := doc.Call("getElementById", "close-btn")
	themeToggler := doc.Call("getElementById", "theme-toggler")

	done := make(chan struct{})

	menuClick := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		sideMenu.Get("style").Set("display", "block")
		return nil

	})
	defer menuClick.Release()

	closeClick := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		sideMenu.Get("style").Set("display", "none")
		return nil
	})
	defer closeClick.Release()

	themeToggle := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		doc.Get("body").Get("classList").Call("toggle", "dark-theme-variables")
		themeToggler.Call("querySelector", "span:nth-child(1)").Get("classList").Call("toggle", "active")
		themeToggler.Call("querySelector", "span:nth-child(2)").Get("classList").Call("toggle", "active")
		return nil
	})
	defer themeToggle.Release()

	menuBtn.Call("addEventListener", "click", menuClick)
	CloseBtn.Call("addEventListener", "click", closeClick)
	themeToggler.Call("addEventListener", "click", themeToggle)

	<-done

}
