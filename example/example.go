package main

import (
	"github.com/maczam/ymlog"
	"os"
	"time"
)

func main() {
	///file log
	logger := ymlog.NewLogger(&ymlog.FileLoggerWriter{
		FileName:         "./logs/error.log",
		MaxSizeByteSize:  1024 * 1024 * 3,
		ChanBufferLength: 1024,
		WriteFileBuffer:  1024,
	},
	)
	logger2 := ymlog.NewLogger(&ymlog.FileLoggerWriter{
		FileName:         "./logs/error_%Y-%M-%D:%H%m%s.log",
		RotateType:       ymlog.RotateMinute,
		MaxSizeByteSize:  1024 * 1024 * 3,
		ChanBufferLength: 1024,
		WriteFileBuffer:  1024,
	},
	)
	logger.InfoString("init NewLogger log")

	///console
	logger1 := ymlog.NewLogger(&ymlog.OutLoggerWriter{
		Out: os.Stdout,
	})
	logger1.InfoString("init ConsoleLoggerWriter log")

	for {
		time.Sleep(time.Second * 1)
		//logger.InfoString(time.Now().String())
		logger1.InfoString(time.Now().String())
		logger2.InfoString(time.Now().String())
	}
}
