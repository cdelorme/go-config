package config

/**
 * DataBag Struct
 *
 * Author: Casey Delow <cdelorme@gmail.com>
 * Date: 2014-8-6
 */

type DataBag struct {
	Data map[string]interface{}
}

func (bag *DataBag) build() {
	if bag.Data == nil {
		bag.Data = make(map[string]interface{})
	}
}

func (bag *DataBag) Set(key string, value interface{}) {
	bag.build()
	bag.Data[key] = value
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

func (bag *DataBag) GetFloat(fallback float32, keys ...string) float32 {
	if res, ok := bag.Get(fallback, keys...).(float32); ok {
		return res
	}
	return fallback
}
