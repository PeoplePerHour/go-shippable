package shippable

import (
	/*
			"bytes"
			"encoding/json"
			"fmt"
			"io/ioutil"
			"os"
			"path"
			"reflect"
			"strings"
			"time"
		"testing"
	*/
	"net/http"
	"net/http/httptest"
	"net/url"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the Shippable client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

// setup sets up a test HTTP server along with a shippable.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// shippable client configured to use test server
	client = NewClient("token")
	url, _ := url.Parse(server.URL)
	client.Endpoint = url
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}
