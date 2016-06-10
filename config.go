package ya_go_config

import (
	"encoding/json"
	"os"
	"fmt"
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

func (this *Config) LoadE(jsonFilePath string) error {
	file, _ := os.Open(jsonFilePath)
	decoder := json.NewDecoder(file)
	return decoder.Decode(&this.values)
}

func (this *Config) Load(jsonFilePath string) {
	var err = this.LoadE(jsonFilePath)
	if err != nil {
		panic(fmt.Sprintf("Cannot load config file '%s', reason: %v", jsonFilePath, err))
	}
}

func (this *Config) getValueOrDefaultE(key string) (interface{}, error) {
	val, ok := this.values[key]
	if !ok {
		val, ok = this.defaults[key]
		if !ok {
			return nil, fmt.Errorf("Key '%s' is not found in config.json nor config defaults", key)
		}
	}
	return val, nil
}

func (this *Config) GetStrE(key string) (res string, err error) {
	val, err := this.getValueOrDefaultE(key)
	if err != nil {
		return res, err
	}
	res, ok := val.(string)
	if !ok {
		return res, fmt.Errorf("Type assertion '%v' to string failed", val)
	}
	return res, nil
}

func (this *Config) GetStr(key string) string {
	res, err := this.GetStrE(key)
	if err != nil {
		panic(err)
	}
	return res
}

func (this *Config) GetIntE(key string) (res int, err error) {
	val, err := this.GetFloat64E(key) // NOTE: JSON stores any number as float
	if err != nil {
		return res, fmt.Errorf("Could not get number value for key '%s', reason: %v", key, err)
	}
	res = int(val) // TODO: panic if fractional part is not zero
	return res, nil
}

func (this *Config) GetInt(key string) int {
	res, err := this.GetIntE(key)
	if err != nil {
		panic(err)
	}
	return res
}

func (this *Config) GetFloat64E(key string) (res float64, err error) {
	val, err := this.getValueOrDefaultE(key)
	if err != nil {
		return res, err
	}
	res, ok := val.(float64)
	if !ok {
		return res, fmt.Errorf("Type assertion '%v' to float failed", val)
	}
	return res, nil
}

func (this *Config) GetFloat64(key string) float64 {
	res, err := this.GetFloat64E(key)
	if err != nil {
		panic(err)
	}
	return res
}


func (this *Config) GetBoolE(key string) (res bool, err error) {
	val, err := this.getValueOrDefaultE(key)
	if err != nil {
		return res, err
	}
	res, ok := val.(bool)
	if !ok {
		fmt.Errorf("Type assertion '%v' to bool failed", val)
	}
	return res, nil
}

func (this *Config) GetBool(key string) bool {
	var res, err = this.GetBoolE(key)
	if err != nil {
		panic(err)
	}
	return res
}


