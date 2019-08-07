package nested

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/icrowley/fake"
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
		"extras": "{\"lorem_ipsum\":{\"url\":\"www.lorem-ipsum.com\",\"id\":12},\"lorem_bacon\":{\"url\":\"www.lorem_bacon.com\",\"id\":\"da9883jw32dl12j120un9sa87ds5asn\"}}",
		"extras_error": "{\"lorem_ipsu:\"www.lorem_bacon.com\",\"id\":\"da9883jw32dl12j120un9sa87ds5asn\"}}",
	},
}

func TestNew(t *testing.T) {
	out := New(nil)
	if reflect.ValueOf(out).Type() != reflect.ValueOf(Map{}).Type() {
		t.Errorf("Expected an %T, but got %v", Map{}, out)
	}

	out = New(map[string]interface{}{"name": "Rodrigo Lopes"})
	if reflect.ValueOf(out).Type() != reflect.ValueOf(Map{}).Type() {
		t.Errorf("Expected an %T, but got %v", Map{}, out)
	}

}

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
func ExampleInterface() {
	data := map[string]interface{}{
		"person": map[string]interface{}{
			"name":  "Rodrigo",
			"level": 3,
		},
		"session": map[string]interface{}{
			"token": "62vsy29v8y4v248v5y97v1e21v35ce97",
		},
	}

	session, found := Interface("session", data)
	fmt.Println(session, found)
	// output: map[token:62vsy29v8y4v248v5y97v1e21v35ce97] true
}
func ExampleMap_Interface() {
	data := map[string]interface{}{
		"person": map[string]interface{}{
			"name":  "Rodrigo",
			"level": 3,
		},
		"session": map[string]interface{}{
			"token": "62vsy29v8y4v248v5y97v1e21v35ce97",
		},
	}

	session, found := New(data).Interface("session")
	fmt.Println(session, found)
	// output: map[token:62vsy29v8y4v248v5y97v1e21v35ce97] true
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

func TestGetInterface(t *testing.T) {
	tests := []struct {
		Parameter     string
		ExpectedFirst interface{}
		Data          map[string]interface{}
	}{
		{
			Parameter:     "",
			ExpectedFirst: nil,
			Data:          data,
		},
		{
			Parameter:     "advert.id",
			ExpectedFirst: "12",
			Data:          data,
		},
		{
			Parameter:     "advert.status.code",
			ExpectedFirst: "active",
			Data:          data,
		},
		{
			Parameter:     "advert.status.ttl",
			ExpectedFirst: 123123,
			Data:          data,
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
			Data: data,
		},
		{
			Parameter:     "advert.bananas",
			ExpectedFirst: nil,
			Data:          data,
		},
		{
			Parameter:     "",
			ExpectedFirst: nil,
			Data:          map[string]interface{}{},
		},
	}

	for key, test := range tests {
		t.Run(fmt.Sprintf("Test #%d", key), func(t *testing.T) {
			actual := GetInterface(test.Parameter, test.Data)
			if !reflect.DeepEqual(test.ExpectedFirst, actual) {
				t.Errorf("[%d] expected param1: %T(%v), but got param1: %T(%v)",
					key,
					test.ExpectedFirst, test.ExpectedFirst, // param1
					actual, actual,                         // actual
				)
			}
		})
	}
}
func ExampleGetInterface() {
	data := map[string]interface{}{
		"person": map[string]interface{}{
			"name":  "Rodrigo",
			"level": 3,
		},
		"session": map[string]interface{}{
			"token": "62vsy29v8y4v248v5y97v1e21v35ce97",
		},
	}

	session := GetInterface("session", data)
	fmt.Println(session)
	// output: map[token:62vsy29v8y4v248v5y97v1e21v35ce97]
}
func ExampleMap_GetInterface() {
	data := map[string]interface{}{
		"person": map[string]interface{}{
			"name":  "Rodrigo",
			"level": 3,
		},
		"session": map[string]interface{}{
			"token": "62vsy29v8y4v248v5y97v1e21v35ce97",
		},
	}

	session := New(data).GetInterface("session")
	fmt.Println(session)
	// output: map[token:62vsy29v8y4v248v5y97v1e21v35ce97]
}
func BenchmarkGetInterface(b *testing.B) {
	total := 10

	bench := make([]map[string]interface{}, total)
	for i := 0; i < total; i++ {
		bench[i] = randomData()
	}

	for n := 0; n < total; n++ {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GetInterface("advert.status.ttl", bench[n])
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
func ExampleInt() {
	data := map[string]interface{}{
		"person": map[string]interface{}{
			"name":  "Rodrigo",
			"level": 3,
		},
	}

	level, found := Int("person.level", data)
	fmt.Println(level, found)
	// output: 3 true
}
func ExampleMap_Int() {
	data := map[string]interface{}{
		"person": map[string]interface{}{
			"name":  "Rodrigo",
			"level": 3,
		},
	}

	level, found := New(data).Int("person.level")
	fmt.Println(level, found)
	// output: 3 true
}
func BenchmarkInt(b *testing.B) {
	total := 10

	bench := make([]map[string]interface{}, total)
	for i := 0; i < total; i++ {
		bench[i] = randomData()
	}

	for n := 0; n < total; n++ {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int("advert.status.ttl", bench[n])
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	tests := []struct {
		Parameter     string
		ExpectedFirst int
		Data          map[string]interface{}
	}{
		{
			Parameter:     "advert.status.ttl",
			ExpectedFirst: 123123,
			Data:          data,
		},
		{
			Parameter:     "advert.id",
			ExpectedFirst: 0,
			Data:          data,
		},
		{
			Parameter:     "advert.title.id",
			ExpectedFirst: 0,
			Data:          data,
		},
		{
			Parameter:     "advert.bananas",
			ExpectedFirst: 0,
			Data:          data,
		},
	}

	for key, test := range tests {
		t.Run(fmt.Sprintf("Test #%d", key), func(t *testing.T) {
			actual := GetInt(test.Parameter, test.Data)
			if actual != test.ExpectedFirst {
				t.Errorf("[%d] expected param1: %T(%v), but got param1: %T(%v)",
					key,
					test.ExpectedFirst, test.ExpectedFirst, // param1
					actual, actual,                         // actual
				)
			}
		})
	}
}
func ExampleGetInt() {
	data := map[string]interface{}{
		"person": map[string]interface{}{
			"name":  "Rodrigo",
			"level": 3,
		},
	}

	level := GetInt("person.level", data)
	fmt.Println(level)
	// output: 3
}
func ExampleMap_GetInt() {
	data := map[string]interface{}{
		"person": map[string]interface{}{
			"name":  "Rodrigo",
			"level": 3,
		},
	}

	level := New(data).GetInt("person.level")
	fmt.Println(level)
	// output: 3
}
func BenchmarkGetInt(b *testing.B) {
	total := 10

	bench := make([]map[string]interface{}, total)
	for i := 0; i < total; i++ {
		bench[i] = randomData()
	}

	for n := 0; n < total; n++ {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GetInt("advert.status.ttl", bench[n])
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
func ExampleString() {
	data := map[string]interface{}{
		"person": map[string]interface{}{
			"name": "Rodrigo",
		},
	}

	name, found := String("person.name", data)
	fmt.Println(name, found)
	// output: Rodrigo true
}
func ExampleMap_String() {
	data := map[string]interface{}{
		"person": map[string]interface{}{
			"name": "Rodrigo",
		},
	}

	name, found := New(data).String("person.name")
	fmt.Println(name, found)
	// output: Rodrigo true
}
func BenchmarkString(b *testing.B) {
	total := 10

	bench := make([]map[string]interface{}, total)
	for i := 0; i < total; i++ {
		bench[i] = randomData()
	}

	for n := 0; n < total; n++ {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				String("advert.title", bench[n])
			}
		})
	}
}

func TestGetString(t *testing.T) {

	tests := []struct {
		Parameter     string
		ExpectedFirst string
		Data          map[string]interface{}
	}{
		{
			Parameter:     "advert.title",
			ExpectedFirst: "Lorem Ipsum",
			Data:          data,
		},
		{
			Parameter:     "advert.status.ttl",
			ExpectedFirst: "",
			Data:          data,
		},
		{
			Parameter:     "advert.bananas",
			ExpectedFirst: "",
			Data:          data,
		},
	}

	for key, test := range tests {
		t.Run(fmt.Sprintf("Test #%d", key), func(t *testing.T) {
			actual := GetString(test.Parameter, test.Data)
			if actual != test.ExpectedFirst {
				t.Errorf("[%s] expected param1: %T(%v), but got param1: %T(%v)",
					test.Parameter,
					test.ExpectedFirst, test.ExpectedFirst, // param1
					actual, actual,                         // actual
				)
			}
		})
	}
}
func ExampleGetString() {
	data := map[string]interface{}{
		"person": map[string]interface{}{
			"name": "Rodrigo",
		},
	}

	name := GetString("person.name", data)
	fmt.Println(name)
	// output: Rodrigo
}
func ExampleMap_GetString() {
	data := map[string]interface{}{
		"person": map[string]interface{}{
			"name": "Rodrigo",
		},
	}

	name := New(data).GetString("person.name")
	fmt.Println(name)
	// output: Rodrigo
}
func BenchmarkGetString(b *testing.B) {
	total := 10

	bench := make([]map[string]interface{}, total)
	for i := 0; i < total; i++ {
		bench[i] = randomData()
	}

	for n := 0; n < total; n++ {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GetString("advert.title", bench[n])
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
func ExampleTime() {
	data := map[string]interface{}{
		"session": map[string]interface{}{
			"expire": "2018-08-08T18:00:00Z",
		},
	}

	expire, found := Time("session.expire", data, time.RFC3339)
	fmt.Println(expire, found)
	// output: 2018-08-08 18:00:00 +0000 UTC true
}
func ExampleMap_Time() {
	data := map[string]interface{}{
		"session": map[string]interface{}{
			"expire": "2018-08-08T18:00:00Z",
		},
	}

	expire, found := New(data).Time("session.expire", time.RFC3339)
	fmt.Println(expire, found)
	// output: 2018-08-08 18:00:00 +0000 UTC true
}
func BenchmarkTime(b *testing.B) {
	total := 10

	bench := make([]map[string]interface{}, total)
	for i := 0; i < total; i++ {
		bench[i] = randomData()
	}

	for n := 0; n < total; n++ {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Time("advert.timer.date_time", bench[n], time.RFC3339)
			}
		})
	}
}

func TestGetTime(t *testing.T) {
	tests := []struct {
		Layout        string
		Parameter     string
		ExpectedFirst int64
		Data          map[string]interface{}
	}{
		{
			Layout:        "02/01/2006",
			Parameter:     "advert.timer.birth",
			ExpectedFirst: 538876800000000000,
			Data:          data,
		},
		{
			Layout:        time.RFC3339,
			Parameter:     "advert.timer.date_time",
			ExpectedFirst: 538945200000000000, // "1987-01-29T19:00:00Z00:00",
			Data:          data,
		},
		{
			Layout:        "",
			Parameter:     "advert.timer.date_time",
			ExpectedFirst: 538945200000000000, // "1987-01-29T19:00:00Z00:00",
			Data:          data,
		},
		{
			Layout:        time.ANSIC,
			Parameter:     "advert.timer.date_time",
			ExpectedFirst: -6795364578871345152, // "1987-01-29T19:00:00Z00:00",
			Data:          data,
		},
		{
			Parameter:     "advert.owner.id",
			ExpectedFirst: -6795364578871345152,
			Data:          data,
		},
		{
			Parameter:     "advert.bananas",
			ExpectedFirst: -6795364578871345152,
			Data:          data,
		},
	}

	for key, test := range tests {
		t.Run(fmt.Sprintf("Test #%d", key), func(t *testing.T) {
			actual := GetTime(test.Parameter, test.Data, test.Layout)
			if actual.UnixNano() != test.ExpectedFirst {
				t.Errorf("[%s] expected param1: %T(%v), but got param1: %T(%v)",
					test.Parameter,
					test.ExpectedFirst, test.ExpectedFirst, // param1
					actual, actual.UnixNano(),              // actual
				)
			}
		})
	}
}
func ExampleGetTime() {
	data := map[string]interface{}{
		"session": map[string]interface{}{
			"expire": "2018-08-08T18:00:00Z",
		},
	}

	expire := GetTime("session.expire", data, time.RFC3339)
	fmt.Println(expire)
	// output: 2018-08-08 18:00:00 +0000 UTC
}
func ExampleMap_GetTime() {
	data := map[string]interface{}{
		"session": map[string]interface{}{
			"expire": "2018-08-08T18:00:00Z",
		},
	}

	expire := New(data).GetTime("session.expire", time.RFC3339)
	fmt.Println(expire)
	// output: 2018-08-08 18:00:00 +0000 UTC
}
func BenchmarkGetTime(b *testing.B) {
	total := 10

	bench := make([]map[string]interface{}, total)
	for i := 0; i < total; i++ {
		bench[i] = randomData()
	}

	for n := 0; n < total; n++ {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GetTime("advert.timer.date_time", bench[n], time.RFC3339)
			}
		})
	}
}

func TestSubFromString(t *testing.T) {
	getType := func(t interface{}) string {
		if t == nil {
			return "nil"
		}
		return reflect.TypeOf(t).Name()
	}

	tests := []struct {
		ParameterFirst string
		ExpectedFirst  interface{}
		ExpectedSecond bool
		Data           map[string]interface{}
	}{
		{
			ParameterFirst: "advert.extras",
			ExpectedFirst:  New(Map{}),
			ExpectedSecond: true,
			Data:           data,
		},
		{
			ParameterFirst: "extras.ttl",
			ExpectedFirst:  Map{},
			ExpectedSecond: false,
			Data:           data,
		},
		{
			ParameterFirst: "advert.extras_error",
			ExpectedFirst:  New(Map{}),
			ExpectedSecond: false,
			Data:           data,
		},
	}

	for key, test := range tests {
		t.Run(fmt.Sprintf("Test #%d", key), func(t *testing.T) {
			actual, result := SubFromString(test.ParameterFirst, test.Data)
			if getType(actual) != getType(test.ExpectedFirst) || result != test.ExpectedSecond {
				t.Errorf("[%s] expected param1: %T(%v) and param2: %T(%v), but got param1: %T(%v) and param2: %T(%v)",
					test.ParameterFirst,
					test.ExpectedFirst, test.ExpectedFirst,   // param1
					test.ExpectedSecond, test.ExpectedSecond, // param2
					actual, getType(actual),                  // actual
					result, result,                           // result
				)
			}
		})
	}
}
func ExampleSubFromString() {
	data := map[string]interface{}{
		"extras": "{\"lorem_ipsum\":{\"url\":\"www.lorem-ipsum.com\",\"id\":12},\"lorem_bacon\":{\"url\":\"www.lorem_bacon.com\",\"id\":\"da9883jw32dl12j120un9sa87ds5asn\"}}",
	}

	name, found := SubFromString("extras", data)
	fmt.Println(name, found)
	// output: map[lorem_bacon:map[id:da9883jw32dl12j120un9sa87ds5asn url:www.lorem_bacon.com] lorem_ipsum:map[id:12 url:www.lorem-ipsum.com]] true
}
func ExampleMap_SubFromString() {
	data := map[string]interface{}{
		"extras": "{\"lorem_ipsum\":{\"url\":\"www.lorem-ipsum.com\",\"id\":12},\"lorem_bacon\":{\"url\":\"www.lorem_bacon.com\",\"id\":\"da9883jw32dl12j120un9sa87ds5asn\"}}",
	}

	name, found := New(data).SubFromString("extras")
	fmt.Println(name, found)
	// output: map[lorem_bacon:map[id:da9883jw32dl12j120un9sa87ds5asn url:www.lorem_bacon.com] lorem_ipsum:map[id:12 url:www.lorem-ipsum.com]] true
}
func BenchmarkSubFromString(b *testing.B) {
	total := 10

	bench := make([]map[string]interface{}, total)
	for i := 0; i < total; i++ {
		bench[i] = randomData()
	}

	for n := 0; n < total; n++ {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SubFromString("advert.title", bench[n])
			}
		})
	}
}

func TestGetSubFromString(t *testing.T) {
	getType := func(t interface{}) string {
		if t == nil {
			return "nil"
		}
		return reflect.TypeOf(t).Name()
	}

	tests := []struct {
		ParameterFirst string
		ExpectedFirst  interface{}
		Data           map[string]interface{}
	}{
		{
			ParameterFirst: "advert.extras",
			ExpectedFirst:  New(Map{}),
			Data:           data,
		},
		{
			ParameterFirst: "extras.ttl",
			ExpectedFirst:  Map{},
			Data:           data,
		},
		{
			ParameterFirst: "advert.extras_error",
			ExpectedFirst:  New(Map{}),
			Data:           data,
		},
	}

	for key, test := range tests {
		t.Run(fmt.Sprintf("Test #%d", key), func(t *testing.T) {
			actual := GetSubFromString(test.ParameterFirst, test.Data)
			if getType(actual) != getType(test.ExpectedFirst){
				t.Errorf("[%s] expected param1: %T(%v), but got param1: %T(%v)",
					test.ParameterFirst,
					test.ExpectedFirst, test.ExpectedFirst,   // param1
					actual, actual,                           // result
				)
			}
		})
	}
}
func ExampleGetSubFromString() {
	data := map[string]interface{}{
		"extras": "{\"lorem_ipsum\":{\"url\":\"www.lorem-ipsum.com\",\"id\":12},\"lorem_bacon\":{\"url\":\"www.lorem_bacon.com\",\"id\":\"da9883jw32dl12j120un9sa87ds5asn\"}}",
	}

	name := GetSubFromString("extras", data)
	fmt.Println(name)
	// output: map[lorem_bacon:map[id:da9883jw32dl12j120un9sa87ds5asn url:www.lorem_bacon.com] lorem_ipsum:map[id:12 url:www.lorem-ipsum.com]]
}
func ExampleMap_GetSubFromString() {
	data := map[string]interface{}{
		"extras": "{\"lorem_ipsum\":{\"url\":\"www.lorem-ipsum.com\",\"id\":12},\"lorem_bacon\":{\"url\":\"www.lorem_bacon.com\",\"id\":\"da9883jw32dl12j120un9sa87ds5asn\"}}",
	}

	name := New(data).GetSubFromString("extras")
	fmt.Println(name)
	// output: map[lorem_bacon:map[id:da9883jw32dl12j120un9sa87ds5asn url:www.lorem_bacon.com] lorem_ipsum:map[id:12 url:www.lorem-ipsum.com]]
}
func BenchmarkGetSubFromString(b *testing.B) {
	total := 10

	bench := make([]map[string]interface{}, total)
	for i := 0; i < total; i++ {
		bench[i] = randomData()
	}

	for n := 0; n < total; n++ {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GetSubFromString("advert.title", bench[n])
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
