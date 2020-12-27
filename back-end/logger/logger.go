package logger

import (
	"go/build"
	"log"
	"os"
	"path/filepath"
)

var (
	//Log is used to export logger module
	Log *log.Logger
)

func init() {
	// set location of log file
	path := filepath.Join(build.Default.Dir, "logs", "info.log")
	dir := filepath.Dir(path)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
	file, err1 := os.Create(path)
	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + path)
}
