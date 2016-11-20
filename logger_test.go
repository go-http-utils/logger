package logger

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type LoggerSuite struct {
	suite.Suite

	req *http.Request
	rl  *responseLogger
	w   *testWriter
}

func (s *LoggerSuite) SetupTest() {
	s.req = httptest.NewRequest(http.MethodGet, "/", nil)
	s.req.RemoteAddr = "192.0.2.1:1234"

	s.rl = &responseLogger{rw: testResponseWriter{}, start: time.Now()}
	s.w = &testWriter{}

	s.rl.Write([]byte("test-logger"))
}

func (s *LoggerSuite) TestTiny() {
	lh := loggerHanlder{
		h:          http.NotFoundHandler(),
		formatType: TinyLoggerType,
		writer:     s.w,
	}
	lh.write(s.rl, s.req)

	s.Equal("GET / 200 11 - 0.000 ms\n", string(s.w.Bytes))
}

func (s *LoggerSuite) TestShort() {
	lh := loggerHanlder{
		h:          http.NotFoundHandler(),
		formatType: ShortLoggerType,
		writer:     s.w,
	}
	lh.write(s.rl, s.req)

	s.Equal("192.0.2.1:1234 - GET / HTTP/1.1 200 11 - 0.000 ms\n", string(s.w.Bytes))
}

func (s *LoggerSuite) TestDev() {
	lh := loggerHanlder{
		h:          http.NotFoundHandler(),
		formatType: DevLoggerType,
		writer:     s.w,
	}
	lh.write(s.rl, s.req)

	s.Equal("GET / 200 0.000 ms - 11\n", string(s.w.Bytes))
}

func (s *LoggerSuite) TestCommon() {
	lh := loggerHanlder{
		h:          http.NotFoundHandler(),
		formatType: CommonLoggerType,
		writer:     s.w,
	}
	lh.write(s.rl, s.req)

	s.Equal(`192.0.2.1:1234 - - [`+s.rl.start.Format(timeFormat)+`] "GET / HTTP/1.1" 200 11`+"\n", string(s.w.Bytes))
}

func (s *LoggerSuite) TestCombined() {
	lh := loggerHanlder{
		h:          http.NotFoundHandler(),
		formatType: CombineLoggerType,
		writer:     s.w,
	}
	lh.write(s.rl, s.req)

	s.Equal(`192.0.2.1:1234 - - [`+s.rl.start.Format(timeFormat)+`] "GET / HTTP/1.1" 200 11 "" ""`+"\n", string(s.w.Bytes))
}

func TestLogger(t *testing.T) {
	suite.Run(t, new(LoggerSuite))
}

type testResponseWriter struct {
	header http.Header
}

func (trw testResponseWriter) Header() http.Header {
	if trw.header != nil {
		return trw.header
	}

	return make(http.Header)
}

func (trw testResponseWriter) Write(bytes []byte) (int, error) {
	return len(bytes), nil
}

func (trw testResponseWriter) WriteHeader(status int) {}

type testWriter struct {
	Bytes []byte
}

func (tw *testWriter) Write(b []byte) (n int, err error) {
	tw.Bytes = append(tw.Bytes, b...)

	return len(b), nil
}
