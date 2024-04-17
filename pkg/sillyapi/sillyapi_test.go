package sillyapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestGetJokeReal(t *testing.T) {
	t.Skip()

	joke, err := GetJoke()
	if err != nil {
		t.Errorf("received unexpected error: %v", err)
	}

	gotLowerCase, err := regexp.Match("[a-z]", []byte(joke))
	if err != nil {
		t.Errorf("failed regexp: %v", err)
	}
	if gotLowerCase {
		t.Errorf("go unexpected lower case: %s", joke)
	}
}

func TestGetJokeFake(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/jokes/random", func(w http.ResponseWriter, r *http.Request) {
		payload := map[string]string{
			"icon_url": "https://assets.chucknorris.host/img/avatar/chuck-norris.png",
			"id":       "tudUv6WpRiefRiv4gYUHgg",
			"url":      "",
			"value":    "the 300 spartans would have not stood a chance against the pesians if Chuck Norris had not let them use the 'total gym'.",
		}

		body, err := json.Marshal(payload)
		if err != nil {
			t.Errorf("marshal error: %v", err)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(body))
	})
	server := httptest.NewServer(mux)

	joke, err := GetJoke(server.URL + "/jokes/random")
	if err != nil {
		t.Errorf("received unexpected error: %v", err)
	}

	gotLowerCase, err := regexp.Match("[a-z]", []byte(joke))
	if err != nil {
		t.Errorf("failed regexp: %v", err)
	}
	if gotLowerCase {
		t.Errorf("go unexpected lower case: %s", joke)
	}
}
