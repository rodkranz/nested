package nested

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

// ErrInvalidInputType when input is invalid or cannot be casted.
var ErrInvalidInputType = errors.New("this is not a valid input")

// Map type with functions bind
type Map map[string]interface{}

// New returns new Map instance or do a cast type for Map
func New(in map[string]interface{}) Map {
	if in == nil {
		in = make(map[string]interface{})
	}

	return Map(in)
}

// NewFromJSON returns new Map instance when in is a json valid
func NewFromJSON(in string) (Map, error) {
	var m Map
	if err := json.Unmarshal([]byte(in), &m); err != nil {
		return nil, ErrInvalidInputType
	}
	return m, nil
}

// NewFromInterface return new map instance if can cast input to map[string]interface{}.
func NewFromInterface(in interface{}) (Map, error) {
	if m, ok := in.(map[string]interface{}); ok {
		return New(m), nil
	}

	return nil, ErrInvalidInputType
}

// GetInterface returns the interface value from position that you passed by argument
func (m Map) GetInterface(position string) interface{} {
	value, _ := m.Interface(position)
	return value
}

// Interface returns the value from position that you pass separately by . (dot)
// the first value is a value that you are looking for and second is bool if found the field or not
// if the field is not found it returns nil and false.
func (m Map) Interface(position string) (interface{}, bool) {
	pos := strings.Split(position, ".")

	t := m
	for key, posKey := range pos {
		v, ok := t[posKey]
		if !ok {
			break
		}

		if newValue, ok := v.(map[string]interface{}); ok {
			if key+1 == len(pos) {
				return v, true
			}

			t = newValue
		} else if newString, ok := v.(string); ok {
			return newString, true
		} else if newInt, ok := v.(int); ok {
			return newInt, true
		}
	}

	return nil, false
}

// GetString returns the string value from position that you passed by argument
func (m Map) GetString(position string) string {
	value, _ := m.String(position)
	return value
}

// String returns the string value from position that you passed by argument and a bool if found the field.
// if it doesn't find the field the returns is "" and false.
func (m Map) String(position string) (value string, ok bool) {
	var valueTmp interface{}
	if valueTmp, ok = m.Interface(position); !ok {
		return "", false
	}
	if value, ok = valueTmp.(string); !ok {
		return "", false
	}
	return value, true
}

// GetInt returns the int value from position that you passed by argument
func (m Map) GetInt(position string) int {
	value, _ := m.Int(position)
	return value
}

// Int returns the int value from position that you passed by argument and a bool if found the field.
// if it doesn't find the field the returns is 0 and false.
func (m Map) Int(position string) (value int, ok bool) {
	var valueTmp interface{}
	if valueTmp, ok = m.Interface(position); !ok {
		return 0, false
	}
	if value, ok = valueTmp.(int); !ok {
		return 0, false
	}
	return value, true
}

// GetTime returns the time value from position that you passed by argument
func (m Map) GetTime(position, layout string) time.Time {
	value, _ := m.Time(position, layout)
	return value
}

// Time returns the time.Time value from position that you passed by argument and a bool if found the field.
// if it doesn't find the field the returns is time.Time default and false.
// By default the layout is time.RFC3339, you can change the layout using a new one as second parameter
func (m Map) Time(position, layout string) (value time.Time, ok bool) {
	var valueTmp string
	if valueTmp, ok = m.String(position); !ok {
		return time.Time{}, false
	}

	if layout == "" {
		layout = time.RFC3339
	}

	var err error
	if value, err = time.Parse(layout, valueTmp); err != nil {
		return time.Time{}, false
	}

	return value, true
}

// SubFromString return Map from string json format if json is valid.
func (m Map) SubFromString(position string) (Map, bool) {
	subData, ok := m.String(position)
	if !ok {
		return nil, false
	}

	var value Map
	if err := json.Unmarshal([]byte(subData), &value); err != nil {
		return nil, false
	}

	return value, true
}

// GetSubFromString returns the time value from position that you passed by argument
func (m Map) GetSubFromString(position string) Map {
	value, _ := m.SubFromString(position)
	return value
}

// Interface is helper for function Interface from Map.
func Interface(position string, mapper map[string]interface{}) (interface{}, bool) {
	return New(mapper).Interface(position)
}

// GetInterface is helper for function GetInterface from Map.
func GetInterface(position string, mapper map[string]interface{}) interface{} {
	return New(mapper).GetInterface(position)
}

// String is helper for function String from Map.
func String(position string, mapper map[string]interface{}) (string, bool) {
	return New(mapper).String(position)
}

// GetString is helper for function GetString from Map.
func GetString(position string, mapper map[string]interface{}) string {
	return New(mapper).GetString(position)
}

// Int is helper for function Int from Map.
func Int(position string, mapper map[string]interface{}) (int, bool) {
	return New(mapper).Int(position)
}

// GetInt is helper for function GetInt from Map.
func GetInt(position string, mapper map[string]interface{}) int {
	return New(mapper).GetInt(position)
}

// Time is helper for function Time from Map.
func Time(position string, mapper map[string]interface{}, layout string) (time.Time, bool) {
	return New(mapper).Time(position, layout)
}

// GetTime is helper for function GetTime from Map.
func GetTime(position string, mapper map[string]interface{}, layout string) time.Time {
	return New(mapper).GetTime(position, layout)
}

// SubFromString is helper for function SubFromString from Map.
func SubFromString(position string, mapper map[string]interface{}) (Map, bool) {
	return New(mapper).SubFromString(position)
}

// GetSubFromString is helper for function GetSubFromString from Map.
func GetSubFromString(position string, mapper map[string]interface{}) Map {
	return New(mapper).GetSubFromString(position)
}
