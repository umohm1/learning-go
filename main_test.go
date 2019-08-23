package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(handler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusOK)
	}

	expected := `Hello World!`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

	func TestRouter(t *testing.T) {
		// Instantiate the router using the constructor function
		// efined previously
		r := newRouter()

		// Create a new server using the "httptest" libraries `NewServer` method
		// Documentation : https://golang.org/pkg/net/http/httptest/#NewServer
		mockServer := httptest.NewServer(r)

		// The mock server created runs a server and exposes its location in the URL attribute
		// We make a GET request to the "hello" route we defined in the router
		resp, err := http.Get(mockServer.URL + "/hello")

		// Handle any unexpected error
		if err != nil {
			t.Fatal(err)
		}

		// Want status to be 200 (ok)
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Status should be ok, got %d", resp.StatusCode)
		}

		// Response body is read, and converted to a string
		defer resp.Body.Close()
		// read the body into a bunch of bytes (b)
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		// convert the bytes into a string
		respString := string(b)
		expected := "Hello World!"

		// Response should match what's defined in the handler
		// If it's not "Hello World!", then it confirms that the route is correct
		if respString != expected {
			t.Errorf("Response should be %s, got %s", expected, respString)
		}

	}

	func TestRouterForNonExistentRoute(t *testing.T) {
		r := newRouter()
		mockServer := httptest.NewServer(r)
		resp, err := http.Post(mockServer.URL+ "/hello", "", nil)

		if err != nil {
			t.Fatal(err) 
		}

		// set status to 405 (method not allowed)
		if resp.StatusCode != http. StatusMethodNotAllowed {
			t.Errorf("Status should be 405, got %d", resp.StatusCode)

		}

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
