package loghelper

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	// Info : Discard
	Info *log.Logger
	// Warning : Stdout
	Warning *log.Logger
	// Error : Stderr
	Error *log.Logger
	// Login .
	Login *log.Logger
)

var errlog *os.File
var loginlog *os.File

func set(
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer,
	loginHandle io.Writer) {

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Login = log.New(loginHandle,
		"LOG: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func init() {
	errlog = getErrLogFile()
	loginlog = getLogFile()
	set(ioutil.Discard, os.Stdout, errlog, loginlog)

	//Info.Println("Special Information")
	//Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
	Login.Println("There is the log message")
}

// Free : close log file
func Free() {
	errlog.Close()
	loginlog.Close()
}

func getErrLogFile() *os.File {
	//logPath := filepath.Join(os.Getenv("GOPATH"), "/src/Agenda/data/error.log")
	logPath := "data/error.log"
	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file open error : %v", err)
	}
	return file
	// defer file.Close()
}

func getLogFile() *os.File {
	logPath := "data/login.log"
	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file open error : %v", err)
	}
	return file
}