package logs

import (
	"io"
	"log"
	"os"
)

var (
	INFO  *log.Logger
	ERROR *log.Logger
)

func initLog(
	infoHandle io.Writer) {

	INFO = log.New(infoHandle,
		"[+] ",
		log.Ldate|log.Ltime)

	ERROR = log.New(infoHandle,
		"[-] ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// Set creates log file and sets writer
func Set() *os.File {
	logInfo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Fail to open log file")
	}

	initLog(logInfo)

	return logInfo
}
