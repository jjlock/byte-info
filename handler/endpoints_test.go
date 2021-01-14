package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jjlock/byte-scraper-api/scraper"
	"github.com/matryer/is"
)

func init() {
	// disable logging during testing
	log.SetOutput(ioutil.Discard)
}

var errorTests = map[string]struct {
	given    error
	expected errorResponse
}{
	// Given a 4xx response (except 404) from byte.co respond with a 500 HTTP status code
	"400": {
		&scraper.RequestError{StatusCode: 400, Message: "test scraper error message 400"},
		errorResponse{Status: 500, Message: "test response error message 500"},
	},
	// Given a 404 response from byte.co respond with a 404 HTTP status code
	"404": {
		&scraper.RequestError{StatusCode: 404, Message: "test scraper error message 404"},
		errorResponse{Status: 404, Message: "test response error message 404"},
	},
	// Given a 5xx response from byte.co respond with a 503 HTTP status code
	"500": {
		&scraper.RequestError{StatusCode: 500, Message: "test scraper error message 500"},
		errorResponse{Status: 503, Message: "test response error message 503"},
	},
	// Given an error not of type *scraper.RequestError respond with a 500 HTTP status code
	"Non RequestError Error": {
		errors.New("non RequestError error"),
		errorResponse{Status: 500, Message: "test response error message 500"},
	},
}

type mockScraper struct {
	user *scraper.User
	byte *scraper.Byte
}

func (ms *mockScraper) GetUser(username string) (*scraper.User, error) {
	if data, ok := errorTests[username]; ok {
		return nil, data.given
	}

	return ms.user, nil
}

func (ms *mockScraper) GetByte(id string) (*scraper.Byte, error) {
	if data, ok := errorTests[id]; ok {
		return nil, data.given
	}

	return ms.byte, nil
}

// newTestScraperHandler returns a ScraperHandler instance for testing given mock
// objects to return in mockScraper methods.
func newTestScraperHandler(t *testing.T, mocks ...interface{}) *ScraperHandler {
	ms := &mockScraper{}
	for _, mock := range mocks {
		switch m := mock.(type) {
		case *scraper.User:
			ms.user = m
		case *scraper.Byte:
			ms.byte = m
		default:
			t.Fatal("Unexpected data type used to initialize the mock scraper")
		}
	}

	handler := &ScraperHandler{
		scraper: ms,
		router:  mux.NewRouter(),
	}

	handler.routes()

	return handler
}

// checkHeader checks if a http.Response header has the correct fields and values.
func checkHeader(t *testing.T, resp *http.Response) {
	ct := "application/json; charset=utf-8"
	actualct := resp.Header.Get("Content-Type")
	if actualct != ct {
		t.Errorf(`In response header, expected Content-Type: "%s" but got Content-Type: "%s"`, ct, actualct)
	}

	xcto := "nosniff"
	actualxcto := resp.Header.Get("X-Content-Type-Options")
	if actualxcto != xcto {
		t.Errorf(`In response header, expected X-Content-Type-Options: "%s" but got X-Content-Type-Options: "%s"`, xcto, actualxcto)
	}
}

// testEndpoint tests a handler for a given endpoint assuming a successful 200 response from byte.co.
func testEndpoint(t *testing.T, sh *ScraperHandler, path string, expected, actual interface{}) {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	sh.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected response with a 200 HTTP status code but got a %d HTTP status code", resp.StatusCode)
	}

	checkHeader(t, resp)

	// actual should be an empty composite literal of expected's type
	if err := json.NewDecoder(resp.Body).Decode(actual); err != nil {
		t.Fatal(err)
	}

	is := is.New(t)
	is.Equal(actual, expected)
}

// testErrors runs tests on a handler assuming a non-200 response from byte.co or any other errors
// encountered when handling a client request.
func testErrors(t *testing.T, sh *ScraperHandler, basepath string) {
	for test, data := range errorTests {
		t.Run("Handle "+test, func(t *testing.T) {
			req, err := http.NewRequest("GET", basepath+"/"+test, nil)
			if err != nil {
				t.Fatal(err)
			}

			w := httptest.NewRecorder()
			sh.ServeHTTP(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			if resp.StatusCode != data.expected.Status {
				t.Errorf("Expected response with a %d HTTP status code but got a %d HTTP status code", data.expected.Status, resp.StatusCode)
			}

			checkHeader(t, resp)

			actual := &errorResponse{}
			if err := json.NewDecoder(resp.Body).Decode(actual); err != nil {
				t.Fatal(err)
			}

			if actual.Status != data.expected.Status {
				t.Errorf("In JSON response, expected Status: %d but got Status: %d", data.expected.Status, actual.Status)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	user := &scraper.User{
		Username:        "byte",
		ProfileImageURL: "https://testcdn.com/profiles/byte.jpg",
		Bio:             "say something nice",
		RecentByteIDs:   []string{"4rHqwfCzgCB", "75COWrlEIjq"},
		RecentByteURLs:  []string{"https://byte.co/@byte/4rHqwfCzgCB", "https://byte.co/@byte/75COWrlEIjq"},
		URL:             "https://byte.co/@byte",
	}

	tsh := newTestScraperHandler(t, user)
	testEndpoint(t, tsh, "/api/users/byte", user, &scraper.User{})
}

func TestGetUserErrors(t *testing.T) {
	tsh := newTestScraperHandler(t)
	testErrors(t, tsh, "/api/users")
}

func TestGetByte(t *testing.T) {
	byte := &scraper.Byte{
		ID:           "4rHqwfCzgCB",
		User:         "byte",
		UserURL:      "https://byte.co/byte",
		ThumbnailURL: "https://testcdn.com/videos/byte.jpg",
		Caption:      "one day he will be free",
		CreatedAt:    "2mo",
		Loops:        160000,
		URLs:         []string{"https://byte.co/byte/4rHqwfCzgCB", "https://byte.co/b/4rHqwfCzgCB"},
	}

	tsh := newTestScraperHandler(t, byte)
	testEndpoint(t, tsh, "/api/bytes/4rHqwfCzgCB", byte, &scraper.Byte{})
}

func TestGetByteErrors(t *testing.T) {
	tsh := newTestScraperHandler(t)
	testErrors(t, tsh, "/api/bytes")
}
