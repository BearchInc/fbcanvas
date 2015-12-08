package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"appengine"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func paypalWebhook(w http.ResponseWriter, r*http.Request) {
	c := appengine.NewContext(r)


	decoder := json.NewDecoder(r.Body)
	var objmap map[string]*json.RawMessage
	err := decoder.Decode(&objmap)
	if err != nil {
		c.Infof("Error: %+v", err)
		w.Write([]byte("Error!"))
		return
	}

	c.Infof("Map keys: %+v", objmap)

	var resource json.RawMessage
	err = json.Unmarshal(*objmap["resource"], &resource)

	c.Infof("resources: %+v", string(resource))
	w.Write([]byte("Success"))
}

func init() {
	http.HandleFunc("/webhook", paypalWebhook)
	http.HandleFunc("/", handler)
}
