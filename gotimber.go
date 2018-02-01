package gotimber

import (
	"encoding/base64"
	"encoding/json"
	"time"
)

const (
	_warning   = "warning"
	_debug     = "debug"
	_info      = "info"
	_notice    = "notice"
	_error     = "error"
	_critical  = "critical"
	_alert     = "alert"
	_emergency = "emergency"
)

type GoTimber struct {
	apiKey string
	client *HttpClient
}

type goTimberJson struct {
	Dt      string      `json:"dt"`
	Level   string      `json:"level"`
	Message string      `json:"message"`
	Context interface{} `json:"context,omitempty"`
	Event   interface{} `json:"event,omitempty"`
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func NewGoTimber(apiKey string) *GoTimber {
	encodedAPIKey := base64.StdEncoding.EncodeToString([]byte(apiKey[:]))
	goTimber := &GoTimber{apiKey: encodedAPIKey}
	goTimber.client = &HttpClient{address: "https://logs.timber.io/frames", gotimber: goTimber}
	return goTimber
}

func prepareLogDataToSend(message string, level string) []byte {
	goTimberStruct := &goTimberJson{Dt: time.Now().String(), Level: level, Message: message}
	goTimberJSON, err := json.Marshal(goTimberStruct)
	checkError(err)
	return goTimberJSON
}

func (t *GoTimber) Warning(message string) *GoTimber {
	goTimberJSON := prepareLogDataToSend(message, _warning)
	t.client.postLog(goTimberJSON)
	return t
}

func (t *GoTimber) Debug(message string) *GoTimber {
	goTimberJSON := prepareLogDataToSend(message, _debug)
	t.client.postLog(goTimberJSON)
	return t
}

func (t *GoTimber) Info(message string) *GoTimber {
	goTimberJSON := prepareLogDataToSend(message, _info)
	t.client.postLog(goTimberJSON)
	return t
}

func (t *GoTimber) Notice(message string) *GoTimber {
	goTimberJSON := prepareLogDataToSend(message, _notice)
	t.client.postLog(goTimberJSON)
	return t
}

func (t *GoTimber) Error(message string) *GoTimber {
	goTimberJSON := prepareLogDataToSend(message, _error)
	t.client.postLog(goTimberJSON)
	return t
}

func (t *GoTimber) Critical(message string) *GoTimber {
	goTimberJSON := prepareLogDataToSend(message, _critical)
	t.client.postLog(goTimberJSON)
	return t
}

func (t *GoTimber) Alert(message string) *GoTimber {
	goTimberJSON := prepareLogDataToSend(message, _alert)
	t.client.postLog(goTimberJSON)
	return t
}

func (t *GoTimber) Emergency(message string) *GoTimber {
	goTimberJSON := prepareLogDataToSend(message, _emergency)
	t.client.postLog(goTimberJSON)
	return t
}

func (t *GoTimber) GetTimberStructure() *goTimberJson {
	return &goTimberJson{}
}

func (t *GoTimber) SendCustomTimberStructure(timberStructure *goTimberJson) {
	timberJson, err := json.Marshal(timberStructure)
	checkError(err)
	t.client.postLog(timberJson)
}
