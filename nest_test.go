package nested

import (
	"testing"
	"fmt"
	"reflect"
)

var data = map[string]interface{}{
	"advert": map[string]interface{}{
		"id":    "12",
		"title": "Lorem Ipsum",
		"status": map[string]interface{}{
			"code": "active",
			"url":  "www.loremipsum.com",
			"ttl":  123123,
		},
		"contact": map[string]interface{}{
			"name": "daniel3",
			"phones": []string{
				"790123123",
				"790123546",
			},
		},
	},
}

// Test for Nested function
func TestNested(t *testing.T) {

	tests := []struct {
		Parameter      string
		ExpectedFirst  interface{}
		ExpectedSecond bool
		Data           map[string]interface{}
	}{
		{
			Parameter:      "advert.id",
			ExpectedFirst:  "12",
			ExpectedSecond: true,
			Data:           data,
		},
		{
			Parameter:      "advert.status.code",
			ExpectedFirst:  "active",
			ExpectedSecond: true,
			Data:           data,
		},
		{
			Parameter:      "advert.status.ttl",
			ExpectedFirst:  123123,
			ExpectedSecond: true,
			Data:           data,
		},
		{
			Parameter: "advert.contact",
			ExpectedFirst: map[string]interface{}{
				"name": "daniel3",
				"phones": []string{
					"790123123",
					"790123546",
				},
			},
			ExpectedSecond: true,
			Data:           data,
		},
		{
			Parameter:      "advert.bananas",
			ExpectedFirst:  nil,
			ExpectedSecond: false,
			Data:           data,
		},
		{
			Parameter:      "",
			ExpectedFirst:  nil,
			ExpectedSecond: false,
			Data:           map[string]interface{}{},
		},
	}

	for key, test := range tests {
		t.Run(fmt.Sprintf("Test #%d", key), func(t *testing.T) {
			actual, result := Interface(test.Parameter, test.Data)
			if !reflect.DeepEqual(test.ExpectedFirst, actual) || result != test.ExpectedSecond {
				t.Errorf("[%d] expected param1: %T(%v) and param2: %T(%v), but got param1: %T(%v) and param2: %T(%v)",
					key,
					test.ExpectedFirst, test.ExpectedFirst,   // param1
					test.ExpectedSecond, test.ExpectedSecond, // param2
					actual, actual,                           // actual
					result, result,                           // result
				)
			}
		})
	}
}

func TestInt(t *testing.T) {

	tests := []struct {
		Parameter      string
		ExpectedFirst  interface{}
		ExpectedSecond bool
		Data           map[string]interface{}
	}{
		{
			Parameter:      "advert.status.ttl",
			ExpectedFirst:  123123,
			ExpectedSecond: true,
			Data:           data,
		},
		{
			Parameter:      "advert.id",
			ExpectedFirst:  0,
			ExpectedSecond: false,
			Data:           data,
		},
		{
			Parameter:      "advert.title.id",
			ExpectedFirst:  0,
			ExpectedSecond: false,
			Data:           data,
		},
	}

	for key, test := range tests {
		t.Run(fmt.Sprintf("Test #%d", key), func(t *testing.T) {
			actual, result := Int(test.Parameter, test.Data)
			if actual != test.ExpectedFirst || result != test.ExpectedSecond {
				t.Errorf("[%d] expected param1: %T(%v) and param2: %T(%v), but got param1: %T(%v) and param2: %T(%v)",
					key,
					test.ExpectedFirst, test.ExpectedFirst,   // param1
					test.ExpectedSecond, test.ExpectedSecond, // param2
					actual, actual,                           // actual
					result, result,                           // result
				)
			}
		})
	}
}

func TestString(t *testing.T) {
	data := map[string]interface{}{
		"advert": map[string]interface{}{
			"id":    "12",
			"title": "Lorem Ipsum",
			"status": map[string]interface{}{
				"code": "active",
				"url":  "www.loremipsum.com",
				"ttl":  123123,
			},
			"contact": map[string]interface{}{
				"name": "daniel3",
				"phones": []string{
					"790123123",
					"790123546",
				},
			},
		},
	}

	tests := []struct {
		Parameter      string
		ExpectedFirst  interface{}
		ExpectedSecond bool
		Data           map[string]interface{}
	}{
		{
			Parameter:      "advert.contact.ttl",
			ExpectedFirst:  123123,
			ExpectedSecond: true,
			Data:           data,
		},
		{
			Parameter:      "advert.id",
			ExpectedFirst:  0,
			ExpectedSecond: false,
			Data:           data,
		},
		{
			Parameter:      "advert.title.id",
			ExpectedFirst:  0,
			ExpectedSecond: false,
			Data:           data,
		},
	}

	for key, test := range tests {
		t.Run(fmt.Sprintf("Test #%d", key), func(t *testing.T) {
			actual, result := Int(test.Parameter, test.Data)
			if actual != test.ExpectedFirst || result != test.ExpectedSecond {
				t.Errorf("[%d] expected param1: %T(%v) and param2: %T(%v), but got param1: %T(%v) and param2: %T(%v)",
					key,
					test.ExpectedFirst, test.ExpectedFirst,   // param1
					test.ExpectedSecond, test.ExpectedSecond, // param2
					actual, actual,                           // actual
					result, result,                           // result
				)
			}
		})
	}
}
