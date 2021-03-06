package spec

import (
	"github.com/waybeams/waybeams/pkg/events"
)

type CharCallback func(r rune)

type Window interface {
	ResizableWriter
	ResizableReader

	BeginFrame()
	Close()
	EndFrame()
	FrameRate() int
	Init()
	OnResize(handler events.EventHandler) events.Unsubscriber
	PixelRatio() float64
	PollEvents()
	ShouldClose() bool
	UpdateInput(root ReadWriter)
}

type InputController interface {
	Update(root ReadWriter)
}
