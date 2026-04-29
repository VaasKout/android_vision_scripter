// Package models ...
package models

import (
	"fmt"
	"strings"
)

// AdbDevice ...
type AdbDevice struct {
	Serial        string `json:"serial"`
	Brand         string `json:"brand"`
	Device        string `json:"device"`
	Locale        string `json:"locale"`
	Model         string `json:"model"`
	OsVersion     string `json:"os_version"`
	Manufacturer  string `json:"manufacturer"`
	MarketingName string `json:"marketing_name"`
	ScrcpyRunning bool   `json:"scrcpy_running"`
}

// FilterLocales ...
func (a *AdbDevice) FilterLocales() {
	if a == nil || a.Locale == "" {
		return
	}
	var systemLocales = strings.Split(a.Locale, ",")
	if len(systemLocales) == 0 {
		return
	}
	a.Locale = systemLocales[0]
}

// ToModelOs ...
func (a *AdbDevice) ToModelOs() string {
	if a == nil || a.Model == "" || a.OsVersion == "" {
		return ""
	}
	return fmt.Sprintf("%s_%s", a.Model, a.OsVersion)
}
