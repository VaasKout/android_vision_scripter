package usecases

import (
	"android_vision_scripter/internal/cv"
	"android_vision_scripter/internal/filesdb"
	"android_vision_scripter/pkg/core/file"
	"android_vision_scripter/pkg/models"
	"errors"
	"fmt"
	"image"
	"path/filepath"
	"time"
)

// Attempts to find template zone or text
const (
	Attempts = 3
)

func (i *interactorImpl) RunScript(serial string, scriptName string, socketPort int) error {
	if serial == "" {
		return errors.New(SerialIsEmptyError)
	}
	if scriptName == "" {
		return errors.New(ScriptNameIsEmpty)
	}

	path := i.getScriptRunner(serial, scriptName)
	if path == "" {
		return errors.New("script not found")
	}

	script := i.getScriptFromFile(path)
	if script == nil || len(script.Steps) == 0 {
		return errors.New("script steps are empty")
	}
	started := i.StartScrcpyServer(serial, socketPort)
	if !started {
		var errMsg = fmt.Sprintf("couldn't start scrcpy server for %s", serial)
		return errors.New(errMsg)
	}

	go i.startVideoListener(serial, script)
	return nil
}

func (i *interactorImpl) startVideoListener(
	serial string,
	script *models.Script,
) {
	var doneCh = make(chan struct{})
	defer i.logger.Info(
		fmt.Sprintf(
			"closing scrcpy connection for %s... 🛑",
			serial,
		),
	)

	go func() {
		defer close(doneCh)
		i.executeScript(serial, script)
	}()

	<-doneCh
	i.CloseConnection(serial)
}

func (i *interactorImpl) executeScript(
	serial string,
	script *models.Script,
) {
	var clientConnection *ClientConnection
	if result, ok := i.clientsCache.Get(serial); ok {
		clientConnection = &result
	}
	if clientConnection == nil || clientConnection.VideoPort == 0 {
		return
	}
	defer i.logger.Info("closing videostream... 🛑")
	go i.scrcpy.ReadVideoStream(serial, nil)

	for _, step := range script.Steps {
		stepZone, err := i.initStepZone(
			serial,
			script.Name,
			&step,
		)
		if err != nil {
			i.logger.Error(err.Error())
			return
		}

		var offsetX = 0
		var offsetY = 0
		var lastTimeStamp int64
		for index, event := range step.Events {
			var data = &event.Data
			if index == 0 && stepZone.IsNotEmpty() && step.Action != "" {
				offsetX, offsetY = i.countOffset(data, stepZone)
			}

			data.ApplyOffset(offsetX, offsetY)
			var delay = event.Time - lastTimeStamp
			time.Sleep(time.Duration(delay) * time.Millisecond)
			lastTimeStamp = event.Time

			i.scrcpy.WriteControlData(serial, *data)
		}
		time.Sleep(500 * time.Millisecond) //animation delay
	}
}

func (i *interactorImpl) countOffset(
	data *models.ControlBytes,
	stepZone *models.Rectangle,
) (int, int) {
	if stepZone.IsEmpty() || data == nil {
		return 0, 0
	}

	x, y := data.GetXY()
	randX, randY := stepZone.GetRandomXY()
	return randX - x, randY - y
}

func (i *interactorImpl) initStepZone(
	serial string,
	scriptName string,
	step *models.ScriptStep,
) (*models.Rectangle, error) {
	scriptDir := i.getScriptDir(serial, scriptName)

	tesseractDir := i.filesDB.CreateLogsDir(serial, filesdb.TesseractDir)
	if tesseractDir == "" {
		return &models.Rectangle{}, fmt.Errorf("tesseract dir was not found")
	}

	var imgRect = &image.Rectangle{}
	for range Attempts {
		var err error
		imgRect, err = i.findRectangleByStep(serial, step, scriptDir, tesseractDir)
		if err != nil {
			i.logger.Error(err.Error())
			time.Sleep(1 * time.Second)
			continue
		}

		break
	}

	stepZone := models.ImgRectangleToDomain(imgRect)
	if stepZone.IsEmpty() && step.Action != "" {
		return &models.Rectangle{}, fmt.Errorf(
			"zone not found in script: %s - ID:%d", scriptName, step.ID,
		)
	}

	return stepZone, nil
}

func (i *interactorImpl) findRectangleByStep(
	serial string,
	step *models.ScriptStep,
	scriptDir string,
	tesseractDir string,
) (*image.Rectangle, error) {
	var imgRect = &image.Rectangle{}

	mat, err := i.scrcpy.GetMatFromLastFrame(serial, true)
	if err != nil {
		return imgRect, err
	}
	if mat == nil {
		return imgRect, fmt.Errorf("mat is nil")
	}

	if step.Template {
		tmpImage := filepath.Join(scriptDir, fmt.Sprintf("%d.png", step.ID))
		if !file.Exists(tmpImage) {
			return imgRect, fmt.Errorf("template not found: %s", tmpImage)
		}

		templateRect, err := i.cv.FindImage(mat, tmpImage)
		if err != nil {
			return imgRect, err
		}

		if models.ImageRectIsEmpty(templateRect) {
			return imgRect, fmt.Errorf("template %d.png not found", step.ID)
		}

		if step.Action == models.EventOnTemplate {
			imgRect = templateRect
		}
	}

	if step.Text != "" {
		var device = i.GetDevice(serial)
		var ocrParams = cv.InitOcrParams(
			step.Text,
			device.Locale,
			cv.PsmText,
			cv.OemText,
			cv.WhiteTheme,
		)
		rectangles, err := i.cv.FindTextRectangles(
			mat,
			tesseractDir,
			ocrParams,
		)
		var textRect *image.Rectangle
		if len(rectangles) > 0 && err == nil {
			textRect = rectangles[0].Rectangle.ToImageRectangle()
		}

		if models.ImageRectIsEmpty(textRect) {
			return imgRect, fmt.Errorf("text %s not found", step.Text)
		}

		if step.Action == models.EventOnText {
			imgRect = textRect
		}
	}
	return imgRect, nil
}
