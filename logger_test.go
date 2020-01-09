package ymlog

import "testing"

func TestNewLogger1(t *testing.T) {

	common := NewLogger(&FileLoggerWriter{
		FileName:    "/pdata/3m.log",
		RotateDaily: true,
		RotateSize:  true,
		MaxSize:     3, // megabytes
		MaxBackup:   3,
	})

	for {
		common.InfoString("common := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWriter")
	}
}

func Test100M(t *testing.T) {
	common := NewLogger(&FileLoggerWriter{
		ChanBufferLength: 1024 * 10,
		FileName:         "/pdata/ymlog/100m_%Y%M%D%H_%m%s.log",
		RotateDaily:      true,
		RotateSize:       true,
		MaxSize:          512, // megabytes
		MaxBackup:        50,
		WriteFileBuffer:  50,
	})

	for {
		common.InfoString(`common := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWriter`)
	}
}

func Test50M(t *testing.T) {
	common := NewLogger(&FileLoggerWriter{
		ChanBufferLength: 1024 * 10,
		FileName:         "/pdata/ymlog/50m_%Y%M%D%H.log",
		RotateDaily:      true,
		RotateSize:       true,
		MaxSize:          50, // megabytes
		MaxBackup:        50,
		WriteFileBuffer:  50,
	})

	for {
		common.InfoString(`common := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWritercommon := NewLogger(&FileLoggerWriter`)
	}
}
