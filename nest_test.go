package nested

import (
	"testing"
	"fmt"
	"reflect"

	"github.com/icrowley/fake"
	"time"
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
		"timer": map[string]interface{}{
			"date_time": "1987-01-29T19:00:00Z",
			"birth":     "29/01/1987",
		},
	},
}

// Test for Nested function
func TestInterface(t *testing.T) {
	tests := []struct {
		Parameter      string
		ExpectedFirst  interface{}
		ExpectedSecond bool
		Data           map[string]interface{}
	}{
		{
			Parameter:      "",
			ExpectedFirst:  nil,
			ExpectedSecond: false,
			Data:           data,
		},
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
		ExpectedFirst  int
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

		{
			Parameter:      "advert.bananas",
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
	tests := []struct {
		Parameter      string
		ExpectedFirst  string
		ExpectedSecond bool
		Data           map[string]interface{}
	}{
		{
			Parameter:      "advert.title",
			ExpectedFirst:  "Lorem Ipsum",
			ExpectedSecond: true,
			Data:           data,
		},
		{
			Parameter:      "advert.status.ttl",
			ExpectedFirst:  "",
			ExpectedSecond: false,
			Data:           data,
		},
		{
			Parameter:      "advert.bananas",
			ExpectedFirst:  "",
			ExpectedSecond: false,
			Data:           data,
		},
	}

	for key, test := range tests {
		t.Run(fmt.Sprintf("Test #%d", key), func(t *testing.T) {
			actual, result := String(test.Parameter, test.Data)
			if actual != test.ExpectedFirst || result != test.ExpectedSecond {
				t.Errorf("[%s] expected param1: %T(%v) and param2: %T(%v), but got param1: %T(%v) and param2: %T(%v)",
					test.Parameter,
					test.ExpectedFirst, test.ExpectedFirst,   // param1
					test.ExpectedSecond, test.ExpectedSecond, // param2
					actual, actual,                           // actual
					result, result,                           // result
				)
			}
		})
	}
}

func TestTime(t *testing.T) {
	tests := []struct {
		Layout         string
		Parameter      string
		ExpectedFirst  int64
		ExpectedSecond bool
		Data           map[string]interface{}
	}{
		{
			Layout:         "02/01/2006",
			Parameter:      "advert.timer.birth",
			ExpectedFirst:  538876800000000000,
			ExpectedSecond: true,
			Data:           data,
		},
		{
			Layout:         time.RFC3339,
			Parameter:      "advert.timer.date_time",
			ExpectedFirst:  538945200000000000, // "1987-01-29T19:00:00Z00:00",
			ExpectedSecond: true,
			Data:           data,
		},
		{
			Layout:         "",
			Parameter:      "advert.timer.date_time",
			ExpectedFirst:  538945200000000000, // "1987-01-29T19:00:00Z00:00",
			ExpectedSecond: true,
			Data:           data,
		},
		{
			Layout:         time.ANSIC,
			Parameter:      "advert.timer.date_time",
			ExpectedFirst:  -6795364578871345152, // "1987-01-29T19:00:00Z00:00",
			ExpectedSecond: false,
			Data:           data,
		},
		{
			Parameter:      "advert.owner.id",
			ExpectedFirst:  -6795364578871345152,
			ExpectedSecond: false,
			Data:           data,
		},
		{
			Parameter:      "advert.bananas",
			ExpectedFirst:  -6795364578871345152,
			ExpectedSecond: false,
			Data:           data,
		},
	}

	for key, test := range tests {
		t.Run(fmt.Sprintf("Test #%d", key), func(t *testing.T) {
			actual, result := Time(test.Parameter, test.Data, test.Layout)
			if actual.UnixNano() != test.ExpectedFirst || result != test.ExpectedSecond {
				t.Errorf("[%s] expected param1: %T(%v) and param2: %T(%v), but got param1: %T(%v) and param2: %T(%v)",
					test.Parameter,
					test.ExpectedFirst, test.ExpectedFirst,   // param1
					test.ExpectedSecond, test.ExpectedSecond, // param2
					actual, actual.UnixNano(),                // actual
					result, result,                           // result
				)
			}
		})
	}
}

func randomData() map[string]interface{} {
	return map[string]interface{}{
		"advert": map[string]interface{}{
			"id":    fake.Digits(),
			"title": fake.Title(),
			"status": map[string]interface{}{
				"code": fake.Color(),
				"url":  "www.loremipsum.com",
				"ttl":  int(time.Now().UnixNano()),
			},
			"contact": map[string]interface{}{
				"name": "daniel3",
				"phones": []string{
					fake.Phone(),
					fake.Phone(),
				},
			},
			"timer": map[string]interface{}{
				"date_time": "1987-01-29T19:00:00Z00:00",
				"birth":     "29/01/1987",
			},
		},
	}
}

func BenchmarkInterface(b *testing.B) {
	total := 10

	bench := make([]map[string]interface{}, total)
	for i := 0; i < total; i++ {
		bench[i] = randomData()
	}

	for n := 0; n < total; n++ {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Interface("advert.status.ttl", bench[n])
			}
		})
	}
}