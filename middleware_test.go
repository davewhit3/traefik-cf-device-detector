package traefik_cf_device_detector_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	mw "github.com/davewhit3/traefik-cf-device-detector"
)

func TestBasic(t *testing.T) {
	called := false
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) { called = true })

	instance, err := mw.New(context.TODO(), next, mw.CreateConfig(), "traefikuseragent")
	if err != nil {
		t.Fatalf("Error creating %v", err)
	}

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)

	instance.ServeHTTP(recorder, req)
	if recorder.Result().StatusCode != http.StatusOK {
		t.Fatalf("Invalid return code")
	}
	if called != true {
		t.Fatalf("next handler was not called")
	}
}

func TestParse(t *testing.T) {
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	instance, _ := mw.New(context.TODO(), next, mw.CreateConfig(), "traefikuseragent")

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	req.Header.Set(mw.UserAgentHeader, "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.97 Safari/537.11")

	instance.ServeHTTP(recorder, req)

	assertHeader(t, req, mw.DeviceIsMobileHeader, "false")
	assertHeader(t, req, mw.DeviceIsDesktopHeader, "true")
	assertHeader(t, req, mw.DeviceIsTabletHeader, "false")
}

func assertHeader(t *testing.T, req *http.Request, key, expected string) {
	t.Helper()
	if req.Header.Get(key) != expected {
		t.Fatalf("invalid value of header [%s] != %s", key, req.Header.Get(key))
	}
}
