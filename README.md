# GoTimber

Client for [timber.io](https://timber.io/) written in pure GO. See example for more information

## How to install

```
go get github.com/axotion/GoTimber
```

## Example usage 

```go
timber := gotimber.NewGoTimber("timber-token-api")

	//Send only message
	go timber.Info("Async info test")
	go timber.Alert("Async alert test")
	
	//Set manually fields for Context
	timberStructure := timber.GetTimberStructure()
	timberStructure.Dt = time.Now().String()
	timberStructure.Level = "warning"
	timberStructure.Message = "Hello World!"

	user := &gotimber.User{Email: "test@test.com", ID: 1, Name: "John SMITH"}
	http := &gotimber.Http{Method: "POST", Path: "/test", RemoteAddr: "127.0.0.1", RequestID: "123432"}
	timberStructure.Context = &gotimber.Context{Http: http, User: user}

	//Custom event
	timberStructure.Event = map[string]map[string]string{"payment": map[string]string{"id": "1", "success": "ok"}, "order": map[string]string{"id": "2"}}

	go timber.SendCustomTimberStructure(timberStructure)
    time.Sleep(time.Second)
    

