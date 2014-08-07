package config

/**
 * Settings Struct
 *
 * Author: Casey Delow <cdelorme@gmail.com>
 * Date: 2014-8-7
 */

import (
	"errors"
)

type Settings struct {
	Data *map[string]interface{}
}

func (settings *Settings) Set(value interface{}, keys ...string) error {
	if _, ok := value.(Config); ok {
		return errors.New("storing config inside config is not allowed")
	}
	if settings.Data == nil {
		data := make(map[string]interface{})
		settings.Data = &data
	}
	if len(keys) == 1 {
		(*settings.Data)[keys[0]] = value
	} else if len(keys) > 1 {
		return settings.setDeep(settings.Data, value, keys...)
	}
	return nil
}

func (settings *Settings) setDeep(current *map[string]interface{}, value interface{}, keys ...string) error {
	if _, ok := (*current)[keys[0]]; !ok {
		(*current)[keys[0]] = make(map[string]interface{})
	}
	if len(keys) > 1 {
		next, ok := (*current)[keys[0]].(map[string]interface{})
		if !ok {
			return errors.New("failed to cast type to map[string]interface, was it a struct?")
		}
		settings.setDeep(&next, value, keys[1:]...)
	} else {
		(*current)[keys[0]] = value
	}
	return nil
}

func (settings *Settings) Get(fallback interface{}, keys ...string) interface{} {
	if len(keys) > 1 {
		return settings.Get(fallback, keys[1:]...)
	} else if len(keys) == 1 {
		if data, exists := (*settings.Data)[keys[0]]; exists {
			return data
		}
	}
	return fallback
}

func (settings *Settings) GetSettings(keys ...string) Settings {
	data := settings.Get(nil, keys...)
	if res, ok := data.(map[string]interface{}); ok {
		return Settings{Data: &res}
	}
	return Settings{}
}

func (settings *Settings) GetString(fallback string, keys ...string) string {
	if res, ok := settings.Get(fallback, keys...).(string); ok {
		return res
	}
	return fallback
}

func (settings *Settings) GetInt(fallback int, keys ...string) int {
	if res, ok := settings.Get(fallback, keys...).(int); ok {
		return res
	}
	return fallback
}

func (settings *Settings) GetFloat32(fallback float32, keys ...string) float32 {
	if res, ok := settings.Get(fallback, keys...).(float32); ok {
		return res
	}
	return fallback
}

func (settings *Settings) GetBool(fallback bool, keys ...string) bool {
	if res, ok := settings.Get(fallback, keys...).(bool); ok {
		return res
	}
	return fallback
}
