package main

import "github.com/maczam/ymlog"

func main() {

	///file log
	logger := ymlog.NewLogger(&ymlog.FileLoggerWriter{
		ChanBufferLength: 10240,
		FileName:         "/pdata/log/test/example.log",
		RotateDaily:      true,
		RotateSize:       true,
		MaxSize:          512, // megabytes
		MaxBackup:        2,
		WriteFileBuffer:  10,
	},
	)
	logger.InfoString("init NewLogger log")

	///console
	logger1 := ymlog.NewLogger(&ymlog.ConsoleLoggerWriter{})
	logger1.InfoString("init ConsoleLoggerWriter log")
}
