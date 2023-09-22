package ymlog

import (
	"testing"
)

func TestNewLogger1(t *testing.T) {

	common := NewLogger(&FileLoggerWriter{
		FileName:        "/pdata/log/test/3m%Y%M%D-%H%m.log",
		MaxSizeByteSize: 1024 * 1024 * 3,
	})

	for {
		common.InfoString(`packageymlogimport("testing""time")funcTestNewLogger1(t*testing.T){common:=NewLogger(&FileLoggerWriter{FileName:"/pdata/3m.log",RotateDuration:time.Minute*1,MaxSize:3,//megabytesMaxBackup:3,})for{common.InfoString("")}}funcTest100M(t*testing.T){common:=NewLogger(&FileLoggerWriter{ChanBufferLength:1024*10,FileName:"/pdata/ymlog/100m_%Y%M%D%H_%m%s.log",RotateDuration:time.Minute*1,MaxSize:512,//megabytesMaxBackup:50,WriteFileBuffer:50,})for{common.InfoString()}}funcTest50M(t*testing.T){common:=NewLogger(&FileLoggerWriter{ChanBufferLength:1024*10,FileName:"/pdata/ymlog/50m_%Y%M%D%H.log",RotateDuration:time.Minute*1,MaxSize:50,//megabytesMaxBackup:50,WriteFileBuffer:50,})for{common.InfoString()}}`)
	}
}

func TestNewLoggerLessOneMicrosecond(t *testing.T) {

	common := NewLogger(&FileLoggerWriter{
		FileName:        "/pdata/log/test/3m%Y%M%D-%H%m.log",
		MaxSizeByteSize: 1024 * 1024 * 3,
	})

	for {
		common.InfoString(`packageymlogimport("testing""time")funcTestNewLogger1(t*testing.T){common:=NewLogger(&FileLoggerWriter{FileName:"/pdata/3m.log",RotateDuration:time.Minute*1,MaxSize:3,//megabytesMaxBackup:3,})for{common.InfoString("")}}funcTest100M(t*testing.T){common:=NewLogger(&FileLoggerWriter{ChanBufferLength:1024*10,FileName:"/pdata/ymlog/100m_%Y%M%D%H_%m%s.log",RotateDuration:time.Minute*1,MaxSize:512,//megabytesMaxBackup:50,WriteFileBuffer:50,})for{common.InfoString()}}funcTest50M(t*testing.T){common:=NewLogger(&FileLoggerWriter{ChanBufferLength:1024*10,FileName:"/pdata/ymlog/50m_%Y%M%D%H.log",RotateDuration:time.Minute*1,MaxSize:50,//megabytesMaxBackup:50,WriteFileBuffer:50,})for{common.InfoString()}}`)
	}
}

func Test100M(t *testing.T) {
	common := NewLogger(&FileLoggerWriter{
		ChanBufferLength: 1024 * 10,
		FileName:         "/pdata/log/test/100m_%Y%M%D%H_%m%s.log",
		MaxSizeByteSize:  1024 * 1024 * 100,
		WriteFileBuffer:  50,
	})

	for {
		common.InfoString(`packageymlogimport("testing""time")funcTestNewLogger1(t*testing.T){common:=NewLogger(&FileLoggerWriter{FileName:"/pdata/3m.log",RotateDuration:time.Minute*1,MaxSize:3,//megabytesMaxBackup:3,})for{common.InfoString("")}}funcTest100M(t*testing.T){common:=NewLogger(&FileLoggerWriter{ChanBufferLength:1024*10,FileName:"/pdata/ymlog/100m_%Y%M%D%H_%m%s.log",RotateDuration:time.Minute*1,MaxSize:512,//megabytesMaxBackup:50,WriteFileBuffer:50,})for{common.InfoString()}}funcTest50M(t*testing.T){common:=NewLogger(&FileLoggerWriter{ChanBufferLength:1024*10,FileName:"/pdata/ymlog/50m_%Y%M%D%H.log",RotateDuration:time.Minute*1,MaxSize:50,//megabytesMaxBackup:50,WriteFileBuffer:50,})for{common.InfoString()}}`)
	}
}

func Test50M(t *testing.T) {
	common := NewLogger(&FileLoggerWriter{
		ChanBufferLength: 1024 * 10,
		FileName:         "/pdata/log/test/50m_%Y%M%D%H.log",
		MaxSizeByteSize:  1024 * 1024 * 50, // 50 megabytes
		WriteFileBuffer:  50,
	})

	for {
		common.InfoString(`packageymlogimport("testing""time")funcTestNewLogger1(t*testing.T){common:=NewLogger(&FileLoggerWriter{FileName:"/pdata/3m.log",RotateDuration:time.Minute*1,MaxSize:3,//megabytesMaxBackup:3,})for{common.InfoString("")}}funcTest100M(t*testing.T){common:=NewLogger(&FileLoggerWriter{ChanBufferLength:1024*10,FileName:"/pdata/ymlog/100m_%Y%M%D%H_%m%s.log",RotateDuration:time.Minute*1,MaxSize:512,//megabytesMaxBackup:50,WriteFileBuffer:50,})for{common.InfoString()}}funcTest50M(t*testing.T){common:=NewLogger(&FileLoggerWriter{ChanBufferLength:1024*10,FileName:"/pdata/ymlog/50m_%Y%M%D%H.log",RotateDuration:time.Minute*1,MaxSize:50,//megabytesMaxBackup:50,WriteFileBuffer:50,})for{common.InfoString()}}`)
	}
}
