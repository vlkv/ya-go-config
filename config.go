package ya_go_config

import (
	"encoding/json"
	"os"
	"fmt"
	"reflect"
	"github.com/spf13/cast"
)

type Config struct {
	defaults map[string]interface{}
	values map[string]interface{}
}

func New(defaults map[string]interface{}) *Config {
	var res = new(Config)
	if defaults != nil {
		res.defaults = defaults
	} else {
		res.defaults = make(map[string]interface{})
	}
	res.values = make(map[string]interface{})
	return res
}

func (this *Config) Load(jsonFilePath string) {
	file, _ := os.Open(jsonFilePath)
	decoder := json.NewDecoder(file)
	var err = decoder.Decode(&this.values)
	if err != nil {
		panic(fmt.Sprintf("Cannot load config file %s, reason: %v", jsonFilePath, err))
	}
}

func (this *Config) getValueOrDefault(key string) interface{} {
	var val, ok = this.values[key]
	if !ok {
		val, ok = this.defaults[key]
		if !ok {
			panic(fmt.Sprintf("Key %s is not found in config nor in configDefaults!", key))
		}
	}
	return val
}

func (this *Config) GetStr(key string) string {
	var val = this.getValueOrDefault(key)
	var res, err = cast.ToStringE(val)
	if err != nil {
		panic(fmt.Sprintf("%v is not a string! It is %v", val, reflect.TypeOf(val)))
	}
	return res
}

func (this *Config) GetInt(key string) int {
	var val = this.getValueOrDefault(key)
	// NOTE: JSON stores numbers as floats
	var res, err = cast.ToIntE(val)
	if err != nil {
		panic(fmt.Sprintf("%v is not an int number! It is %v", val, reflect.TypeOf(val)))
	}
	return res
}

func (this *Config) GetFloat64(key string) float64 {
	var val = this.getValueOrDefault(key)
	var res, err = cast.ToFloat64E(val)
	if err != nil {
		panic(fmt.Sprintf("%v is not a float number! It is %v", val, reflect.TypeOf(val)))
	}
	return res
}

