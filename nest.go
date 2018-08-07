package nested

import (
	"strings"
	"time"
)

func Interface(position string, mapper map[string]interface{}) (interface{}, bool) {
	pos := strings.Split(position, ".")

	t := mapper
	for key, posKey := range pos {
		v, ok := t[posKey]
		if !ok {
			return nil, false
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

func String(position string, mapper map[string]interface{}) (value string, ok bool) {
	var valueTmp interface{}
	if valueTmp, ok = Interface(position, mapper); !ok {
		return "", false
	}
	if value, ok = valueTmp.(string); !ok {
		return "", false
	}
	return value, true
}

func Int(position string, mapper map[string]interface{}) (value int, ok bool) {
	var valueTmp interface{}
	if valueTmp, ok = Interface(position, mapper); !ok {
		return 0, false
	}
	if value, ok = valueTmp.(int); !ok {
		return 0, false
	}
	return value, true
}

func Time(position string, mapper map[string]interface{}, layout string) (value time.Time, ok bool) {
	var valueTmp string
	if valueTmp, ok = String(position, mapper); !ok {
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
