package main

import (
	"github.com/maczam/ymlog"
	"time"
)

func main() {
	///file log
	logger := ymlog.NewLogger(&ymlog.FileLoggerWriter{
		FileName:         "/pdata/log/test/3m%Y%M%D-%H%m.log",
		MaxSizeByteSize:  1024 * 1024 * 3,
		RotateDuration:   time.Minute,
		ChanBufferLength: 1024,
		WriteFileBuffer:  1024,
	},
	)
	logger.InfoString("init NewLogger log")

	///console
	logger1 := ymlog.NewLogger(&ymlog.ConsoleLoggerWriter{})
	logger1.InfoString("init ConsoleLoggerWriter log")

	for {
		time.Sleep(time.Second * 90)
		logger.InfoString(time.Now().String())
		logger1.InfoString(time.Now().String())
	}
}
