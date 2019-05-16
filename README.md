# ymlog



### example

```golang


   
logger := logger


	logger := &Logger{
		LogWriter: &FileLogWriter{
			ChanBufferLength: 10240,
			FileName:         "/pdata/log/test/example.log",
			DailyRotate:      true,
			RotateSize:       true,
			MaxSize:          512, // megabytes
			MaxBackup:        2,
			WriteFileBuffer:  10,
		},
	}
	logger.Init()

	logger.InfoString("init log")
```