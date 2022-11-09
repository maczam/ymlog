package benchmarks

import (
	"github.com/maczam/ymlog"
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
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

func newZerologLogger() zerolog.LevelWriter {
	timeFormat := "2006-01-02 15:04:05"
	zerolog.TimeFieldFormat = timeFormat
	fileName := "/pdata/log/test/zerolog.log"
	logFile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: timeFormat}
	multi := zerolog.MultiLevelWriter(consoleWriter, logFile)
	return multi
}
