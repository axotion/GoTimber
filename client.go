package gotimber

import (
	"bytes"
	"net/http"
	"time"
)

type HttpClient struct {
	address  string
	gotimber *GoTimber
}

func (c HttpClient) postLog(jsonLog []byte) {
	client := &http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest("POST", c.address, bytes.NewBuffer(jsonLog))
	checkError(err)
	req.Header.Add("authorization", "Basic "+c.gotimber.apiKey)
	req.Header.Add("content-type", "text/plain")
	res, err := client.Do(req)
	if err != nil || res.StatusCode == 500 {
		time.Sleep(time.Second * 15)
		c.postLog(jsonLog)
	}
}
