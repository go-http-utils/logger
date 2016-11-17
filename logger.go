package logger

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Type represents logger's type
type Type int

const (
	// CombineLoggerType is the standard Apache combined log output
	// format: :remote-addr - :remote-user [:date[clf]] ":method :url HTTP/:http-version" :status :res[content-length] ":referrer" ":user-agent"
	CombineLoggerType Type = iota
	// CommonLoggerType is the standard Apache common log output
	// format: :remote-addr - :remote-user [:date[clf]] ":method :url HTTP/:http-version" :status :res[content-length]
	CommonLoggerType
	// DevLoggerType use colorful response status for development use
	// format: :method :url :status :response-time ms - :res[content-length]
	DevLoggerType
	// ShortLoggerType is shorter than default, also including response time
	// format: :remote-addr :remote-user :method :url HTTP/:http-version :status :res[content-length] - :response-time ms
	ShortLoggerType
	// TinyLoggerType is the minimal ouput
	// format: :method :url :status :res[content-length] - :response-time ms
	TinyLoggerType
)

const (
	combineFormat = `%s - %s "%s %s HTTP/%s" %v %v "%s" "%s"`
)

type loggerHanlder struct {
	h          http.Handler
	FormatType Type
	Dest       io.Writer
}

func (rh loggerHanlder) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	rl := &responseLogger{rw: res}

	rh.h.ServeHTTP(rl, req)

	go rh.WriteLog(rl, req)
}

func (rh loggerHanlder) WriteLog(rl *responseLogger, req *http.Request) {
	fmt.Fprintf(rh.Dest, combineFormat, req.RemoteAddr, "user", req.Method,
		req.RequestURI, "1.1", rl.status, rl.size, req.Referer(), req.UserAgent())
}

type responseLogger struct {
	rw     http.ResponseWriter
	status int
	size   int
}

func (rl *responseLogger) Header() http.Header {
	return rl.rw.Header()
}

func (rl *responseLogger) Write(bytes []byte) (int, error) {
	if rl.status == 0 {
		rl.status = http.StatusOK
	}

	rl.size = rl.size + len(bytes)

	return rl.rw.Write(bytes)
}

func (rl *responseLogger) WriteHeader(status int) {
	rl.status = status

	rl.rw.WriteHeader(status)
}

// Handler returns a http.Handler that wraps h by using
// Apache combined log output and the destination is os.Stdout
func Handler(h http.Handler) http.Handler {
	return loggerHanlder{
		h:          h,
		FormatType: CombineLoggerType,
		Dest:       os.Stdout,
	}
}
