package logging


//V0.1
import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File


	DefaultPrefix = ""
	DefaultCallerDepth = 2

	logger *log.Logger
	logPrefix = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	currentTime string
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	newLogger()
}

func newLogger(){
	filePath := getLogFileFullPath()
	fmt.Println(filepath)
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}



func check(){
	check := checkLogFile()

	if check{
		return
	}else{
		F.Close()
	}
	newLogger()

}

func Debug(v ...interface{}) {
	check()
	setPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {
	check()
	setPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	check()
	setPrefix(WARNING)
	log.Println(v)
	logger.Println(v)
}

func Error(v ...interface{}) {
	check()
	setPrefix(ERROR)
	log.Println(v)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	check()
	setPrefix(FATAL)
	log.Println(v)
	logger.Fatalln(v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)

	fmt.Println(file)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}