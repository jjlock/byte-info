package scraper

import (
	"errors"
	"testing"

	"github.com/matryer/is"
)

func TestGetByte(t *testing.T) {
	ts := newTestScraper(t, 200, "testdata/byte.html")
	expected := &Byte{
		ID:           "4rHqwfCzgCB",
		User:         "byte",
		UserURL:      "https://byte.co/byte",
		ThumbnailURL: "https://testcdn.com/videos/byte.jpg",
		Caption:      "one day he will be free",
		CreatedAt:    "2mo",
		Loops:        160000,
		URLs:         []string{"https://byte.co/byte/4rHqwfCzgCB", "https://byte.co/b/4rHqwfCzgCB"},
	}

	got, err := ts.GetByte("4rHqwfCzgCB")
	if err != nil {
		t.Errorf("GetByte should return a nil error but got: %v", err)
	}

	is := is.New(t)
	is.Equal(got, expected)
}

func TestGetByteNotFound(t *testing.T) {
	ts := newTestScraper(t, 404, "")
	byte, err := ts.GetByte("foo")
	if byte != nil {
		t.Errorf("GetByte should return a nil Byte but got: %#v", byte)
	}

	if err == nil {
		t.Fatal("GetByte should return *RequestError but got a nil error")
	}

	var rerr *RequestError
	if !errors.As(err, &rerr) {
		t.Fatalf("GetByte should return *RequestError but got: %v", err)
	}

	if rerr.StatusCode != 404 {
		t.Fatalf("GetByte should return *RequestError with a 404 HTTP status code but got a %d HTTP status code", rerr.StatusCode)
	}
}
