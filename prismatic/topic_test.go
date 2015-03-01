package prismatic

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the GitHub client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient(nil, "foo")
	url, _ := url.Parse(server.URL)
	client.BaseURL = url

}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func TestSearchForInterest(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/topic/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprintln(w, `{"results": [ {"topic": "Clojure", "id": 1000} ]}`)
	})

	results, _, err := client.Topics.SearchForInterest("Clojure")
	if err != nil {
		t.Error("SearchForInterest returned error: %v", err)
	}

	want := ResultResponse{
		[]Result{{Topic: "Clojure", Id: 1000}},
	}

	if !reflect.DeepEqual(results, want) {
		t.Errorf("SearchForInterest returned %+v, want %+v", results, want)
	}

}

func TestSearchForRelatedTopic(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/topic/topic", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprintln(w, `{"topics": [ {"topic": "Clojure", "id": 1000, "score": 100} ]}`)
	})

	results, _, err := client.Topics.SearchForRelatedTopic(10222)
	if err != nil {
		t.Error("SearchForRelatedTopic returned error: %v", err)
	}

	want := TopicResponse{
		[]Topic{{Topic: "Clojure", Id: 1000, Score: 100}},
	}

	if !reflect.DeepEqual(results, want) {
		t.Errorf("SearchForRelatedTopic returned %+v, want %+v", results, want)
	}
}
