package test

import (
	"android_vision_scripter/internal/server"
	"net/http"
	"testing"
)

// Device paths...
const (
	PingPath   = LocalURL + server.Ping
	DevicePath = LocalURL + server.Devices
)

func TestPingServer(t *testing.T) {
	var url = PingPath
	var data = map[string]any{}
	makeHTTPRequest(
		url,
		http.MethodGet,
		[]byte{},
		&data,
	)
	t.Log(data)
}

func TestGetDevices(t *testing.T) {
	var url = DevicePath
	var data = ""
	makeHTTPRequest(
		url,
		http.MethodGet,
		[]byte{},
		&data,
	)
	t.Log(data)
}
