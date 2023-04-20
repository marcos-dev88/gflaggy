package gflaggy

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
)

var CLIParams = os.Args

type Flag interface {
	Bool() (bool, error)
	String() (string, error)
	Int() (int, error)
	Float32() (float32, error)
	Float64() (float64, error)
	JSON() (map[string]interface{}, error)
}

func (f *flag) Bool() (bool, error) {
	f.Type = Boolean
	return getValue(f, false)
}

func (f *flag) String() (string, error) {
	f.Type = String
	return getValue(f, "")
}

func (f *flag) Int() (int, error) {
	f.Type = Integer
	return getValue(f, 0)
}

func (f *flag) Float32() (float32, error) {
	f.Type = Float32
	return getValue(f, float32(0.0))
}

func (f *flag) Float64() (float64, error) {
	f.Type = Float64
	return getValue(f, 0.0)
}

func (f *flag) JSON() (map[string]interface{}, error) {
	f.Type = JSON
	return getValue(f, map[string]interface{}{})
}

func getValue[T any](f *flag, typeReturn T) (data T, err error) {
	val, err := f.getParam(CLIParams)

	if err != nil {
		return
	}

	if len(val) == 0 && f.Required {
		err = errors.New("flag " + f.Name + " is required")
		return
	}

	var returnData any

	switch f.Type {
	case Boolean:
		if len(val) > 0 {
			returnData = true
			return returnData.(T), nil
		}
		return
	case String:
		returnData = val
		return returnData.(T), nil
	case Integer:
		if len(val) == 0 {
			val = "0"
		}
		convVal, errParse := strconv.Atoi(val)
		if errParse != nil {
			err = errParse
			return
		}
		returnData = convVal
		return returnData.(T), nil
	case Float64:
		if len(val) == 0 {
			val = "0.0"
		}
		convVal, errParse := strconv.ParseFloat(val, 64)
		if errParse != nil {
			err = errParse
			return
		}
		returnData = convVal
		return returnData.(T), nil
	case Float32:
		if len(val) == 0 {
			val = "0.0"
		}
		convVal, errParse := strconv.ParseFloat(val, 32)
		if errParse != nil {
			err = errParse
			return
		}
		returnData = float32(convVal)
		return returnData.(T), nil

	case JSON:
		if len(val) == 0 {
			return
		}
		if err := json.Unmarshal([]byte(val), &data); err != nil {
			return returnData.(T), err
		}
		return data, nil
	}
	return
}

func (f flag) getParam(params []string) (string, error) {
	if len(params) < 2 {
		return "", errors.New("was expected a param to get a flag and nothing came")
	}

	mapIndex := flagIndex(f.Name, params)

	if v, ok := mapIndex[f.Name]; ok {
		if v != DefaultNullValue {
			if f.Type == Boolean {
				return params[v], nil
			}
			return params[v+1], nil
		}
	}

	return "", nil
}

func flagIndex(flagName string, params []string) map[string]int {
	returnIndex := make(map[string]int, 0)
	returnIndex[flagName] = DefaultNullValue
	for i := 0; i < len(params); i++ {
		if params[i] == flagName {
			returnIndex[params[i]] = i
		}
	}

	return returnIndex
}
