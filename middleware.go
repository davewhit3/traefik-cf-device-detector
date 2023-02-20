// Package traefik_cf_device_detector a demo plugin.
package traefik_cf_device_detector

import (
	"context"
	"net/http"
	"strconv"

	"github.com/mileusna/useragent"
)

const (
	// UserAgentHeader header.
	UserAgentHeader = "User-Agent"

	// DeviceIsMobileHeader header.
	DeviceIsMobileHeader = "CloudFront-Is-Mobile-Viewer"

	// DeviceIsDesktopHeader header.
	DeviceIsDesktopHeader = "CloudFront-Is-Desktop-Viewer"

	// DeviceIsTabletHeader header.
	DeviceIsTabletHeader = "CloudFront-Is-Tablet-Viewer"

	// DeviceIsSmartTVHeader header.
	DeviceIsSmartTVHeader = "CloudFront-Is-SmartTV-Viewer"
)

// Config the plugin configuration.
type Config struct{}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// CfDeviceDetector a CfDeviceDetector plugin.
type CfDeviceDetector struct {
	next http.Handler
	name string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &CfDeviceDetector{
		next: next,
		name: name,
	}, nil
}

func (mw *CfDeviceDetector) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ua := useragent.Parse(req.Header.Get(UserAgentHeader))

	req.Header.Set(DeviceIsMobileHeader, strconv.FormatBool(ua.Mobile))
	req.Header.Set(DeviceIsDesktopHeader, strconv.FormatBool(ua.Desktop))
	req.Header.Set(DeviceIsTabletHeader, strconv.FormatBool(ua.Tablet))
	req.Header.Set(DeviceIsSmartTVHeader, strconv.FormatBool(false))

	mw.next.ServeHTTP(rw, req)
}
