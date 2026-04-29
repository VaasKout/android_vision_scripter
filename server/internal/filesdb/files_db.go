// Package filesdb ...
package filesdb

import (
	"android_vision_scripter/config"
	"regexp"
)

var eventRegex = regexp.MustCompile(`\[\s*(\d+\.\d+)]\s+(/dev/input/event\d+)\s+([0-9A-Fa-f]{4})\s+([0-9A-Fa-f]{4})\s+([0-9A-Fa-f]{8})`)

// Basic directories for phone data
const (
	ScreenshotDir = "screenshot"
	RecordsDir    = "records"
	ScriptsDir    = "scripts"
	TesseractDir  = "tesseract"
	Keyboards     = "keyboards"
)

// Script JSONs
const (
	RunJSON = "run.json"
	TmpZone = "tmp.png"
)

// X and Y keys in hex format
const (
	HexXKey = "0035"
	HexYKey = "0036"
)

// FilesDB ...
type FilesDB interface {
	Create
	Read
	Delete
}

type filesDBImpl struct {
	filesProps *config.FilesProps
}

// New instance of FilesDB
func New(
	filesProps *config.FilesProps,
) FilesDB {
	return &filesDBImpl{
		filesProps: filesProps,
	}
}
