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
