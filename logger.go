package ymlog

import (
	"reflect"
	"unsafe"
)

// the associated LogWriter.
type Logger struct {
	logWriter LoggerWrite
}
type LoggerWrite interface {
	start()
	writeLog([]byte)
	close()
}

func (log *Logger) start() {
	log.logWriter.start()
}

/*
*

	每个后面都要增加
*/
func join(s [][]byte, sep byte) []byte {
	if len(s) == 0 {
		return []byte{}
	}
	if len(s) == 1 {
		b := make([]byte, len(s[0])+1)
		copy(b, s[0])
		b[len(s[0])] = sep
		// Just return a copy.
		return b
	}
	n := len(s)
	for _, v := range s {
		n += len(v)
	}

	b := make([]byte, n)
	bp := 0
	//bp := copy(b, s[0])
	//b[bp] = sep
	for _, v := range s[0:] {
		//bp += copy(b[bp:], sep)
		//bp = bp + 1
		bp += copy(b[bp:], v)
		b[bp] = sep
		bp = bp + 1
	}
	return b
}

func (log *Logger) Close() {
	log.logWriter.close()
}

func (log *Logger) Info(arg0 *string) {
	if arg0 != nil {
		bytes := stringToBytes(arg0)
		log.logWriter.writeLog(bytes)
	}
}

func (log *Logger) InfoString(arg0 string) {
	if arg0 != "" {
		log.logWriter.writeLog([]byte(arg0))
	}
}

func (log *Logger) InfoBytes(arg0 []byte) {
	if arg0 != nil {
		log.logWriter.writeLog(arg0)
	}
}

func bytesToString(b []byte) *string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return (*string)(unsafe.Pointer(&sh))
}

func stringToBytes(s *string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(s))
	bh := reflect.SliceHeader{sh.Data, sh.Len, sh.Len}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

/*
*

	new Logger
*/
func NewLogger(logWriter LoggerWrite) *Logger {
	logger := &Logger{
		logWriter: logWriter,
	}
	logger.start()
	return logger
}
