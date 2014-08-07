package config

/**
 * DataBag Struct
 *
 * Author: Casey Delow <cdelorme@gmail.com>
 * Date: 2014-8-7
 */

type DataBag struct {
	Data map[string]interface{}
}

func (bag *DataBag) Set(value interface{}, keys ...string) {
	if bag.Data == nil {
		bag.Data = make(map[string]interface{})
	}
	if len(keys) == 1 {
		bag.Data[keys[0]] = value
	} else if len(keys) > 1 {
		bag.setDeep(&bag.Data, value, keys...)
	}
}

func (bag *DataBag) setDeep(current *map[string]interface{}, value interface{}, keys ...string) {
	if _, ok := (*current)[keys[0]]; !ok {
		(*current)[keys[0]] = make(map[string]interface{})
	}
	if len(keys) > 1 {
		next := (*current)[keys[0]].(map[string]interface{})
		bag.setDeep(&next, value, keys[1:]...)
	} else {
		(*current)[keys[0]] = value
	}
}

func (bag *DataBag) Get(fallback interface{}, keys ...string) interface{} {
	if len(keys) > 1 {
		return bag.Get(fallback, keys[1:]...)
	} else if len(keys) == 1 {
		if data, exists := bag.Data[keys[0]]; exists {
			return data
		}
	}
	return fallback
}

func (bag *DataBag) GetBag(keys ...string) DataBag {
	data := bag.Get(nil, keys...)
	if res, ok := data.(map[string]interface{}); ok {
		return DataBag{Data: res}
	}
	return DataBag{}
}

func (bag *DataBag) GetString(fallback string, keys ...string) string {
	if res, ok := bag.Get(fallback, keys...).(string); ok {
		return res
	}
	return fallback
}

func (bag *DataBag) GetInt(fallback int, keys ...string) int {
	if res, ok := bag.Get(fallback, keys...).(int); ok {
		return res
	}
	return fallback
}

func (bag *DataBag) GetFloat32(fallback float32, keys ...string) float32 {
	if res, ok := bag.Get(fallback, keys...).(float32); ok {
		return res
	}
	return fallback
}

func (bag *DataBag) GetBool(fallback bool, keys ...string) bool {
	if res, ok := bag.Get(fallback, keys...).(bool); ok {
		return res
	}
	return fallback
}
