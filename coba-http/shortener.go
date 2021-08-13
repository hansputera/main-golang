package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Typings
type ShortPayload struct {
	Url string
}

type ShortResultSID struct {
	Url         string
	OriginalUrl string
	CreatedAt   CreatedAtSID
}

type CreatedAtSID struct {
	Date         string `json:"date"`
	Timezone     string `json:"timezone"`
	TimezoneType int    `json:"timezone_type"`
}

type OriginalResponseSID struct {
	LongUrl   string       `json:"long_url"`
	Short     string       `json:"short"`
	CreatedAt CreatedAtSID `json:"created_at"`
}

// End typings

func createShortUrl(payload ShortPayload) (ShortResultSID, error) {
	var jsonString []byte = []byte(`{"url":"` + payload.Url + `"}`)
	response, err := http.Post(config.ShortenerSDotId, "application/json", bytes.NewBuffer(jsonString))
	if err != nil {
		return ShortResultSID{}, err
	} else {
		defer response.Body.Close()
		if response.StatusCode != 200 {
			return ShortResultSID{}, fmt.Errorf("status code: %v", response.StatusCode)
		}
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return ShortResultSID{}, err
		}
		var result OriginalResponseSID
		if err := json.Unmarshal(body, &result); err != nil {
			return ShortResultSID{}, err
		}

		return ShortResultSID{
			Url:         "https://s.id/" + result.Short,
			OriginalUrl: result.LongUrl,
			CreatedAt:   result.CreatedAt,
		}, nil
	}
}
