package sillyapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

var (
	ErrNon200Response = errors.New("non-200 response")
	ErrEmptyJoke      = errors.New("I didn't get the joke")
)

type Joke struct {
	IconURL string `json:"icon_url"`
	ID      string `json:"id"`
	URL     string `json:"url"`
	Value   string `json:"value"`
}

func GetJoke(urls ...string) (string, error) {
	url := "https://api.chucknorris.io/jokes/random"
	if len(urls) > 0 {
		url = urls[0]
	}

	// make api call to a joke service
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	// check result
	if resp.StatusCode != http.StatusOK {
		return "", ErrNon200Response
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// unpack result
	var joke Joke
	err = json.Unmarshal(body, &joke)
	if err != nil {
		return "", err
	}

	// check result contents
	// validation?
	if joke.Value == "" {
		return "", ErrEmptyJoke
	}

	// return the joke
	return strings.ToUpper(joke.Value), nil
}
