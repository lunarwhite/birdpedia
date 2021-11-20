package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	// The mock server we created runs a server and exposes its location in the URL attribute.
	// We make a GET request to the "hello" route we defined in the router.
	resp, err := http.Get(mockServer.URL + "/hello")

	if err != nil {
		t.Fatal(err)
	}
	// We want our status to be 200 (ok).
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	// In the next few lines, the response body is read, and converted to a string.
	defer resp.Body.Close()

	// Read the body into a bunch of bytes (b).
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	// Convert the bytes to a string.
	respString := string(b)
	expected := "Hello World!"
	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

func TestRouterForNonExistentRoute(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	// Most of the code is similar. The only difference is that now we make a
	// request to a route we know we didn't define, like the `POST /hello` route.
	resp, err := http.Post(mockServer.URL+"/hello", "", nil)

	if err != nil {
		t.Fatal(err)
	}
	// We want our status to be 405 (method not allowed)
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", resp.StatusCode)
	}

	// Also mostly the same, except this time, we expect an empty body
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := ""
	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}

}

func TestHandler(t *testing.T) {
	// Here, we form a new HTTP request. This is the request that's going to be passed to our handler.
	// The first argument is the method, the second argument is the route,
	// and the third is the request body, which we don't have in this case.
	req, err := http.NewRequest("GET", "", nil)

	// In case there is an error in forming the request, we fail and stop the test
	if err != nil {
		t.Fatal(err)
	}

	// We use Go's httptest library to create an http recorder.
	// This recorder will act as the target of our http request.
	recorder := httptest.NewRecorder()

	// Create an HTTP handler from our handler function.
	hf := http.HandlerFunc(handler)

	// Serve the HTTP request to our recorder. This is the line that actually
	// executes our the handler that we want to test
	hf.ServeHTTP(recorder, req)

	// Check the status code is what we expect.
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `Hello World!`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
