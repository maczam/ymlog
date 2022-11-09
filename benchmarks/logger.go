package benchmarks

import (
	"github.com/maczam/ymlog"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func newZapLogger() *zap.Logger {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeDuration = zapcore.NanosDurationEncoder
	ec.EncodeTime = zapcore.EpochNanosTimeEncoder
	enc := zapcore.NewJSONEncoder(ec)
	return zap.New(zapcore.NewCore(
		enc,
		getLogWriter(),
		zap.InfoLevel,
	))
}

func newYmLog1024Logger() *ymlog.Logger {
	return ymlog.NewLogger(&ymlog.FileLoggerWriter{
		FileName:         "/pdata/log/test/ymlog1024_%Y%M%D-%H%m.log",
		MaxSizeByteSize:  1024 * 1024 * 100,
		RotateDuration:   time.Hour,
		ChanBufferLength: 1024,
		WriteFileBuffer:  1024,
	},
	)
}

func newYmLog2048Logger() *ymlog.Logger {
	return ymlog.NewLogger(&ymlog.FileLoggerWriter{
		FileName:         "/pdata/log/test/ymlog2048_%Y%M%D-%H%m.log",
		MaxSizeByteSize:  1024 * 1024 * 100,
		RotateDuration:   time.Hour,
		ChanBufferLength: 2048,
		WriteFileBuffer:  2048,
	},
	)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "/pdata/log/test/zap_%Y%M%D-%H%m.log",
		MaxSize:    100,   // M
		MaxBackups: 1,     // 最大备份数量
		MaxAge:     30,    // 最大备份天数
		Compress:   false, // 是否压缩`
	}
	return zapcore.AddSync(lumberJackLogger)
}
