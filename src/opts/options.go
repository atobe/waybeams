package opts

import (
	"events"
	"ui"
)

// ActualWidth will set Component.ActualWidth.
func ActualWidth(value float64) ui.Option {
	return func(d ui.Displayable) error {
		d.SetActualWidth(value)
		return nil
	}
}

// ActualHeight will set Component.ActualHeight.
func ActualHeight(value float64) ui.Option {
	return func(d ui.Displayable) error {
		d.SetActualHeight(value)
		return nil
	}
}

func BgColor(color int) ui.Option {
	return func(d ui.Displayable) error {
		d.SetBgColor(color)
		return nil
	}
}

func Blurred() ui.Option {
	return func(d ui.Displayable) error {
		d.Blur()
		return nil
	}
}

// Children will compose child components onto the current component. The composer
// type must be a function with a signature that matches one of the following:
//   A) func()
//   B) func(b Builder)
//   C) func(d Displayable)
//   D) func(b Builder, d Displayable)
// The outermost Children function usually should receive a builder instance that
// all children will receive and isolated Component definitions generally require
// both arguments to the outer composer.
func Children(composer interface{}) ui.Option {
	return func(d ui.Displayable) error {
		return d.Composer(composer)
	}
}

func Data(data interface{}) ui.Option {
	return func(d ui.Displayable) error {
		d.SetData(data)
		return nil
	}
}

// ExcludeFromLayout will configure Component.ExcludeFromLayout.
func ExcludeFromLayout(value bool) ui.Option {
	return func(d ui.Displayable) error {
		d.SetExcludeFromLayout(value)
		return nil
	}
}

// FlexHeight will set Component.FlexHeight.
func FlexHeight(value float64) ui.Option {
	return func(d ui.Displayable) error {
		d.SetFlexHeight(value)
		return nil
	}
}

// FlexWidth will set Component.FlexWidth.
func FlexWidth(value float64) ui.Option {
	return func(d ui.Displayable) error {
		d.SetFlexWidth(value)
		return nil
	}
}

func Focused() ui.Option {
	return func(d ui.Displayable) error {
		d.Focus()
		return nil
	}
}

func FontColor(color int) ui.Option {
	return func(d ui.Displayable) error {
		d.SetFontColor(color)
		return nil
	}
}

func FontFace(face string) ui.Option {
	return func(d ui.Displayable) error {
		d.SetFontFace(face)
		return nil
	}
}

func FontSize(size int) ui.Option {
	return func(d ui.Displayable) error {
		d.SetFontSize(size)
		return nil
	}
}

func Gutter(value float64) ui.Option {
	return func(d ui.Displayable) error {
		d.SetGutter(value)
		return nil
	}
}

// ID will set the Component.Id.
func ID(value string) ui.Option {
	return func(d ui.Displayable) error {
		d.Model().ID = value
		return nil
	}
}

func IsFocusable(value bool) ui.Option {
	return func(d ui.Displayable) error {
		d.SetIsFocusable(value)
		return nil
	}
}

func IsText(value bool) ui.Option {
	return func(d ui.Displayable) error {
		d.SetIsText(value)
		return nil
	}
}

func IsTextInput(value bool) ui.Option {
	return func(d ui.Displayable) error {
		d.SetIsTextInput(value)
		return nil
	}
}

// HAlign will set Component.HAlign.
func HAlign(align ui.Alignment) ui.Option {
	return func(d ui.Displayable) error {
		d.SetHAlign(align)
		return nil
	}
}

// Height will set Component.Height.
func Height(value float64) ui.Option {
	return func(d ui.Displayable) error {
		// TODO(lbayes): Should use accessor!
		d.Model().Height = value
		return nil
	}
}

func Key(value string) ui.Option {
	return func(d ui.Displayable) error {
		d.Model().Key = value
		return nil
	}
}

// LayoutType will set Component.LayoutType.
func LayoutType(layoutType ui.LayoutTypeValue) ui.Option {
	return func(d ui.Displayable) error {
		d.SetLayoutType(layoutType)
		return nil
	}
}

// MaxHeight will set Component.MaxHeight.
func MaxHeight(value float64) ui.Option {
	return func(d ui.Displayable) error {
		d.SetMaxHeight(value)
		return nil
	}
}

// MaxWidth will set Component.MaxWidth.
func MaxWidth(value float64) ui.Option {
	return func(d ui.Displayable) error {
		d.SetMaxWidth(value)
		return nil
	}
}

// MinHeight will set Component.MinHeight.
func MinHeight(value float64) ui.Option {
	return func(d ui.Displayable) error {
		d.SetMinHeight(value)
		return nil
	}
}

// MinWidth will set Component.MinWidth.
func MinWidth(value float64) ui.Option {
	return func(d ui.Displayable) error {
		d.SetMinWidth(value)
		return nil
	}
}

// Padding will set Component.Padding, which will effectively set padding for
// all four sides as well (bottom, top, left, right, horizontal and vertical).
func Padding(value float64) ui.Option {
	return func(d ui.Displayable) error {
		d.SetPadding(value)
		return nil
	}
}

// PaddingBottom will set Component.PaddingBottom.
func PaddingBottom(value float64) ui.Option {
	return func(d ui.Displayable) error {
		d.SetPaddingBottom(value)
		return nil
	}
}

// PaddingLeft will set Component.PaddingLeft.
func PaddingLeft(value float64) ui.Option {
	return func(d ui.Displayable) error {
		d.SetPaddingLeft(value)
		return nil
	}
}

// PaddingRight will set Component.PaddingRight.
func PaddingRight(value float64) ui.Option {
	return func(d ui.Displayable) error {
		d.SetPaddingRight(value)
		return nil
	}
}

// PaddingTop will set Component.PaddingTop.
func PaddingTop(value float64) ui.Option {
	return func(d ui.Displayable) error {
		d.SetPaddingTop(value)
		return nil
	}
}

// PrefHeight will set Component.PrefHeight.
func PrefHeight(value float64) ui.Option {
	return func(d ui.Displayable) error {
		d.SetPrefHeight(value)
		return nil
	}
}

// PrefWidth will set Component.PrefWidth.
func PrefWidth(value float64) ui.Option {
	return func(d ui.Displayable) error {
		d.SetPrefWidth(value)
		return nil
	}
}

// Size will set Component.Width and Component.Height.
func Size(width, height float64) ui.Option {
	return func(d ui.Displayable) error {
		d.SetWidth(width)
		d.SetHeight(height)
		return nil
	}
}

func StrokeColor(color int) ui.Option {
	return func(d ui.Displayable) error {
		d.SetStrokeColor(color)
		return nil
	}
}

func StrokeSize(size int) ui.Option {
	return func(d ui.Displayable) error {
		d.SetStrokeSize(size)
		return nil
	}
}

func Text(value string) ui.Option {
	return func(d ui.Displayable) error {
		// TODO(lbayes): Sanitize text as user input values can be placed in here.
		// TODO(lbayes): Localize text using Localization map.
		d.SetText(value)
		return nil
	}
}

// Title will set Component.Title.
func Title(value string) ui.Option {
	return func(d ui.Displayable) error {
		d.SetTitle(value)
		return nil
	}
}

func TraitNames(names ...string) ui.Option {
	return func(d ui.Displayable) error {
		d.SetTraitNames(names...)
		return nil
	}
}

// VAlign will set Component.VAlign.
func VAlign(align ui.Alignment) ui.Option {
	return func(d ui.Displayable) error {
		d.SetVAlign(align)
		return nil
	}
}

func View(view ui.RenderHandler) ui.Option {
	return func(d ui.Displayable) error {
		d.SetView(view)
		return nil
	}
}

func Visible(visible bool) ui.Option {
	return func(d ui.Displayable) error {
		d.SetVisible(visible)
		return nil
	}
}

// Width will set Component.Width.
func Width(value float64) ui.Option {
	return func(d ui.Displayable) error {
		d.Model().Width = value
		return nil
	}
}

// X will set Component.X.
func X(pos float64) ui.Option {
	return func(d ui.Displayable) error {
		d.SetX(pos)
		return nil
	}
}

// Y will set Component.Y.
func Y(pos float64) ui.Option {
	return func(d ui.Displayable) error {
		d.SetY(pos)
		return nil
	}
}

// Z will set Component.Z.
func Z(pos float64) ui.Option {
	return func(d ui.Displayable) error {
		d.SetZ(pos)
		return nil
	}
}

// Bag is a collection of Options.
func Bag(opts ...ui.Option) ui.Option {
	return func(d ui.Displayable) error {
		for _, opt := range opts {
			opt(d)
		}
		return nil
	}
}

//-------------------------------------------
// Event Helpers
//-------------------------------------------

// On will apply the provided handler to the provided event name.
func On(eventName string, handler events.EventHandler) ui.Option {
	return func(d ui.Displayable) error {
		d.PushUnsub(d.On(eventName, handler))
		return nil
	}
}

func OnClick(handler events.EventHandler) ui.Option {
	return func(d ui.Displayable) error {
		d.PushUnsub(d.On(events.Clicked, handler))
		return nil
	}
}

func OnEnter(handler events.EventHandler) ui.Option {
	return func(d ui.Displayable) error {
		d.PushUnsub(d.On(events.Entered, handler))
		return nil
	}
}

func OnFrameEntered(handler events.EventHandler) ui.Option {
	return func(d ui.Displayable) error {
		d.PushUnsub(d.Context().OnFrameEntered(handler))
		return nil
	}
}

//-------------------------------------------
// State Helpers
//-------------------------------------------

// TODO(lbayes): Consider introducing AppendState and ReplaceState
func OnState(name string, options ...ui.Option) ui.Option {
	return func(d ui.Displayable) error {
		d.OnState(name, options...)
		return nil
	}
}

func SetState(name string) ui.Option {
	return func(d ui.Displayable) error {
		d.SetState(name)
		return nil
	}
}
