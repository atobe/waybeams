package main

import (
	"github.com/fogleman/ease"
	"runtime"
	. "ui"
	"uiold/context"
	. "ui/controls"
	. "uiold/opts"
)

func init() {
	// We need to do this so that our interactions with CGO (NanoVG/OpenGL) are
	// synchronous.
	runtime.LockOSThread()
}

func createWindow() Displayable {
	return NanoWindow(
		context.New(),

		ID("nano-window"),
		Width(800),
		Height(600),
		Children(func(c Context) {
			// var currentMove ComponentOption
			// moveLeft := Transition(X, 700.0, 0.0, 2000, ease.InOutCubic)
			Trait(c, ".move",
				ExcludeFromLayout(true),
				Transition(c,
					Width,
					25.0,
					100.0,
					1800,
					ease.InOutCubic),
				Transition(c,
					Height,
					25.0,
					100.0,
					1800,
					ease.InOutCubic),
				Transition(c,
					Y,
					0.0,
					490.0,
					2000,
					ease.InOutCubic),
				Transition(c,
					X,
					0.0,
					690.0,
					2000,
					ease.InOutCubic))
			Button(c,
				ID("moving-box"),
				ExcludeFromLayout(true),
				TraitNames("move"),
				Y(200),
				BgColor(0xeebb99ff),
				Width(100),
				Height(100))
		}))
}

func main() {
	win := createWindow()
	win.(*NanoWindowControl).Listen()
}