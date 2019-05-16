# ymlog



### example

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
 
 	///console
 	logger1 := NewLogger(&ConsoleLoggerWriter{},
 	)
 	logger1.InfoString("init log")
 	
```