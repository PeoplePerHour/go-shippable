// Package shippable is a library providing access to Shippable API.
package shippable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	version          = "0.1"
	userAgent        = "shippable-go/" + version
	defaultEndpoint  = "https://api.shippable.com/"
	defaultMediaType = "application/json"
)

// Client holds the client to Shippable API along with an authorization token and
// pointers to each Shippable API service/endpoint.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Endpoint for API requests.  Defaults to the Live Shippable API
	// should have a trailing slash
	Endpoint *url.URL

	// The client's Shippable API token used for Authentication
	Token string

	// User agent used when communicating with the Shippable API.
	UserAgent string

	// Routes/entities used for talking to different parts of the Shippable API.
	Projects *ProjectService
	Accounts *AccountService
	Workflow *WorkflowService
}

// Response is a Shippable API response. This wraps the standard http.Response
// returned from Shippable.
type Response struct {
	*http.Response
	// Would be nice to add some extra fields to be used with convinient getters
}

// newResponse creates a new Response for the provided http.Response.
// TODO: Add extra logic in some added Response fields
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

// NewClient returns a new Shippable API client. An authorization token
// is needed to interact with the service.
func NewClient(token string) (c *Client) {
	httpClient := http.DefaultClient
	endpoint, _ := url.Parse(defaultEndpoint)
	c = &Client{
		client:    httpClient,
		UserAgent: userAgent,
		Endpoint:  endpoint,
		Token:     token,
	}
	c.Projects = &ProjectService{client: c}
	c.Accounts = &AccountService{client: c}
	c.Workflow = &WorkflowService{client: c}

	return
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the Endpoint of the Client.
// Relative URLs should always be specified without a preceding slash.  If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.Endpoint.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", defaultMediaType)
	req.Header.Add("Authorization", "apiToken "+c.Token)
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	return req, nil
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

	if response.StatusCode >= 300 {
		return response, fmt.Errorf("Request: '%s' failed!", ToString(req))
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}
	return response, err
}
