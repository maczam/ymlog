# ymlog [![Build Status](https://travis-ci.org/maczam/ymlog.svg?branch=master)](https://travis-ci.org/maczam/ymlog)&nbsp;[![GoDoc](https://godoc.org/github.com/maczam/ymlog?status.svg)](https://godoc.org/github.com/maczam/ymlog)

ymlog is golang high-performance asynchronous logging framework

## Installation

`go get github.com/maczam/ymlog`



## Quick Start

log file 

```golang

 	///file log
 	logger := NewLogger(&FileLoggerWriter{
 		ChanBufferLength: 10240,
 		FileName:         "/pdata/log/test/example.log",
 		RotateDaily:      true,
 		RotateSize:       true,
 		MaxSize:          512, // megabytes
 		MaxBackup:        2,
 		WriteFileBuffer:  10,
 	},
 	)
 	logger.InfoString("init log")
  	
```

 log console 
 ```
 	///console
 	logger1 := NewLogger(&ConsoleLoggerWriter{},
 	)
 	logger1.InfoString("init log")
  ```
