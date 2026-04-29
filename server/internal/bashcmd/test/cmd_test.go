package bashcmd

import (
	"android_vision_scripter/config"
	"android_vision_scripter/internal/bashcmd"
	"android_vision_scripter/internal/filesdb"
	"android_vision_scripter/pkg/logger"
	"testing"
)

const (
	TestSerial = "xxx" //serial number of the device
)

func TestGetDeviceList(t *testing.T) {
	var fileProps = &config.FilesProps{
		Logs: "./logs/",
	}
	logAPI := logger.New(logger.INFO, true)
	var filesDB = filesdb.New(fileProps)
	var cmdRunner = bashcmd.New(filesDB, logAPI)

	devices := cmdRunner.GetDevicesList()
	t.Log(devices)
	t.Log(len(devices))
}

func TestScreenshot(t *testing.T) {
	var fileProps = &config.FilesProps{
		Logs: "./logs/",
	}
	logAPI := logger.New(logger.INFO, true)
	var filesDB = filesdb.New(fileProps)
	var cmdRunner = bashcmd.New(filesDB, logAPI)

	screenshot := cmdRunner.ScreenShot(TestSerial)
	if screenshot == "" {
		t.Fatal("screenshot is empty")
	}

	t.Log(screenshot)
}
