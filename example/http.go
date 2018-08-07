package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"time"

	"github.com/rodkranz/nested"
)

/**
curl -X POST \
  'http://localhost:8080/?field=advert.contact.name' \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 313cc416-965e-4dc3-ab16-bc76373158be' \
  -d '{
  "advert": {
    "contact": {
      "name": "daniel3",
      "phones": [
        "473-68-42",
        "789-52-84"
      ]
    },
    "id": "91",
    "status": {
      "code": "Orange",
      "ttl": 1533657456135539726,
      "url": "www.loremipsum.com"
    },
    "timer": {
      "birth": "29/01/1987",
      "date_time": "1987-01-29T19:00:00Z00:00"
    },
    "title": "MOLLITIA SUSCIPIT"
  }
}'
 */

func main() {
	index := func(w http.ResponseWriter, r *http.Request) {
		fieldName := r.URL.Query().Get("field")
		if fieldName == "" {
			http.Error(w, fmt.Sprint("use filed in query to defined the path"), http.StatusBadRequest)
			return
		}

		var bodyRequested map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&bodyRequested); err != nil {
			http.Error(w, fmt.Sprintf("cannot parse request because: %s", err), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		output := struct {
			Field   string
			Data    interface{}
			Results map[string]interface{}
		}{
			Field:   fieldName,
			Results: make(map[string]interface{}),
			Data:    bodyRequested,
		}

		{
			result, found := nested.Interface(fieldName, bodyRequested)
			output.Results["Interface"] = map[string]interface{}{"result": result, "found": found}
		}
		{
			result, found := nested.String(fieldName, bodyRequested)
			output.Results["String"] = map[string]interface{}{"result": result, "found": found}
		}
		{
			result, found := nested.Int(fieldName, bodyRequested)
			output.Results["Int"] = map[string]interface{}{"result": result, "found": found}
		}
		{
			result, found := nested.Time(fieldName, bodyRequested, time.RFC3339)
			output.Results["Time"] = map[string]interface{}{"result": result, "found": found}
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&output); err != nil {
			http.Error(w, fmt.Sprintf("cannot parse request because: %s", err), http.StatusBadRequest)
			return
		}
	}

	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
