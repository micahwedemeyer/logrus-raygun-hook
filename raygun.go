package raygun

import (
	"net/http"
	"github.com/MindscapeHQ/raygun4go"
	"github.com/Sirupsen/logrus"
)

type raygunHook struct {
	Client *raygun4go.Client
}

func NewHook(apiKey string, appName string) (*raygunHook, error) {
	client,err := raygun4go.New(apiKey, appName)
	return NewHookFromClient(client),err
}

func NewHookFromClient(client *raygun4go.Client) *raygunHook {
	return &raygunHook{client}
}

func (hook *raygunHook) Fire(logEntry *logrus.Entry) error {
	hook.processEntry(logEntry)
	return hook.Client.CreateError(logEntry.Message)
}

func (hook *raygunHook) processEntry(logEntry *logrus.Entry) {
	if request, ok := logEntry.Data["request"].(*http.Request); ok {
		hook.Client.Request(request)
	}
	hook.Client.CustomData(logEntry.Data)
}

func (hook *raygunHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}
