
# go-config

Yet another generic configuration layer to simplify loading and accessing configuration for any application.


## alternatives

Mine is not the only, nor the first; there are many others:

- [miguel branco's goconfig](https://github.com/miguel-branco/goconfig)
- [robfig's extended from miguel's](https://github.com/robfig/config)
- [Unknwon's goconfig](https://github.com/Unknwon/goconfig)

Each aims to deliver a similar goal; make re-usable configuration accessible.


## sales pitch

My config library aims to deliver the simplest usable implementation.

It offers `Load()` and `Save()` as functions.  It expects the path to be supplied, but will fallback to the user homedir and application name (ex. `os.Args[0]`) or in `/etc/`.

It works with `map[string]interface{}` data, saving and returning it as needed.  _The user is responsible for casting indexes, though I have an independent/uncoupled [library for this as well](https://github.com/cdelorme/go-maps)._

What my library does not have:

- more than 70 lines of code
- unit tests
- complex abstractions
- interfaces
- support for multiple file types and formats


## usage

To attempt to load from a specified file:

    conf, err := config.Load("config.json")

_If no file is found, or supplied, it will look in `~/.appname` and `/etc/appname` in that order to attempt to load configuration data.  You can supply an empty string to have it attempt to load from fallback paths.  If no files were found, or it failed to load valid json, an error will be returned._

You can save any `map[string]interface{}` as json via:

    err := config.Save("config.json", &aMap)

_The map file is supplied by reference.  if the supplied file is an empty string it will attempt to save to `~/.appname`, but it will not attempt to save to `/etc/appname`._
