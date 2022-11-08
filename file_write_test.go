package ymlog

import (
	"fmt"
	"testing"
	"time"
)

func TestRoteMin1(t *testing.T) {
	///file log
	logger := NewLogger(&FileLoggerWriter{
		ChanBufferLength: 10240,
		FileName:         "/pdata/log/test/example_%Y%M%D-%H%m.log.log",
		MaxSizeByteSize:  1024 * 1024,
		RotateDuration:   time.Second * 10,
		WriteFileBuffer:  10,
	},
	)
	logger.InfoString("init NewLogger log")

	///console
	logger1 := NewLogger(&ConsoleLoggerWriter{})
	logger1.InfoString("init ConsoleLoggerWriter log")

	for {
		time.Sleep(time.Second * 90)
		logger.InfoString(time.Now().String())
		logger1.InfoString(time.Now().String())
	}
}

func TestRoteMin2(t *testing.T) {
	///file log
	logger := NewLogger(&FileLoggerWriter{
		ChanBufferLength: 10240,
		FileName:         "/pdata/log/test/example_min2_%Y%M%D-%H%m.log",
		RotateDuration:   time.Second * 2,
		MaxSizeByteSize:  1024 * 1024,
		WriteFileBuffer:  10,
	},
	)
	logger.InfoString("init NewLogger log")

	for i := 0; i < 100; i++ {
		time.Sleep(time.Microsecond * 10)
		logger.InfoString(time.Now().String())
		logger.InfoString(time.Now().String())
		logger.InfoString(time.Now().String())
		logger.InfoString(time.Now().String())
		logger.InfoString(time.Now().String())
		logger.InfoString(time.Now().String())
	}

	fmt.Println("xxxxxxxxxxxxxxxxxx111111111111111")

	time.Sleep(time.Second * 3)

	fmt.Println("xxxxxxxxxxxxxxxxxx1111111111111112222")
	for i := 0; i < 100; i++ {
		time.Sleep(time.Microsecond * 10)
		logger.InfoString(time.Now().String())
		logger.InfoString(time.Now().String())
		logger.InfoString(time.Now().String())
	}
	fmt.Println("xxxxxxxxxxxxxxxxxx111111111111111333333")

	time.Sleep(time.Second * 3)
	fmt.Println("xxxxxxxxxxxxxxxxxx111111111111111444444")
	for i := 0; i < 100; i++ {
		time.Sleep(time.Microsecond * 10)
		logger.InfoString(time.Now().String())
		logger.InfoString(time.Now().String())
		logger.InfoString(time.Now().String())
	}
	fmt.Println("xxxxxxxxxxxxxxxxxx11111111111111155555")

	time.Sleep(time.Second * 100)

}
