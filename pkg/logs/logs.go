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

func InitLog(
	infoHandle io.Writer) {

	INFO = log.New(infoHandle,
		"[+] ",
		log.Ldate|log.Ltime)

	ERROR = log.New(infoHandle,
		"[-] ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func Set() *os.File {
	logInfo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Fail to open log file")
	}

	InitLog(logInfo)

	return logInfo
}
