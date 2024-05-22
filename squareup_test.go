package squareup

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	mux *http.ServeMux

	ctx = context.TODO()

	client *Client

	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient(nil, ModeSandbox)
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	if expected != r.Method {
		t.Errorf("Request method = %v, expected %v", r.Method, expected)
	}
}

type values map[string]string

func testFormValues(t *testing.T, r *http.Request, values values) {
	expected := url.Values{}
	for k, v := range values {
		expected.Add(k, v)
	}

	err := r.ParseForm()
	if err != nil {
		t.Fatalf("parseForm(): %v", err)
	}

	if !reflect.DeepEqual(expected, r.Form) {
		t.Errorf("Request parameters = %v, expected %v", r.Form, expected)
	}
}

func testURLParseError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func testClientServices(t *testing.T, c *Client) {
	services := []string{
		"TerminalAction",
	}
	cp := reflect.ValueOf(c)
	cv := reflect.Indirect(cp)

	for _, s := range services {
		if cv.FieldByName(s).IsNil() {
			t.Errorf("c.%s shouldn't be nil", s)
		}
	}
}

func testClientDefaultBaseURL(t *testing.T, c *Client) {
	if c.BaseURL == nil || c.BaseURL.String() != defaultBaseURL {
		t.Errorf("NewClient BaseURL = %v, expected %v", c.BaseURL, defaultBaseURL)
	}
}

func testClientDefaultUserAgent(t *testing.T, c *Client) {
	if c.UserAgent != userAgent {
		t.Errorf("NewClient UserAgent = %v, expected %v", c.UserAgent, userAgent)
	}
}

func testClientDefaults(t *testing.T, c *Client) {
	testClientDefaultBaseURL(t, c)
	testClientDefaultUserAgent(t, c)
	testClientServices(t, c)
}

func TestNewClient(t *testing.T) {
	c := NewClient(nil, ModeSandbox)
	testClientDefaults(t, c)
}

func TestNewFromToken(t *testing.T) {
	c := NewFromToken("myToken", ModeSandbox)
	testClientDefaults(t, c)
}

func TestNewFromToken_cleaned(t *testing.T) {
	testTokens := []string{"myToken ", " myToken", " myToken ", "'myToken'", " 'myToken' "}
	expected := "Bearer myToken"

	setup()
	defer teardown()

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for _, tt := range testTokens {
		t.Run(tt, func(t *testing.T) {
			c := NewFromToken(tt, ModeSandbox)
			req, _ := c.NewRequest(ctx, http.MethodGet, server.URL+"/foo", nil)
			resp, err := c.Do(ctx, req, nil)
			if err != nil {
				t.Fatalf("Do(): %v", err)
			}

			authHeader := resp.Request.Header.Get("Authorization")
			if authHeader != expected {
				t.Errorf("Authorization header = %v, expected %v", authHeader, expected)
			}
		})
	}
}

func TestNew(t *testing.T) {
	c, err := New(nil, ModeSandbox)

	if err != nil {
		t.Fatalf("New(): %v", err)
	}
	testClientDefaults(t, c)
}

func TestNewRequest_get(t *testing.T) {
	c := NewClient(nil, ModeSandbox)

	inURL, outURL := "/foo", defaultBaseURL+"foo"
	req, _ := c.NewRequest(ctx, http.MethodGet, inURL, nil)

	// test relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewRequest(%v) URL = %v, expected %v", inURL, req.URL, outURL)
	}

	// test the content-type header is not set
	if contentType := req.Header.Get("Content-Type"); contentType != "" {
		t.Errorf("NewRequest() Content-Type = %v, expected empty string", contentType)
	}

	// test default user-agent is attached to the request
	userAgent := req.Header.Get("User-Agent")
	if c.UserAgent != userAgent {
		t.Errorf("NewRequest() User-Agent = %v, expected %v", userAgent, c.UserAgent)
	}
}

func TestCustomUserAgent(t *testing.T) {
	ua := "testing/0.0.1"
	c, err := New(nil, ModeSandbox, SetUserAgent(ua))

	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	expected := fmt.Sprintf("%s %s", ua, userAgent)
	if got := c.UserAgent; got != expected {
		t.Errorf("New() UserAgent = %s; expected %s", got, expected)
	}
}

func TestCustomBaseURL(t *testing.T) {
	baseURL := "http://localhost/foo"
	c, err := New(nil, ModeSandbox, SetBaseURL(baseURL))

	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	expected := baseURL
	if got := c.BaseURL.String(); got != expected {
		t.Errorf("New() BaseURL = %s; expected %s", got, expected)
	}
}

func TestCustomBaseURL_badURL(t *testing.T) {
	baseURL := ":"
	_, err := New(nil, ModeSandbox, SetBaseURL(baseURL))

	testURLParseError(t, err)
}
