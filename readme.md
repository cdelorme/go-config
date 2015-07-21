
# go-config

Yet another generic configuration layer to simplify loading and accessing configuration for any application.


## alternatives

Mine is not the only, nor the first; there are many others:

- [miguel branco's goconfig](https://github.com/miguel-branco/goconfig)
- [robfig's extended from miguel's](https://github.com/robfig/config)
- [Unknwon's goconfig](https://github.com/Unknwon/goconfig)

Each aims to deliver a similar goal of accessing configuration via a module.


## sales pitch

My config library aims to deliver the simplest usable implementation.

It's only operations are `Load()` and `Save()`.  These operations use XDG default paths when no path is supplied.  The `Save()` operation will use `~/.config/appname/appname`, and load will look for `~/.config/appname/appname`, then `~/.appname`, and finally `/etc/appname`.  _The `Load()` operation will also look for .json and .conf extensions at each path before giving up._

If either fails, it will return an error.

Load will return a `map[string]interface{}`, giving you the flexibility of casting to expected types.  A [map support library](https://github.com/cdelorme/go-maps) is also available to help with basic casting and merging.

What my library does not have:

- more than 110 lines of code
- unit tests
- complex abstractions
- interfaces
- support for multiple file types and formats

_Whether you believe all of these to be beneficial is a matter of personal preference,_ but given it's size it should be possible to grasp the complete project in your head at a glance.  This makes it a breeze to pickup and use confidently.


## usage

To import my library:

	import "github.com/cdelorme/go-config"

To attempt to load from a specified file:

    conf, err := config.Load("config.json")

You can save any `map[string]interface{}` as json via:

    err := config.Save("config.json", aMap)

