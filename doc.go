// Package logger is a HTTP logger middleware for Go
//
// import (
//   "net/http"
// 	 "os"
//
//   "github.com/go-http-utils/logger"
// )
//
// mux := http.NewServeMux()
// mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
//   res.Write([]byte("Hello World"))
// })
//
// http.ListenAndServe(":8080", logger.Handler(mux, os.Stdout, logger.DevLoggerType))

package logger
