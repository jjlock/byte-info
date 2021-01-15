package scraper

import (
	"errors"
	"reflect"
	"testing"
)

func TestGetUser(t *testing.T) {
	ts := newTestScraper(t, 200, "testdata/user.html")
	expected := &User{
		Username:        "byte",
		ProfileImageURL: "https://testcdn.com/profiles/byte.jpg",
		Bio:             "say something nice",
		RecentByteIDs:   []string{"4rHqwfCzgCB", "75COWrlEIjq"},
		RecentByteURLs:  []string{"https://byte.co/@byte/4rHqwfCzgCB", "https://byte.co/@byte/75COWrlEIjq"},
		URL:             "https://byte.co/@byte",
	}

	got, err := ts.GetUser("byte")
	if err != nil {
		t.Errorf("GetUser should return a nil error but got: %v", err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("\nGot: %#v\nWant: %#v", got, expected)
	}
}

func TestGetUserNotFound(t *testing.T) {
	ts := newTestScraper(t, 404, "")
	user, err := ts.GetUser("foo")
	if user != nil {
		t.Errorf("GetUser should return a nil User but got: %#v", user)
	}

	if err == nil {
		t.Fatal("GetUser should return *RequestError but got a nil error")
	}

	var rerr *RequestError
	if !errors.As(err, &rerr) {
		t.Fatalf("GetUser should return *RequestError but got: %v", err)
	}

	if rerr.StatusCode != 404 {
		t.Fatalf("GetUser should return *RequestError with a 404 HTTP status code but got a %d HTTP status code", rerr.StatusCode)
	}
}
