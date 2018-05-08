package controls

import (
	"ui/comp"
	"ui/opts"
	"os"
	"runtime"
	"testing"
	"ui"
)

var TestComponent = comp.Define("TestComponent", comp.New)

type FakeComponent struct {
	comp.Component
}

func NewFake() *FakeComponent {
	return &FakeComponent{}
}

// Create a new factory using our component creation function reference.
var Fake = comp.Define("Fake",
	func() ui.Displayable { return NewFake() })

func TestMain(m *testing.M) {
	// This is required if any test uses a OpenTestWindow
	runtime.LockOSThread()
	os.Exit(m.Run())
}

// OpenTestWindow is a component option that will create and launch a new window with your
// component instance displayed inside of it. Any ComponentOptions provided to the call
// will be applied to the newly created window object.
func OpenTestWindow(userOptions ...ui.Option) ui.Option {
	return func(d ui.Displayable) {
		options := []ui.Option{
			opts.Width(800),
			opts.Height(600),
			opts.HAlign(ui.AlignCenter),
			opts.VAlign(ui.AlignCenter),
			opts.Children(func(c ui.Context, w ui.Displayable) {
				w.AddChild(d)
			}),
		}
		options = append(options, userOptions...)

		win := NanoWindow(d.Context(), options...)
		win.(*NanoWindowComponent).Listen()
	}
}