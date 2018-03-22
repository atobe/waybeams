package builder

import "display"

type ComponentFactory func(s display.Surface)

type GlfwWindowHint int

const (
	AutoIconify = iota
	Decorated
	Floating
	Focused
	Iconified
	Maximized
	Resizable
	Visible
)

/*
// TODO(lbayes) Map local hints to glfw library hints
const (
	Focused     GlfwWindowHint = C.GLFW_FOCUSED      // Specifies whether the window will be given input focus when created. This hint is ignored for full screen and initially hidden windows.
	Iconified   Hint = C.GLFW_ICONIFIED    // Specifies whether the window will be minimized.
	Maximized   Hint = C.GLFW_MAXIMIZED    // Specifies whether the window is maximized.
	Visible     Hint = C.GLFW_VISIBLE      // Specifies whether the window will be initially visible.
	Resizable   Hint = C.GLFW_RESIZABLE    // Specifies whether the window will be resizable by the user.
	Decorated   Hint = C.GLFW_DECORATED    // Specifies whether the window will have window decorations such as a border, a close widget, etc.
	Floating    Hint = C.GLFW_FLOATING     // Specifies whether the window will be always-on-top.
	AutoIconify Hint = C.GLFW_AUTO_ICONIFY // Specifies whether fullscreen windows automatically iconify (and restore the previous video mode) on focus loss.
)

*/

const DefaultFrameRate = 60
const DefaultWidth = 1024
const DefaultHeight = 768
const DefaultTitle = "Default Title"

type SurfaceTypeName int

const (
	CairoSurface = iota
	ImageSurface
	FakeSurface
)

type Option func(b *builder) error

type windowHint struct {
	name  GlfwWindowHint
	value interface{}
}

// Surface Option for Builder
func SurfaceType(surfaceType SurfaceTypeName) Option {
	return func(b *builder) error {
		b.surfaceTypeName = surfaceType
		return nil
	}
}

func FrameRate(fps int) Option {
	return func(b *builder) error {
		b.frameRate = fps
		return nil
	}
}

func Size(width int, height int) Option {
	return func(b *builder) error {
		b.width = width
		b.height = height
		return nil
	}
}

func WindowHint(hintName GlfwWindowHint, value interface{}) Option {
	wHint := &windowHint{
		name:  hintName,
		value: value,
	}

	return func(b *builder) error {
		// First remove existing hint by same name if found
		b.removeWindowHint(hintName)

		b.windowHints = append(b.windowHints, wHint)
		return nil
	}
}

func Title(title string) Option {
	return func(b *builder) error {
		b.title = title
		return nil
	}
}
