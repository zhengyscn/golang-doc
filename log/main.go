package main

import (
	"os"

	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("api")

var format = logging.MustStringFormatter(
	`[%{color}%{time:2006-01-02 15:04:05.999Z}] %{level:.4s} %{shortfile} %{id:03x}%{color:reset} %{message}`,
)

func outputFile() {
	fd, _ := os.OpenFile("/tmp/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer fd.Close()
	backend := logging.NewLogBackend(fd, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)

	logger.Debugf("debug %s", "123456")
	logger.Info("info")
	logger.Notice("notice")
	logger.Warning("warning")
	logger.Error("err")
	logger.Critical("crit")
}



func outputConsole() {
	backend := logging.NewLogBackend(os.Stdout, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)

	logger.Debugf("debug %s", "123456")
	logger.Info("info")
	logger.Notice("notice")
	logger.Warning("warning")
	logger.Error("err")
	logger.Critical("crit")
}


func main() {
  	outputFile()
	outputConsole()
  
}
