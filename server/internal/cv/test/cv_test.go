package test

import (
	"android_vision_scripter/config"
	"android_vision_scripter/internal/bashcmd"
	"android_vision_scripter/internal/cv"
	"android_vision_scripter/internal/filesdb"
	"android_vision_scripter/pkg/core/file"
	"android_vision_scripter/pkg/logger"
	"fmt"
	"image"
	"path/filepath"

	"testing"

	"gocv.io/x/gocv"
)

const (
	TestSerial = "xxx"        //serial number of the device
	TestImage  = "./test.png" //to add some template, create screenshot and crop any zone to compare
)

func TestGetTextFromImage(t *testing.T) {
	var fileProps = &config.FilesProps{
		Logs: "./logs",
	}
	var logAPI = logger.New(logger.INFO, true)
	var filesDB = filesdb.New(fileProps)
	var cmdRunner = bashcmd.New(filesDB, logAPI)
	var cvAPI = cv.New(cmdRunner, logAPI)

	screenshot := cmdRunner.ScreenShot(TestSerial)
	if screenshot == "" {
		t.Fatal("screenshot is empty")
	}

	dir, err := file.FindDirectoryOfFile(screenshot)
	if err != nil {
		t.Fatal(err)
	}
	img := gocv.IMRead(screenshot, gocv.IMReadColor)
	if img.Empty() {
		t.Fatal("could not read screenshot image")
	}
	defer img.Close()

	ocrParams := cv.InitOcrParams("", "eng", cv.PsmText, cv.OemText, cv.WhiteTheme)
	ocrResult, err := cvAPI.FindTextRectangles(&img, dir, ocrParams)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ocrResult)
	var rectangles = []image.Rectangle{}
	for _, ocr := range ocrResult {
		var imgRect = ocr.Rectangle.ToImageRectangle()
		if imgRect == nil || cv.ImageRectIsEmpty(imgRect) {
			continue
		}
		rectangles = append(rectangles, *imgRect)
	}

	err = cvAPI.DrawRectangles(img, rectangles, false)
	if err != nil {
		t.Fatal(err)
	}
	params := []int{gocv.IMWriteJpegQuality, 90}
	if ok := gocv.IMWriteWithParams(screenshot, img, params); !ok {
		fmt.Println("could not write image " + screenshot)
	}
}

func TestResetKeyboardKeys(t *testing.T) {
	var fileProps = &config.FilesProps{
		Logs: "./logs",
	}
	var logAPI = logger.New(logger.INFO, true)
	var filesDB = filesdb.New(fileProps)
	var cmdRunner = bashcmd.New(filesDB, logAPI)
	var cvAPI = cv.New(cmdRunner, logAPI)

	screenshot := cmdRunner.ScreenShot(TestSerial)
	if screenshot == "" {
		t.Fatal("screenshot is empty")
	}

	img := gocv.IMRead(screenshot, gocv.IMReadColor)
	if img.Empty() {
		t.Fatal("could not read screenshot image")
	}
	defer img.Close()

	keyboardDir := filesDB.CreateLogsDir(TestSerial, filesdb.Keyboards, "rus")
	tesseractDir := filesDB.CreateLogsDir(TestSerial, filesdb.TesseractDir)
	ocrResult := cvAPI.ResetKeyboardKeys(keyboardDir, tesseractDir, img, "ru", false)

	var screenshotWithRects = filepath.Join(tesseractDir, "screenshot.png")
	var rectangles = []image.Rectangle{}
	for _, result := range ocrResult {
		rectangles = append(rectangles, *result.Rectangle.ToImageRectangle())
	}
	cvAPI.DrawRectangles(img, rectangles, false)
	params := []int{gocv.IMWriteJpegQuality, 90}
	if ok := gocv.IMWriteWithParams(screenshotWithRects, img, params); !ok {
		fmt.Println("could not write image " + screenshotWithRects)
	}
}

func TestGetKeyboardKeys(t *testing.T) {
	var fileProps = &config.FilesProps{
		Logs: "./logs",
	}
	var logAPI = logger.New(logger.INFO, true)
	var filesDB = filesdb.New(fileProps)
	var cmdRunner = bashcmd.New(filesDB, logAPI)
	var cvAPI = cv.New(cmdRunner, logAPI)

	screenshot := cmdRunner.ScreenShot(TestSerial)
	if screenshot == "" {
		t.Fatal("screenshot is empty")
	}

	img := gocv.IMRead(screenshot, gocv.IMReadColor)
	if img.Empty() {
		t.Fatal("could not read screenshot image")
	}
	defer img.Close()

	keyboardDir := filesDB.CreateLogsDir(TestSerial, filesdb.Keyboards, "rus")
	tesseractDir := filesDB.CreateLogsDir(TestSerial, filesdb.TesseractDir)
	keyboardButtons := filesDB.GetFiles(keyboardDir)

	ocrResult := cvAPI.GetKeyboardKeys(keyboardButtons, img)

	var screenshotWithRects = filepath.Join(tesseractDir, "screenshot.png")
	var rectangles = []image.Rectangle{}
	for _, result := range ocrResult {
		rectangles = append(rectangles, *result.Rectangle.ToImageRectangle())
	}
	cvAPI.DrawRectangles(img, rectangles, false)
	if ok := gocv.IMWrite(screenshotWithRects, img); !ok {
		fmt.Println("could not write image " + screenshotWithRects)
	}
}

func TestDrawAllRectangles(t *testing.T) {
	var fileProps = &config.FilesProps{
		Logs: "./logs",
	}
	var logAPI = logger.New(logger.INFO, true)
	var filesDB = filesdb.New(fileProps)
	var cmdRunner = bashcmd.New(filesDB, logAPI)
	var cvAPI = cv.New(cmdRunner, logAPI)

	screenshot := cmdRunner.ScreenShot(TestSerial)
	if screenshot == "" {
		t.Fatal("screenshot is empty")
	}
	img := gocv.IMRead(screenshot, gocv.IMReadColor)
	if img.Empty() {
		t.Fatal("could not read screenshot image")
	}
	defer img.Close()

	rectangles, err := cvAPI.FindAllRectangles(&img)
	if err != nil {
		t.Fatal(err)
	}
	for _, rect := range rectangles {
		t.Log(rect.Max.Y)
	}

	if len(rectangles) == 0 {
		t.Fatal("no rectangles found")
	}

	t.Logf("rects len: %d", len(rectangles))
	err = cvAPI.DrawRectangles(img, rectangles, false)
	if err != nil {
		t.Fatal(err)
	}
	params := []int{gocv.IMWriteJpegQuality, 90}
	if ok := gocv.IMWriteWithParams(screenshot, img, params); !ok {
		fmt.Println("could not write image " + screenshot)
	}
}

func TestDrawSomeRectangle(t *testing.T) {
	var fileProps = &config.FilesProps{
		Logs: "./logs",
	}
	var logAPI = logger.New(logger.INFO, true)
	var filesDB = filesdb.New(fileProps)
	var cmdRunner = bashcmd.New(filesDB, logAPI)
	var cvAPI = cv.New(cmdRunner, logAPI)

	screenshot := cmdRunner.ScreenShot(TestSerial)
	if screenshot == "" {
		t.Fatal("screenshot is empty")
	}

	img := gocv.IMRead(screenshot, gocv.IMReadColor)
	if img.Empty() {
		t.Fatal("could not read screenshot image")
	}
	defer img.Close()

	rectangles, err := cvAPI.FindAllRectangles(&img)
	if err != nil {
		t.Fatal(err)
	}

	if len(rectangles) == 0 {
		t.Fatal("no rectangles found")
	}

	err = cvAPI.DrawRectangles(img, rectangles, false)
	if err != nil {
		t.Fatal(err)
	}
	params := []int{gocv.IMWriteJpegQuality, 90}
	if ok := gocv.IMWriteWithParams(screenshot, img, params); !ok {
		fmt.Println("could not write image " + screenshot)
	}
}

func TestFindTemplate(t *testing.T) {
	var fileProps = &config.FilesProps{
		Logs: "./logs",
	}
	var logAPI = logger.New(logger.INFO, true)
	var filesDB = filesdb.New(fileProps)
	var cmdRunner = bashcmd.New(filesDB, logAPI)
	var cvAPI = cv.New(cmdRunner, logAPI)

	screenshot := cmdRunner.ScreenShot(TestSerial)
	if screenshot == "" {
		t.Fatal("screenshot is empty")
	}

	img := gocv.IMRead(screenshot, gocv.IMReadColor)
	if img.Empty() {
		t.Fatal("could not read screenshot image")
	}
	defer img.Close()

	rectangle, err := cvAPI.FindImage(&img, TestImage)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rectangle)

	err = cvAPI.DrawRectangles(img, []image.Rectangle{*rectangle}, false)
	if err != nil {
		t.Fatal(err)
	}
	params := []int{gocv.IMWriteJpegQuality, 90}
	if ok := gocv.IMWriteWithParams(screenshot, img, params); !ok {
		fmt.Println("could not write image " + screenshot)
	}
}
