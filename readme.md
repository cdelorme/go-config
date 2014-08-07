
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

    conf.Set("anyvalue", "key")

Data can be set deeper, like this:

    conf.Set("content", "deep", "seeded")

_While deep values are allowed, it only supports `map[string]interface` for deep items, it won't change struct fields, nor can it append new fields to them._

You can save changed configuration via:

    conf.Save()

If you want to can break off `Settings` from the configuration like this:

    settings := conf.GetSettings("")

_The `Settings` struct carries all of the abstract Get/Set methods, but none of the `Config` structs file (`Save` & `Load`) or flag methods.  It carries a pointer to the originals data, so changes made to it will be applied back to the source.  This is handy if you want to supply limited access to a subset of configuration to another struct or library._


## nuances & bugs

- if you store a struct and attempt to change an inner value, or add new values it will return an error
- the flags abstraction is very basic and treats all supplied values as strings; config will cast them as you desire
- the code will now allow you to store Config inside Config (it can cause a cyclic error), and will return an error
