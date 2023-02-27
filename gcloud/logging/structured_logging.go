package logging

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"cloud.google.com/go/logging/apiv2/loggingpb"
	ltype "google.golang.org/genproto/googleapis/logging/type"
)

// entry represents Cloud Logging Structured logs
// https://cloud.google.com/logging/docs/structured-logging
type entry struct {
	Severity       string                            `json:"severity,omitempty"`
	Message        string                            `json:"message,omitempty"`
	HTTPRequest    *ltype.HttpRequest                `json:"httpRequest,omitempty"`
	InsertID       string                            `json:"logging.googleapis.com/insertId,omitempty"`
	Labels         map[string]string                 `json:"logging.googleapis.com/labels,omitempty"`
	Operation      *loggingpb.LogEntryOperation      `json:"logging.googleapis.com/operation,omitempty"`
	SourceLocation *loggingpb.LogEntrySourceLocation `json:"logging.googleapis.com/sourceLocation,omitempty"`
	SpanID         string                            `json:"logging.googleapis.com/spanId,omitempty"`
	Trace          string                            `json:"logging.googleapis.com/trace,omitempty"`
	TraceSampled   bool                              `json:"logging.googleapis.com/trace_sampled,omitempty"`
}

func (e *entry) String() string {
	if e.Severity == "" {
		e.Severity = ltype.LogSeverity_INFO.String()
	}

	out, err := json.Marshal(e)
	if err != nil {
		log.Printf("json.Marshal: %v", err)
	}

	return string(out)
}

type Logger struct {
	logger *log.Logger
	entry  *entry

	labelFuncs map[string]func(v string) string
	mu         sync.Mutex
}

func New(out io.Writer, prefix string, flag int) *Logger {
	return &Logger{
		logger: log.New(out, prefix, flag),
		entry: &entry{
			Labels: map[string]string{},
		},
		labelFuncs: map[string]func(v string) string{},
	}
}

func (lg *Logger) printf(serverity, format string, v ...interface{}) {
	lg.mu.Lock()
	defer lg.mu.Unlock()

	lg.logger.Println(lg.fmtEntry(serverity, fmt.Sprintf(format, v...)))
}

func (lg *Logger) println(serverity string, v ...interface{}) {
	lg.mu.Lock()
	defer lg.mu.Unlock()

	lg.logger.Println(lg.fmtEntry(serverity, fmt.Sprint(v...)))
}

func (lg *Logger) fmtEntry(serverity, msg string) *entry {
	l := make(map[string]string, len(lg.entry.Labels))
	for k, v := range lg.entry.Labels {
		f := lg.labelFuncs[k]
		if f != nil {
			l[k] = f(v)

			continue
		}

		l[k] = v
	}

	return &entry{
		Message:  msg,
		Labels:   l,
		Severity: serverity,
	}
}

func (lg *Logger) SetOutput(out io.Writer) {
	lg.mu.Lock()
	defer lg.mu.Unlock()

	lg.logger.SetOutput(out)
}

func (lg *Logger) SetPrefix(prefix string) {
	lg.mu.Lock()
	defer lg.mu.Unlock()

	lg.logger.SetPrefix(prefix)
}

func (lg *Logger) SetFlags(flag int) {
	lg.mu.Lock()
	defer lg.mu.Unlock()

	lg.logger.SetFlags(flag)
}

func (lg *Logger) SetEntryHTTPRequest(r *http.Request) {
	lg.mu.Lock()
	defer lg.mu.Unlock()

	if r != nil {
		lg.entry.HTTPRequest = &ltype.HttpRequest{
			RequestMethod: r.Method,
			RequestUrl:    r.URL.String(),
			Protocol:      r.Proto,
			RemoteIp:      r.RemoteAddr,
		}
	}
}

func (lg *Logger) SetEntryInsertID(id string) {
	lg.mu.Lock()
	defer lg.mu.Unlock()

	lg.entry.InsertID = id
}

func (lg *Logger) SetEntryLabel(k, v string) {
	lg.mu.Lock()
	defer lg.mu.Unlock()

	lg.entry.Labels[k] = v
}

func (lg *Logger) SetEntryLabels(l map[string]string) {
	lg.mu.Lock()
	defer lg.mu.Unlock()

	lg.entry.Labels = l
}

func (lg *Logger) SetEntryLabelFunc(k string, f func(v string) string) {
	lg.mu.Lock()
	defer lg.mu.Unlock()

	lg.labelFuncs[k] = f
}

func (lg *Logger) SetEntryOperation(o *loggingpb.LogEntryOperation) {
	lg.mu.Lock()
	defer lg.mu.Unlock()

	lg.entry.Operation = o
}

func (lg *Logger) SetEntrySourceLocation(sl *loggingpb.LogEntrySourceLocation) {
	lg.mu.Lock()
	defer lg.mu.Unlock()

	lg.entry.SourceLocation = sl
}

func (lg *Logger) SetEntrySpanID(id string) {
	lg.mu.Lock()
	defer lg.mu.Unlock()

	lg.entry.SpanID = id
}

func (lg *Logger) SetEntryTrace(t string) {
	lg.mu.Lock()
	defer lg.mu.Unlock()

	lg.entry.Trace = t
}

func (lg *Logger) SetEntryTraceSampled(t bool) {
	lg.mu.Lock()
	defer lg.mu.Unlock()

	lg.entry.TraceSampled = t
}

func (lg *Logger) Defaultf(format string, v ...interface{}) {
	lg.printf(ltype.LogSeverity_DEFAULT.String(), format, v...)
}

func (lg *Logger) Debugf(format string, v ...interface{}) {
	lg.printf(ltype.LogSeverity_DEBUG.String(), format, v...)
}

func (lg *Logger) Printf(format string, v ...interface{}) {
	lg.printf(ltype.LogSeverity_INFO.String(), format, v...)
}

func (lg *Logger) Noticef(format string, v ...interface{}) {
	lg.printf(ltype.LogSeverity_NOTICE.String(), format, v...)
}

func (lg *Logger) Warningf(format string, v ...interface{}) {
	lg.printf(ltype.LogSeverity_WARNING.String(), format, v...)
}

func (lg *Logger) Errorf(format string, v ...interface{}) {
	lg.printf(ltype.LogSeverity_ERROR.String(), format, v...)
}

func (lg *Logger) Criticalf(format string, v ...interface{}) {
	lg.printf(ltype.LogSeverity_CRITICAL.String(), format, v...)
}

func (lg *Logger) Alertf(format string, v ...interface{}) {
	lg.printf(ltype.LogSeverity_ALERT.String(), format, v...)
}

func (lg *Logger) Emergencyf(format string, v ...interface{}) {
	lg.printf(ltype.LogSeverity_EMERGENCY.String(), format, v...)
}

func (lg *Logger) Defaultln(v ...interface{}) {
	lg.println(ltype.LogSeverity_DEFAULT.String(), v...)
}

func (lg *Logger) Debugln(v ...interface{}) {
	lg.println(ltype.LogSeverity_DEBUG.String(), v...)
}

func (lg *Logger) Println(v ...interface{}) {
	lg.println(ltype.LogSeverity_INFO.String(), v...)
}

func (lg *Logger) Noticeln(v ...interface{}) {
	lg.println(ltype.LogSeverity_NOTICE.String(), v...)
}

func (lg *Logger) Warningln(v ...interface{}) {
	lg.println(ltype.LogSeverity_WARNING.String(), v...)
}

func (lg *Logger) Errorln(v ...interface{}) {
	lg.println(ltype.LogSeverity_ERROR.String(), v...)
}

func (lg *Logger) Criticalln(v ...interface{}) {
	lg.println(ltype.LogSeverity_CRITICAL.String(), v...)
}

func (lg *Logger) Alertln(v ...interface{}) {
	lg.println(ltype.LogSeverity_ALERT.String(), v...)
}

func (lg *Logger) Emergencyln(v ...interface{}) {
	lg.println(ltype.LogSeverity_EMERGENCY.String(), v...)
}

var std = New(os.Stderr, "", 0)

func Default() *Logger { return std }

func SetOutput(out io.Writer) {
	std.SetOutput(out)
}

func SetPrefix(prefix string) {
	std.SetPrefix(prefix)
}

func SetFlags(flag int) {
	std.SetFlags(flag)
}

func SetEntryHTTPRequest(r *http.Request) {
	std.SetEntryHTTPRequest(r)
}

func SetEntryInsertID(id string) {
	std.SetEntryInsertID(id)
}

func SetEntryLabel(k, v string) {
	std.SetEntryLabel(k, v)
}

func SetEntryLabels(l map[string]string) {
	std.SetEntryLabels(l)
}

func SetEntryLabelFunc(k string, f func(v string) string) {
	std.SetEntryLabelFunc(k, f)
}

func SetEntryOperation(o *loggingpb.LogEntryOperation) {
	std.SetEntryOperation(o)
}

func SetEntrySourceLocation(sl *loggingpb.LogEntrySourceLocation) {
	std.SetEntrySourceLocation(sl)
}

func SetEntrySpanID(id string) {
	std.SetEntrySpanID(id)
}

func SetEntryTrace(t string) {
	std.SetEntryTrace(t)
}

func SetEntryTraceSampled(t bool) {
	std.SetEntryTraceSampled(t)
}

func Defaultf(format string, v ...interface{}) {
	std.Defaultf(format, v...)
}

func Debugf(format string, v ...interface{}) {
	std.Debugf(format, v...)
}

func Printf(format string, v ...interface{}) {
	std.Printf(format, v...)
}

func Noticef(format string, v ...interface{}) {
	std.Noticef(format, v...)
}

func Warningf(format string, v ...interface{}) {
	std.Warningf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	std.Errorf(format, v...)
}

func Criticalf(format string, v ...interface{}) {
	std.Criticalf(format, v...)
}

func Alertf(format string, v ...interface{}) {
	std.Alertf(format, v...)
}

func Emergencyf(format string, v ...interface{}) {
	std.Emergencyf(format, v...)
}

func Defaultln(v ...interface{}) {
	std.Defaultln(v...)
}

func Debugln(v ...interface{}) {
	std.Debugln(v...)
}

func Println(v ...interface{}) {
	std.Println(v...)
}

func Noticeln(v ...interface{}) {
	std.Noticeln(v...)
}

func Warningln(v ...interface{}) {
	std.Warningln(v...)
}

func Errorln(v ...interface{}) {
	std.Errorln(v...)
}

func Criticalln(v ...interface{}) {
	std.Criticalln(v...)
}

func Alertln(v ...interface{}) {
	std.Alertln(v...)
}

func Emergencyln(v ...interface{}) {
	std.Emergencyln(v...)
}
