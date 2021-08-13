package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func SID(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("Method not allowed"))
	} else {
		requestBody, _ := ioutil.ReadAll(r.Body)
		var body ShortPayload
		if err := json.Unmarshal(requestBody, &body); err != nil {
			w.Write([]byte(err.Error()))
		} else {
			if _, err := url.ParseRequestURI(body.Url); err != nil {
				w.Write([]byte("URL is required"))
			} else {
				result, err := createShortUrl(ShortPayload{
					Url: body.Url,
				})
				if err != nil {
					w.Write([]byte(err.Error()))
				} else {
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(result)
				}
			}
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
