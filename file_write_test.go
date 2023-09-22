package ymlog

import (
	"os"
	"testing"
	"time"
)

func TestRoteMin1(t *testing.T) {
	///file log
	logger := NewLogger(&FileLoggerWriter{
		ChanBufferLength: 10240,
		FileName:         "/pdata/log/test/example_%Y%M%D-%H%m.log.log",
		MaxSizeByteSize:  1024 * 1024,
		//RotateType:
		//RotateDuration:   time.Second * 10,
		WriteFileBuffer: 10,
	},
	)
	logger.InfoString("init NewLogger log")

	///console
	logger1 := NewLogger(&OutLoggerWriter{
		Out: os.Stdout,
	})
	logger1.InfoString("init OutLoggerWriter log")

	for {
		time.Sleep(time.Second * 90)
		logger.InfoString(time.Now().String())
		logger1.InfoString(time.Now().String())
	}
}

func TestRoteMin2(t *testing.T) {
	///file log
	logger := NewLogger(&FileLoggerWriter{
		ChanBufferLength: 1023,
		FileName:         "/Users/hexin/Documents/temproot/example_min2_%Y%M%D-%H%m.log",
		//RotateDuration:   time.Second * 2,
		RotateType:      RotateMinute,
		MaxSizeByteSize: 1024 * 1024,
		WriteFileBuffer: 10,
	},
	)
	logger.InfoString("init NewLogger log")

	for {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Microsecond * 10)
			logger.InfoString(time.Now().String())
			logger.InfoString(time.Now().String())
			logger.InfoString(time.Now().String())
			logger.InfoString(time.Now().String())
			logger.InfoString(time.Now().String())
			logger.InfoString(time.Now().String())
		}
	}
}
