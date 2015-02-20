/*
Package prismatic provides a client for using the Prismatic Interests API.
Access different parts of the  API using the various
services:
         apiToken := "api-token"
         client := prismatic.NewClient(nil, apiToken)
         // search for an interest
         results, _, err := client.Topics.SearchForInterest("Clojure")
The full Prismatic Interests API is documented at https://github.com/Prismatic/interest-graph.
*/
package prismatic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "http://interest-graph.getprismatic.com/"
)

// A Client manages communication with the Prismatic Interest API.
type Client struct {
	// HTTP client used to communicate with the API
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// Prismatic issued API Token.
	ApiToken string

	// Services used for talking to different parts of the Prismatic API.
	Topics *TopicService
}

// NewClient returns a new Prismatic API client.  If a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(httpClient *http.Client, ApiToken string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	if ApiToken == "" {
		log.Fatalln("ApiToken must be provided as parameter to NewClient()")
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, ApiToken: ApiToken}
	c.Topics = &TopicService{client: c}

	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.  If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, payload url.Values) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	u := c.BaseURL.ResolveReference(rel)

	log.Printf("url: %v", u)

	body := &bytes.Buffer{}

	if payload != nil {
		body = bytes.NewBufferString(payload.Encode())
	}

	req, err := http.NewRequest(method, u.String(), body)

	//set the api token header
	req.Header.Set("X-API-TOKEN", c.ApiToken)

	if err != nil {
		return nil, err
	}

	return req, nil
}

// Response is a Prismatic API response.  This wraps the standard http.Response
// returned from Prismatic.
type Response struct {
	*http.Response
}

// newResponse creats a new Response for the provided http.Response.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%v: %v: %v", e.Response.Request.Method, e.Response.Request.URL,
		e.Response.StatusCode)
}

// CheckResponse checks the API response for errors, and returns them if
// present.  A response is considered an error if it has a status code outside
// the 200 range.  API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse.  Any other
// response body will be silently ignored.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

// Do sends an API request and returns the API response.  The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := newResponse(resp)

	err = CheckResponse(resp)

	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return response, err

}
