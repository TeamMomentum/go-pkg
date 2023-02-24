package logging

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	ltype "google.golang.org/genproto/googleapis/logging/type"
)

var once sync.Once

// logEntry represents Cloud Logging Structured logs
// https://cloud.google.com/logging/docs/structured-logging
type logEntry struct {
	Severity    string             `json:"severity,omitempty"`
	Message     string             `json:"message,omitempty"`
	HTTPRequest *ltype.HttpRequest `json:"httpRequest,omitempty"`
	Labels      map[string]string  `json:"logging.googleapis.com/labels,omitempty"`
}

func (le logEntry) String() string {
	if le.Severity == "" {
		le.Severity = ltype.LogSeverity_INFO.String()
	}

	out, err := json.Marshal(le)
	if err != nil {
		log.Printf("json.Marshal: %v", err)
	}

	return string(out)
}

type Entry struct {
	httpRequest *http.Request
	labels      map[string]string

	mu         sync.Mutex
	labelFuncs map[string]func(v string) string
}

func New() *Entry {
	once.Do(func() {
		log.SetFlags(0)
	})

	return &Entry{
		labels:     map[string]string{},
		labelFuncs: map[string]func(v string) string{},
	}
}

func (e *Entry) printf(serverity, format string, v ...interface{}) {
	l := make(map[string]string, len(e.labels))
	for k, v := range e.labels {
		f := e.labelFuncs[k]
		if f != nil {
			l[k] = f(v)

			continue
		}

		l[k] = v
	}

	sle := logEntry{
		Message:  fmt.Sprintf(format, v...),
		Labels:   l,
		Severity: serverity,
	}

	if r := e.httpRequest; r != nil {
		sle.HTTPRequest = &ltype.HttpRequest{
			RequestMethod: r.Method,
			RequestUrl:    r.URL.String(),
			Protocol:      r.Proto,
			RemoteIp:      r.RemoteAddr,
		}
	}

	log.Println(sle)
}

func (e *Entry) SetHTTPRequest(r *http.Request) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.httpRequest = r
}

func (e *Entry) SetLabel(k, v string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.labels[k] = v
}

func (e *Entry) SetLabels(l map[string]string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.labels = l
}

func (e *Entry) SetLabelFunc(k string, f func(v string) string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.labelFuncs[k] = f
}

func (e *Entry) Defaultf(s string, v ...interface{}) {
	e.printf(ltype.LogSeverity_DEFAULT.String(), s, v...)
}

func (e *Entry) Debugf(s string, v ...interface{}) {
	e.printf(ltype.LogSeverity_DEBUG.String(), s, v...)
}

func (e *Entry) Infof(s string, v ...interface{}) {
	e.printf(ltype.LogSeverity_INFO.String(), s, v...)
}

func (e *Entry) Notice(s string, v ...interface{}) {
	e.printf(ltype.LogSeverity_NOTICE.String(), s, v...)
}

func (e *Entry) Warningf(s string, v ...interface{}) {
	e.printf(ltype.LogSeverity_WARNING.String(), s, v...)
}

func (e *Entry) Errorf(s string, v ...interface{}) {
	e.printf(ltype.LogSeverity_ERROR.String(), s, v...)
}

func (e *Entry) Criticalf(s string, v ...interface{}) {
	e.printf(ltype.LogSeverity_CRITICAL.String(), s, v...)
}

func (e *Entry) Alertf(s string, v ...interface{}) {
	e.printf(ltype.LogSeverity_ALERT.String(), s, v...)
}

func (e *Entry) Emergencyf(s string, v ...interface{}) {
	e.printf(ltype.LogSeverity_EMERGENCY.String(), s, v...)
}
