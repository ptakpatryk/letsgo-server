package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
)

func newTestApplication(t *testing.T) *application {
	return &application{
		// infoLog: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		// errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(io.Discard, "", 0),
		errorLog: log.New(io.Discard, "", 0),
	}
}

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewTLSServer(h)

  jar, err := cookiejar.New(nil)
  if err != nil {
    t.Fatal()
  }

  ts.Client().Jar = jar

  ts.Client().CheckRedirect = func (req *http.Request, via []*http.Request) error  {
    return http.ErrUseLastResponse
  }
  
	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	return rs.StatusCode, rs.Header, string(body)
}
