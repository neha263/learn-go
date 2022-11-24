package main

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	//open a specified file and writes log into that
	// file, err := os.OpenFile("logfile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	// if err != nil {
	// 	log.Fatal("error occured while opening file", err.Error())
	// }
	// log.SetOutput(file)

	file, err := os.OpenFile("logfile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal("error occured while opening file", err.Error())
	}
	wrt := io.MultiWriter(os.Stdout, file)
	log.SetOutput(wrt)

	// log.SetOutput(os.Stdout)
	log.SetReportCaller(false)
	log.SetLevel(log.DebugLevel)
	log.Info("this is info log")
	log.Error("this is error log")
	log.Debug("this is debug log")
	log.Warn("this is worning log")
	log.Fatal("this is fatel log")
	log.Trace("this is trace log")
}
