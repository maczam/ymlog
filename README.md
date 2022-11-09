# ymlog [![Build Status](https://travis-ci.org/maczam/ymlog.svg?branch=master)](https://travis-ci.org/maczam/ymlog)&nbsp;[![GoDoc](https://godoc.org/github.com/maczam/ymlog?status.svg)](https://godoc.org/github.com/maczam/ymlog)

ymlog is golang high-performance asynchronous logging framework

## Installation

`go get github.com/maczam/ymlog`



## Quick Start

log file 

```golang

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
  	
```

 log console 
 ```
 	///console
 	logger1 := NewLogger(&ConsoleLoggerWriter{},
 	)
 	logger1.InfoString("init log")
  ```


Benchmark test file
``` 
    goos: windows
    goarch: amd64
    pkg: benchmarks
    cpu: AMD Ryzen 7 4800H with Radeon Graphics
    BenchmarkWithString
    BenchmarkWithString/Zap
    BenchmarkWithString/Zap-16                150902              8166 ns/op
    BenchmarkWithString/YmLog-1024
    BenchmarkWithString/YmLog-1024-16         841738              1231 ns/op
    BenchmarkWithString/YmLog-2048
    BenchmarkWithString/YmLog-2048-16        1000000              1008 ns/op
    PASS
```
