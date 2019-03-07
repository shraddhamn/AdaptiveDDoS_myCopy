
package main
// from http://www.goinggo.net/2013/11/using-log-package-in-go.html

import (
	"io"
	"log"
	"os"
)

var (
	_DEBUG   *log.Logger
	_INFO    *log.Logger
	_WARNING *log.Logger
	_ERROR   *log.Logger
)

func initLog(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	_DEBUG = log.New(traceHandle,
		"BEGUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	_INFO = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	_WARNING = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	_ERROR = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}


func init() {
	f1, err := os.OpenFile("debug.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
	    log.Fatalf("Error opening log file: %v", err)
	}
	f2, err := os.OpenFile("stats.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
	    log.Fatalf("Error opening log file: %v", err)
	}
	initLog(f1, f2, f1, f1)
	//initLog(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
}