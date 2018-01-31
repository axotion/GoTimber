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

type Http struct {
	Method     string `json:"method,omitempty"`
	Path       string `json:"path,omitempty"`
	RemoteAddr string `json:"remote_addr,omitempty"`
	RequestID  string `json:"request_id,omitempty"`
}

type User struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type Context struct {
	Http *Http
	User *User
}
type goTimberJson struct {
	Dt      string      `json:"dt"`
	Level   string      `json:"level"`
	Message string      `json:"message"`
	Context *Context    `json:"context,omitempty"`
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

func (t *GoTimber) Warning(message string) *GoTimber {
	goTimberStruct := &goTimberJson{Dt: time.Now().String(), Level: _warning, Message: message}
	goTimberJSON, err := json.Marshal(goTimberStruct)
	checkError(err)
	t.client.postLog(goTimberJSON)
	return t
}

func (t *GoTimber) Debug(message string) *GoTimber {
	goTimberStruct := &goTimberJson{Dt: time.Now().String(), Level: _debug, Message: message}
	goTimberJSON, err := json.Marshal(goTimberStruct)
	checkError(err)
	t.client.postLog(goTimberJSON)
	return t
}

func (t *GoTimber) Info(message string) *GoTimber {
	goTimberStruct := &goTimberJson{Dt: time.Now().String(), Level: _info, Message: message}
	goTimberJSON, err := json.Marshal(goTimberStruct)
	checkError(err)
	t.client.postLog(goTimberJSON)
	return t
}

func (t *GoTimber) Notice(message string) *GoTimber {
	goTimberStruct := &goTimberJson{Dt: time.Now().String(), Level: _notice, Message: message}
	goTimberJSON, err := json.Marshal(goTimberStruct)
	checkError(err)
	t.client.postLog(goTimberJSON)
	return t
}

func (t *GoTimber) Error(message string) *GoTimber {
	goTimberStruct := &goTimberJson{Dt: time.Now().String(), Level: _error, Message: message}
	goTimberJSON, err := json.Marshal(goTimberStruct)
	checkError(err)
	t.client.postLog(goTimberJSON)
	return t
}

func (t *GoTimber) Critical(message string) *GoTimber {
	goTimberStruct := &goTimberJson{Dt: time.Now().String(), Level: _critical, Message: message}
	goTimberJSON, err := json.Marshal(goTimberStruct)
	checkError(err)
	t.client.postLog(goTimberJSON)
	return t
}

func (t *GoTimber) Alert(message string) *GoTimber {
	goTimberStruct := &goTimberJson{Dt: time.Now().String(), Level: _alert, Message: message}
	goTimberJSON, err := json.Marshal(goTimberStruct)
	checkError(err)
	t.client.postLog(goTimberJSON)
	return t
}

func (t *GoTimber) Emergency(message string) *GoTimber {
	goTimberStruct := &goTimberJson{Dt: time.Now().String(), Level: _emergency, Message: message}
	goTimberJSON, err := json.Marshal(goTimberStruct)
	checkError(err)
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
