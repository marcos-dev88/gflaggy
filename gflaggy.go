package gflaggy

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"strconv"
)

var CLIParams = os.Args

type Flag interface {
	String() (string, error)
	Int() (int, error)
	Float() (float64, error)
	JSON() (map[string]interface{}, error)
}

func (f *flag) String() (string, error) {
	return getValue(f.Required, f.Name, "")
}

func (f *flag) Int() (int, error) {
	return getValue(f.Required, f.Name, 0)
}

func (f *flag) Float() (float64, error) {
	return getValue(f.Required, f.Name, 0.0)
}

func (f *flag) JSON() (map[string]interface{}, error) {
	return getValue(f.Required, f.Name, map[string]interface{}{})
}

func getValue[T any](isRequired bool, flagName string, typeReturn T) (data T, err error) {
	typeDesiredReturn := reflect.TypeOf(typeReturn).String()

	val, err := getParam(flagName, CLIParams)

	if err != nil {
		return
	}

	if len(val) == 0 && isRequired {
		err = errors.New("flag " + flagName + " is required")
		return
	}

	var returnData any

	switch typeDesiredReturn {
	case "string":
		returnData = val
		return returnData.(T), nil
	case "int":
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
	case "float64":
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
	case "map[string]interface {}":
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

func getParam(flagName string, params []string) (string, error) {
	if len(params) < 2 {
		return "", errors.New("was expected a param to get a flag and nothing came")
	}

	mapIndex := flagIndex(flagName, params)

	if v, ok := mapIndex[flagName]; ok {
		if v != DefaultNullValue {
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
