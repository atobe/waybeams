
## This software is experimental. Contributions are welcome, but production use is discouraged

![CircleCi](https://circleci.com/gh/waybeams/waybeams.svg?style=shield&circle-token=:circle-token)

# Waybeams

With Waybeams you can quickly and reliably create and test tiny (<4MB). delightful applications that can thrive on a multitude of surfaces (Windows, macOS, Linux, Android, iOS, Raspberry Pi, Beaglebone, etc.).

According to Merriam Webster, A Waybeam is, ": a beam supporting a way; specifically : either of two longitudinal beams resting on transverse girders and supporting the rails of a road crossing a bridge"

We like to think of Waybeams (the tools here) as providing a solid structural foundation that makes it possible for us to safely and quickly transport enormous quantities of user facing features to production.

![Waybeams Image](media/waybeams-home.jpg)

_[Image](https://www.flickr.com/photos/charlyamato/13417543435/) provided courtesy of [Carlos Amato](https://www.flickr.com/photos/charlyamato/) and the [Creative Commons](https://creativecommons.org/licenses/by-nc-nd/2.0/) license(s)._

Waybeams is built using the [Go language](https://golang.org/) and an OpenGL rendering surface (currently, [NanoVGO](https://github.com/shibukawa/nanovgo))

Waybeams provides (or for now, aspires to provide):
* Simple, fast, reactive and composable GUI toolkit
* Cross-platform, GPU accelerated drawing surface
* Pure Go component declaration and configuration
* Tiny, blazing fast, constraints-based layouts
* Component Trait assignment using web-like selectors
* Headless (insanely fast) environment for UI tests
* Isolated visual environment for test-driven development on UI components
* Automated image rendering surface (from tests) for release validation

# Getting Started

Step one: Install some version of Go to the system > 1.4. This version will be used purely to build our local version.

Once this is done, change to this directory and run the following (on Darwin/macOS or Linux):

```bash
source setup-env.sh
make dev-install
```

# What is, "pure Go component declaration?"
Rather than using some separate language (or many languages) to describe a user interface, we use _one_. We use pure Go to describe behavior, style _and_ structure. A core thesis of this work, is that this decision alone can deliver significant reductions in cognitive load, development time and even runtime performance.

A very simple Waybeams application might look something like the following:
```go
package main

import (
  . "display"
  "runtime"
)

func main() {
  runtime.LockOSThread()
  win, _ := NanoWindow(NewBuilder(), Title("Hello World"), Children(func(b Builder) {
    Label(b, FlexWidth(1), Height(40), FontSize(36), Text("Hello World!"))
  }))

  win.(Window).Init()
}
```

There are other benefits to this approach. Our compile times are exceptionally fast, component trees build extremely fast as there is no translation layer and the more we conceive of these declarations as pure functions, the simpler and more stable our applications become.

Unlike just about every other graphical toolkti, these components also happen to be extremely amenable to automated testing. Even though we can exercise components using automated tests, we also have the ability to instantly launch a visual environment that allows you to explore a component instantiation from an individual test case.

# Development environment
This project is being actively developed on OS X and Linux and should build properly in both environments.

If you're on Windows, and interested in contributing, you'll need to get some kind of Posix-ish environment working (probably MingW these days?), and things should probably mostly work? Please let us know if there's anything we can do to help.

### Manual Prerequisites
You'll need to get the following installed on your computer in order to proceed:
* Git
* Make
* Any version of Go (since Go is now bootstrapped)

### Notes
We will download, build and install a specific version of Go and any other tools into the `${PROJECT_PATH}/lib` folder to ensure that all development happens against the same version everywhere.

We are currently building against the [Nanovg](https://github.com/memononen/nanovg) 2d drawing library. I expect this to change in the future in order to deliver  support for rich text rendering. The primary option being considered is [Skia](https://skia.org/), but has not been integrated because the C interface is incomplete and notoriously difficult to rationalize and the build dependencies are also non-trivial. I also spent quite a lot of time getting Cairo working, but ran into difficulties with Pango (rich text layout) and temporarily gave up.

If you have experience with Skia or Cairo/Pango, and would like to contribute, please let us know!

## Download and install
```
mkdir waybeams; cd waybeams
git clone https://github.com/waybeams/waybeams.git .
source setup-env.sh
make dev-install
```

## Run tests
```
make test
```
Or to get verbose test output:
```
make test-v
```

I use a Python script called, "[when-changed.py](https://github.com/joh/when-changed)" to watch source files and re-run `make test` whenever a file changes. I place the file in my path and use the following command:
```
when-changed.py src/**/*.go  -c "make test"
```

## Build & run binary for development
This should build the binaries from the latest sources on your computer
```
make build
```
Or to build & run in one step
```
make run
make run-anim
```
Build artifacts can be found in `./out`
