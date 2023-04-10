package log

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var lgr *logger
var errlgr *errLogger

func NewConfig() Configure {
	return &config{
		Dir:  "logs",
		Path: ".",
	}
}

func Init(c Configure) error {
	config, ok := c.(*config)
	if !ok {
		return errors.New("invalid config type")
	}

	logDir := config.Path + "/" + config.Dir
	logFile := logDir + "/" + "process.log"
	errorLogFile := logDir + "/" + "error.log"

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return err
	}

	file, err := os.OpenFile(logFile, os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	lgr = &logger{
		logFile: logFile,
	}

	errorFile, err := os.OpenFile(errorLogFile, os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer errorFile.Close()

	errlgr = &errLogger{
		logFile: errorLogFile,
	}

	return nil
}

type Configure interface {
	SetPath(path string)
}

type config struct {
	Dir  string
	Path string
}

func (c *config) SetPath(path string) {
	c.Path = path
}

type logger struct {
	mutex     sync.Mutex
	logFile   string
	isEnabled bool
}

type errLogger struct {
	mutex     sync.Mutex
	logFile   string
	isEnabled bool
}

func Info(msg string) {
	lgr.mutex.Lock()
	defer lgr.mutex.Unlock()

	file, err := os.OpenFile(lgr.logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return
	}

	if fileInfo.Size() > 1024*1024*5 {
		divs := strings.Split(lgr.logFile, "/")
		fileName := divs[len(divs)-1]

		logDir := strings.Join(divs[:len(divs)-1], "/")

		newFilePath := logDir + "/" + time.Now().Format("2006_01_02_150405") + fileName

		err := os.Rename(lgr.logFile, newFilePath)
		if err != nil {
			return
		}

		_, err = os.Create(lgr.logFile)
		if err != nil {
			return
		}
	}

	message := time.Now().Format(time.RFC3339) + " " + msg

	file.WriteString(message + "\n")
}

func Error(msg string) {
	errlgr.mutex.Lock()
	defer errlgr.mutex.Unlock()

	file, err := os.OpenFile(errlgr.logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return
	}

	if fileInfo.Size() > 1024*1024*5 {
		divs := strings.Split(errlgr.logFile, "/")
		fileName := divs[len(divs)-1]

		logDir := strings.Join(divs[:len(divs)-1], "/")

		newFilePath := logDir + "/" + time.Now().Format("2006_01_02_150405") + fileName

		err := os.Rename(errlgr.logFile, newFilePath)
		if err != nil {
			return
		}

		_, err = os.Create(errlgr.logFile)
		if err != nil {
			return
		}
	}

	message := time.Now().Format(time.RFC3339) + " " + msg

	file.WriteString(message + "\n")
}

func Fatalln(err error) {
	fmt.Println(err)
	os.Exit(1)
}
