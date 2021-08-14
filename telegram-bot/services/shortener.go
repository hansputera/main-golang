package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Shortener_SDotIDResult struct {
	Url         string
	OriginalUrl string
	CreatedAt   Shortener_SDotIDCreatedAt
}

type Shortener_SDotIDCreatedAt struct {
	Date         string `json:"date"`
	Timezone     string `json:"timezone"`
	TimezoneType int    `json:"timezone_type"`
}

type Shortener_SDotIDOriginalResult struct {
	LongUrl   string                    `json:"long_url"`
	Short     string                    `json:"short"`
	CreatedAt Shortener_SDotIDCreatedAt `json:"created_at"`
}

func Shortener_SDotID(url string) (Shortener_SDotIDResult, error) {
	jsonString := []byte(`{"url":"` + url + `"}`)
	response, err := http.Post("https://home.s.id/api/public/link/shorten", "application/json", bytes.NewBuffer(jsonString))
	if err != nil {
		return Shortener_SDotIDResult{}, err
	} else {
		defer response.Body.Close()
		if response.StatusCode != 200 {
			return Shortener_SDotIDResult{}, fmt.Errorf("rate limit issue")
		} else {
			body, errorBody := ioutil.ReadAll(response.Body)
			if errorBody != nil {
				return Shortener_SDotIDResult{}, errorBody
			}
			var result Shortener_SDotIDOriginalResult
			if err := json.Unmarshal(body, &result); err != nil {
				return Shortener_SDotIDResult{}, err
			}
			return Shortener_SDotIDResult{
				Url:         "https://s.id/" + result.Short,
				OriginalUrl: result.LongUrl,
				CreatedAt:   result.CreatedAt,
			}, nil
		}
	}
}
