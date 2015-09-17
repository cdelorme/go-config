
# go-config

Yet another generic configuration layer to simplify loading and accessing configuration for any application.


## alternatives

Mine is not the only, nor the first; there are many others:

- [miguel branco's goconfig](https://github.com/miguel-branco/goconfig)
- [robfig's extended from miguel's](https://github.com/robfig/config)
- [Unknwon's goconfig](https://github.com/Unknwon/goconfig)

Each aims to deliver a similar goal of accessing configuration via a module.


## sales pitch

My config library aims to deliver the simplest re-usable cross-platform implementation for dealing with json configuration.

It has two operations:

- `Load()`
- `Save()`

It builds a list of standard paths based on environment variables, and provides the default path via the `ConfigPath` package variable.  It automatically attempts extensionless plus `.json` and `.conf` naming conventions.

Both methods may return an error, and `Load()` returns a go native type `map[string]interface{}`, giving you the flexibility to cast and work with data as desired.  A [support library](https://github.com/cdelorme/go-maps) can be used to simplify conversion or access.

Things my library does not include:

- more than 110 lines of code
- complex abstractions
- interfaces
- support for many file types/formats
- unit tests

_Whether you believe all of these to be beneficial is a matter of personal preference,_ but given it's size it should be possible to grasp the complete project in your head at a glance.  This makes it a breeze to pickup and use confidently.

**This package was intended to be used alongside the [go-option](https://github.com/cdelorme/go-option) and [go-env](https://github.com/cdelorme/go-env) packages to handle application configuration.**


## standard configuration paths

This library works with the following prioritized list of standard storage paths:

- `{appName}{.conf,.json}`
- `{APPDATA}/{appName}{.conf,.json}`
- `{XDG_CONFIG_DIR}/{appName}{.conf,.json}`
- `{HOME}/.{appName}{.conf,.json}`
- `/etc/{appName}{.conf,.json}`
- `/etc/{appName}/{appName}{.conf,.json}`

When you run `Load()` with no arguments, it will search this list returning the first file found.  Similarly it will use `APPDATA`/`XDG_CONFIG_DIR` as the default `Save()` path when no arguments are supplied.


## usage

To import my library:

	import "github.com/cdelorme/go-config"

To attempt to load from a specified file:

    conf, err := config.Load("config.json")

_Load will remember what file configuration was loaded from in the package-variable `ConfigFile`, which by default uses `APPDATA` or `XDG_CONFIG_DIR` standard paths._

You can save any object that can be marshaled into json via:

    err := config.Save(aMap, "config.json")

_If no file is supplied, it will attempt to use `ConfigFile`.  By default json marshal uses pretty-printed output for readability._
