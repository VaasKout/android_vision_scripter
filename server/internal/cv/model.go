package cv

import (
	"android_vision_scripter/pkg/models"
	"encoding/json"
	"fmt"
	"strings"
)

// Output files for cv
const (
	EdgesPng  = "edges.png"
	OcrJSON   = "ocr_results.json"
	OutputTsv = "output"
)

// OCRResult ...
type OCRResult struct {
	Text      string           `json:"text"`
	Rectangle models.Rectangle `json:"rectangle"`
}

// OCRJsonToArray ...
func OCRJsonToArray(ocrJSON []byte) []OCRResult {
	var ocrResult []OCRResult
	err := json.Unmarshal(ocrJSON, &ocrResult)
	if err != nil {
		fmt.Println(err)
		return []OCRResult{}
	}
	return ocrResult
}

// TsvToOCRResult ...
func TsvToOCRResult(data string) []OCRResult {
	var results []OCRResult
	lines := strings.Split(data, "\n")
	for i, line := range lines {
		if i == 0 || len(line) == 0 {
			continue // Skip header or empty lines
		}
		fields := strings.Split(line, "\t")
		if len(fields) < 12 {
			continue
		}

		// Extract values
		text := fields[11]
		if text == "" {
			continue
		}

		var x, y, width, height int
		fmt.Sscanf(fields[6], "%d", &x)
		fmt.Sscanf(fields[7], "%d", &y)
		fmt.Sscanf(fields[8], "%d", &width)
		fmt.Sscanf(fields[9], "%d", &height)

		results = append(results, OCRResult{
			Text: text,
			Rectangle: models.Rectangle{
				LeftX:   x,
				TopY:    y,
				RightX:  x + width,
				BottomY: y + height,
			},
		})
	}
	return results
}

// IsEmpty ...
func (o *OCRResult) IsEmpty() bool {
	return o == nil || o.Rectangle.IsEmpty() || o.Text == ""
}

// TemplateResult ...
type TemplateResult struct {
	Rectangle models.Rectangle
	Path      string
}
