
# go-appconfig

Yet another generic configuration layer to simplify loading and accessing configuration for any application.


## alternatives

Mine is not the only, nor the first; there are many others:

- [miguel branco's goconfig](https://github.com/miguel-branco/goconfig)
- [robfig's extended from miguel's](https://github.com/robfig/config)
- [Unknwon's goconfig](https://github.com/Unknwon/goconfig)

Each aims to deliver a similar goal, make re-usable configuration accessible.


## sales pitch

My library aims to deliver simplicity, featuring:

- works with json
- provides an abstracted interface for dynamic access

What it does not feature:

- unit tests
- complex abstractions
- interfaces
- support for multiple file types and formats

Planned features:

- connect with `flags` for overridable config parameters


## usage

The usage is quite simple, create the config object with the expected config file path:

    conf := config.Config{File: "path/to/conf.json"}

You may now load the configuration:

    conf.Load()

Here you can access contents, supplying a fallback or "default value" if the key is not found, and one or more keys to access deep embedded values:

    anInterfae := conf.Get(nil, "key")
    aString := conf.GetString("", "key", "keyTwo")
    anInt := conf.GetInt(0, "key")
    aFloat := conf.GetFloat(0.0, "key")

You can change the stored data with `Set()`:

    conf.Set("key", "anyvalue")

_It does not allow variable depth, if you want a deeper object, you'll have to assemble it before storing it._

You can save changed configuration via:

    conf.Save()

If you want to can break off a `DataBag` from the configuration like this:

    bag := conf.GetBag("")

_The `DataBag` struct gives you all the abstract methods of interacting with the config values, without the `Save` & `Load` methods, and is convenient when you only want to pass a limited portion of the configuration to dependent code._
