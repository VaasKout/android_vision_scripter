package models

import (
	"encoding/json"
	"fmt"
)

// Script actions
const (
	EventOnTemplate = "event_on_template"
	EventOnText     = "event_on_text"
	TypeText        = "type_text"
)

// Script ...
type Script struct {
	Name  string       `json:"name"`
	Steps []ScriptStep `json:"steps"`
}

// ToJSON ...
func (s *Script) ToJSON() []byte {
	result, err := json.Marshal(s)
	if err != nil {
		fmt.Println("Script ToJson " + err.Error())
		return []byte{}
	}
	return result
}

// ScriptStep ...
type ScriptStep struct {
	ID       int     `json:"id,omitempty"`
	Events   []Event `json:"events,omitempty"`
	Action   string  `json:"action,omitempty"`
	Template bool    `json:"template,omitempty"`
	Text     string  `json:"text,omitempty"`
	Command  string  `json:"command,omitempty"`
}

// Event ...
type Event struct {
	Time int64        `json:"time"`
	Data ControlBytes `json:"data"`
}
