# logrus-raygun-hook
A Raygun.io hook for logrus

## Usage

```go
import (
  log "github.com/Sirupsen/logrus"
  "github.com/squirkle/logrus-raygun-hook"
)

func init() {
  hook, _ := raygun.NewHook("yourApiKey", "appName")
  log.AddHook(hook)

  // ...or, if you want to use the raygun4go Client directly
  client, _ := raygun4go.New(apiKey, appName)
  defer client.HandleError()
  log.AddHook(raygun.NewHookFromClient(client))
}
```

## Request Logging

The hook will look for a `http.Request` object in each log entry with the key of `request`. If found, it will pass this through to Raygun so all the request metadata is logged correctly.

#### Example

```go
  log.WithFields(log.Fields{
    "request" : someHttpRequest,
  }).Error("Some error occurred")
```

## Project status
This library is a **work in progress**. Be aware of the possibility of upcoming improvements/API changes.
