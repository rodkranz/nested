package nested

import (
	"strings"
	"time"
)

// Map type with functions bind
type Map map[string]interface{}

// New returns new Map instance or do a cast type for Map
func New(in map[string]interface{}) Map {
	if in == nil {
		in = make(map[string]interface{})
	}

	return Map(in)
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

// String returns the string value from position that you passed by arguments and a bool if found the field.
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

// Int returns the int value from position that you passed by arguments and a bool if found the field.
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

// Time returns the time.Time value from position that you passed by arguments and a bool if found the field.
// if it doesn't find the field the returns is time.Time default and false.
// By default the layout is time.RFC3339, you can change the layout using a new one as second parameter
func (m Map) Time(position string, layout string) (value time.Time, ok bool) {
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

// Interface is helper for function Interface from Map.
func Interface(position string, mapper map[string]interface{}) (interface{}, bool) {
	return New(mapper).Interface(position)
}

// String is helper for function String from Map.
func String(position string, mapper map[string]interface{}) (value string, ok bool) {
	return New(mapper).String(position)
}

// Int is helper for function Int from Map.
func Int(position string, mapper map[string]interface{}) (value int, ok bool) {
	return New(mapper).Int(position)
}

// Time is helper for function Time from Map.
func Time(position string, mapper map[string]interface{}, layout string) (value time.Time, ok bool) {
	return New(mapper).Time(position, layout)
}
