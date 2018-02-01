# GoTimber

Client for [timber.io](https://timber.io/) written in pure GO. See example for more information

## How to install

```
go get github.com/axotion/GoTimber
```
## Effect of example usage
![Screen of ](https://imgur.com/MagefjBl.png)

## Example usage 

```go
package main

import (
	"gotimber/gotimber"
	"time"
)

func main() {
	timber := gotimber.NewGoTimber("2142_6ed835c3e5a9354a:8480d4fab4c9288bd96d476a0f253f0aa8ab8eba476d2b4c35ed1722e53aab5a")

	//Send only message
	go timber.Info("Async info test")
	go timber.Alert("Async alert test")

	//Set manually fields for Context
	timberStructure := timber.GetTimberStructure()
	timberStructure.Dt = time.Now().String()
	timberStructure.Level = "warning"
	timberStructure.Message = "Hello World!"

	//Custom context #1 way (or just create your own struct)

	httpContext := struct {
		UserID int `json:"user_id"` 
	} { UserID: 1}

	timberStructure.Context = struct {
		HTTP interface{} `json:"http"`
	}{ HTTP: httpContext}

	//Custom event #2 way
	timberStructure.Event = map[string]map[string]string{"payment": map[string]string{"id": "1", "success": "ok"}, "order": map[string]string{"id": "2"}}

	go timber.SendCustomTimberStructure(timberStructure)
	time.Sleep(time.Second)
}



    

