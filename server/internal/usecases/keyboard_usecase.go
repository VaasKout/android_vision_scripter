package usecases

import (
	"android_vision_scripter/internal/cv"
	"android_vision_scripter/internal/filesdb"
	"android_vision_scripter/pkg/models"
	"fmt"
	"path/filepath"

	"gocv.io/x/gocv"
)

// KeyboardUseCase ...
type KeyboardUseCase interface {
	GetKeyboardKeys(serial string, locale string) []cv.OCRResult
	EditKeyboardKey(
		serial string,
		locale string,
		name string,
		rectangle *models.Rectangle,
	) bool
	ResetKeyboardKeys(serial string, locale string, upperCase bool) []cv.OCRResult
}

func (i *interactorImpl) GetKeyboardKeys(
	serial string,
	locale string,
) []cv.OCRResult {
	screenshot := i.cmd.ScreenShot(serial)
	if screenshot == "" {
		return []cv.OCRResult{}
	}

	img := gocv.IMRead(screenshot, gocv.IMReadColor)
	if img.Empty() {
		return []cv.OCRResult{}
	}
	defer img.Close()

	var device = i.GetDevice(serial)
	var modelOs = device.ToModelOs()
	if modelOs == "" {
		return []cv.OCRResult{}
	}

	var tesseractLocale = cv.DefaultOCRLanguage
	if result, ok := cv.TesseractLocaleMap[locale]; ok {
		tesseractLocale = result
	}

	keyboardDir := i.filesDB.CreateDBDir(modelOs, filesdb.Keyboards, tesseractLocale)
	keyboardButtons := i.filesDB.GetFiles(keyboardDir)
	if len(keyboardButtons) == 0 {
		return []cv.OCRResult{}
	}
	return i.cv.GetKeyboardKeys(keyboardButtons, img)
}

func (i *interactorImpl) EditKeyboardKey(
	serial string,
	locale string,
	name string,
	rectangle *models.Rectangle,
) bool {
	if serial == "" || name == "" || rectangle.IsEmpty() {
		return false
	}

	screenshot := i.cmd.ScreenShot(serial)
	if screenshot == "" {
		return false
	}

	var device = i.GetDevice(serial)
	var modelOs = device.ToModelOs()
	if modelOs == "" {
		return false
	}

	var imgRect = rectangle.ToImageRectangle()
	var tesseractLocale = cv.DefaultOCRLanguage
	if result, ok := cv.TesseractLocaleMap[locale]; ok {
		tesseractLocale = result
	}

	keyboardDir := i.filesDB.CreateDBDir(modelOs, filesdb.Keyboards, tesseractLocale)
	var keyboardKeyPath = filepath.Join(keyboardDir, fmt.Sprintf("%s.png", name))
	i.cv.CutZone(screenshot, keyboardKeyPath, imgRect)

	return true
}

func (i *interactorImpl) ResetKeyboardKeys(
	serial string,
	locale string,
	upperCase bool,
) []cv.OCRResult {
	screenshot := i.cmd.ScreenShot(serial)
	if screenshot == "" {
		return []cv.OCRResult{}
	}

	var device = i.GetDevice(serial)
	var modelOs = device.ToModelOs()
	if modelOs == "" {
		return []cv.OCRResult{}
	}

	var tesseractLocale = cv.DefaultOCRLanguage
	if result, ok := cv.TesseractLocaleMap[locale]; ok {
		tesseractLocale = result
	}

	i.filesDB.DeletePathInDBDir(modelOs, filesdb.Keyboards, tesseractLocale)
	keyboardDir := i.filesDB.CreateDBDir(modelOs, filesdb.Keyboards, tesseractLocale)
	tesseractDir := i.filesDB.CreateLogsDir(serial, filesdb.TesseractDir)

	return i.cv.ResetKeyboardKeys(
		keyboardDir,
		tesseractDir,
		screenshot,
		tesseractLocale,
		upperCase,
	)
}
