
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

It's only operations are `Load()` and `Save()`.

The `Load()` operation accepts a `filePath`, but if none is provided it will look in `XDG_CONFIG_DIR`, then in `HOME`, and finally in `/etc/` as standard paths for application configuration.  It will try to find the config file by the application name, then application name plus `.json`, and finally application name plus `.conf`.  It will return a `map[string]interface{}`, or an error.

The `Save()` operation accepts a file path, with an XDG compatible default (eg. `$XDG_CONFIG_DIR/appname/appname`), and a `map[string]interface{}` of data.  It will build the necessary path if it does not exist.  If it fails at any point an error will be returned.

Since both of these deal with an ambiguous `interface{}` returned type, the user is responsible for casting the values correctly.  **Since this is a common interaction in golang, I build a [library for maps](https://github.com/cdelorme/go-maps), which you may also find useful.**

What my library does not have:

- more than 110 lines of code
- unit tests
- complex abstractions
- interfaces
- support for multiple file types and formats

_Whether you believe all of these to be beneficial is a matter of personal preference,_ but given it's size it should be possible to grasp the complete project in your head at a glance.


## usage

To import my library:

	import "github.com/cdelorme/go-config"

To attempt to load from a specified file:

    conf, err := config.Load("config.json")

You can save any `map[string]interface{}` as json via:

    err := config.Save("config.json", aMap)
