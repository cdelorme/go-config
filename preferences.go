package config

import (
	"errors"
	"flag"
)

type Preferences struct {
	Options []Option
}

func (pref *Preferences) Option(name string, fallback interface{}, desc string, keys ...string) error {
	fb, ok := fallback.(string)
	if !ok {
		return errors.New("unable to accept provided fallback")
	}
	o := Option{Name: name, Description: desc, Keys: keys}
	o.Value = flag.String(name, fb, desc)
	pref.Options = append(pref.Options, o)
	return nil
}

func (pref *Preferences) Parse() {
	flag.Parse()
}
