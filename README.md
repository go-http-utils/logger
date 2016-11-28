# logger
[![Build Status](https://travis-ci.org/go-http-utils/logger.svg?branch=master)](https://travis-ci.org/go-http-utils/logger)
[![Coverage Status](https://coveralls.io/repos/github/go-http-utils/logger/badge.svg?branch=master)](https://coveralls.io/github/go-http-utils/logger?branch=master)

HTTP logger middleware for Go

## Installation

```sh
go get -u github.com/go-http-utils/logger
```

## Documentation

https://godoc.org/github.com/go-http-utils/logger

## Usage

```go
import (
  "net/http"
  "os"

  "github.com/go-http-utils/logger"
)

mux := http.NewServeMux()
mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
  res.Write([]byte("Hello World"))
})

http.ListenAndServe(":8080", logger.Handler(mux, os.Stdout, logger.DevLoggerType))
```

## Supportted log output format

### CombineLoggerType

CombineLoggerType is the standard Apache combined log output

```
:remote-addr - :remote-user [:date[clf]] ":method :url HTTP/:http-version" :status :res[content-length] ":referrer" ":user-agent"
```

### CommonLoggerType

CommonLoggerType is the standard Apache common log output

```
:remote-addr - :remote-user [:date[clf]] ":method :url HTTP/:http-version" :status :res[content-length]
```

### DevLoggerType

DevLoggerType is useful for development

```
:method :url :status :response-time ms - :res[content-length]
```

### ShortLoggerType

ShortLoggerType is shorter than common, including response time

```
:remote-addr :remote-user :method :url HTTP/:http-version :status :res[content-length] - :response-time ms
```

### TinyLoggerType

TinyLoggerType is the minimal ouput

```
:method :url :status :res[content-length] - :response-time ms
```
