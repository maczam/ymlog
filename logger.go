package ymlog

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

const default_log_buffer_length int = 2048

var (
	//logBufferLength = 2048
	line = []byte("\n")
)

// the associated LogWriter.
type Logger struct {
	LogWriter *FileLogWriter
}

func (log *Logger) start() {
	w := log.LogWriter
	if w == nil {
		panic("log.LogWriter error")
	}
	if w.MaxSize > 0 {
		w.maxSize_bytesize = w.MaxSize * 1024 * 1024
	}

	//ChanBufferLength
	if w.msgChan == nil {
		if w.ChanBufferLength > 0 {
			w.msgChan = make(chan []byte, w.ChanBufferLength)
		} else {
			w.msgChan = make(chan []byte, default_log_buffer_length)
		}
	}

	//
	if w.WriteFileBuffer == 0 {
		w.WriteFileBuffer = 1
	}

	now := time.Now()
	w.fileRotate(true, &now)

	go func() {
		defer func() {
			if w.file != nil {
				w.file.Close()
			}
		}()

		var msgBody []byte
		//https://blog.drkaka.com/batch-get-from-golangs-buffered-channel-9638573f0c6e 批量保存日志
		for {
			msgBody = msgBody[:0]
			// 判断是否按照天来滚动
			now := time.Now()
			if w.checkRotate(&now) {
				fmt.Println(fmt.Sprintf("%s RotateLog>name>%s,Hour>%d,hourly_opendate>%d,Day>%d,daily_opendate>%d", now.Format(time.RFC3339), w.FileName, now.Hour(), w.hourly_opendate, now.Day(), w.daily_opendate))
				if err := w.fileRotate(false, &now); err != nil {
					fmt.Fprintf(os.Stderr, "FileLogWriter(%q): %s\n", w.FileName, err)
					return
				}
			}

			//var items [][]byte
			msgBody = append(msgBody, <-w.msgChan...)
			msgBody = append(msgBody, line...)
			//items = append(items, <-w.msgChan)
			//批量获取
		Remaining:
			for i := 1; i < w.WriteFileBuffer; i++ {
				select {
				case item := <-w.msgChan:
					//items = append(items, item)
					msgBody = append(msgBody, item...)
					msgBody = append(msgBody, line...)
				default:
					break Remaining
				}
			}
			//msgBody := join(items, line)
			n, err := w.file.Write(msgBody)
			if err != nil {
				fmt.Fprintf(os.Stderr, "FileLogWriter(%q): %s\n", w.FileName, err)
				return
			}
			w.maxsize_cursize += n
		}
	}()
}

/**
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
	close(log.LogWriter.msgChan)
	log.LogWriter.Close()
}

func (log *Logger) Info(arg0 *string) {
	if arg0 != nil {
		bytes := stringToBytes(arg0)
		log.LogWriter.LogWrite(bytes)
	}
}

func (log *Logger) InfoString(arg0 string) {
	if arg0 != "" {
		//bytes := stringToBytes(&arg0)
		log.LogWriter.LogWrite([]byte(arg0))
	}
}

func (log *Logger) InfoBytes(arg0 []byte) {
	if arg0 != nil {
		log.LogWriter.LogWrite(arg0)
	}
}

type FileLogWriter struct {
	msgChan          chan []byte
	maxsize_cursize  int
	maxSize_bytesize int
	daily_opendate   int
	hourly_opendate  int
	file             *os.File

	FileName string

	// Rotate
	RotateSize   bool
	DailyRotate  bool
	HourlyRotate bool

	// Keep old logfiles (.001, .002, etc)
	MaxBackup int
	MaxSize   int //单位为M

	//buffer
	ChanBufferLength int
	WriteFileBuffer  int
}

// This is the FileLogWriter's output method
func (w *FileLogWriter) LogWrite(msg []byte) {
	w.msgChan <- msg
}

func (w *FileLogWriter) checkRotate(now *time.Time) bool {
	if (w.RotateSize && w.maxSize_bytesize > 0 && w.maxsize_cursize >= w.maxSize_bytesize) ||
		(w.HourlyRotate && now.Hour() != w.hourly_opendate) ||
		(w.DailyRotate && now.Day() != w.daily_opendate) {
		return true
	} else {
		return false
	}
}

func (w *FileLogWriter) Close() {
	w.file.Sync()
}

/**
  真正文件切割，切割的时候方法加锁，并且做二次认证
*/
func (w *FileLogWriter) fileRotate(init bool, now *time.Time) error {
	//拿到做
	if !w.checkRotate(now) {
		fileName := getActualPathReplacePattern(w.FileName)
		checkDir(fileName)
		fd, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
		if err != nil {
			return err
		}
		w.maxsize_cursize = 0
		w.file = fd
		w.daily_opendate = now.Day()
		w.hourly_opendate = now.Hour()
		return nil
	}

	if w.file != nil {
		w.file.Sync()
		w.file.Close()
	}

	fileName := getActualPathReplacePattern(w.FileName)
	fmt.Println(fmt.Sprintf("%s fileRotate>name>%s,Hour>%d,hourly_opendate>%d,Day>%d,daily_opendate>%d,fileName>%s", now.Format(time.RFC3339), w.FileName, now.Hour(), w.hourly_opendate, now.Day(), w.daily_opendate, fileName))

	_, err := os.Lstat(fileName)
	if err == nil {
		if !init {
			num := 1
			fname := ""
			num = w.MaxBackup - 1
			for ; num >= 1; num-- {
				fname = fileName + fmt.Sprintf(".%d", num)
				nfname := fileName + fmt.Sprintf(".%d", num+1)
				_, err = os.Lstat(fname)
				if err == nil {
					os.Rename(fname, nfname)
				}
			}
			err = os.Rename(fileName, fname)
			if err != nil {
				return fmt.Errorf("Rotate: %s\n", err)
			}
		}
	}

	// Open the log file
	checkDir(fileName)
	fd, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	w.file = fd

	w.daily_opendate = now.Day()
	w.hourly_opendate = now.Hour()
	w.maxsize_cursize = 0
	return nil
}

func checkDir(fileName string) {
	dir := filepath.Dir(fileName)
	os.MkdirAll(dir, os.ModePerm)
}

func getActualPathReplacePattern(pattern string) string {
	now := time.Now()
	Y := fmt.Sprintf("%d", now.Year())
	M := fmt.Sprintf("%02d", now.Month())
	D := fmt.Sprintf("%02d", now.Day())
	H := fmt.Sprintf("%02d", now.Hour())
	m := fmt.Sprintf("%02d", now.Minute())

	pattern = strings.Replace(pattern, "%Y", Y, -1)
	pattern = strings.Replace(pattern, "%M", M, -1)
	pattern = strings.Replace(pattern, "%D", D, -1)
	pattern = strings.Replace(pattern, "%H", H, -1)
	pattern = strings.Replace(pattern, "%m", m, -1)
	return pattern
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

/**
  new Logger
*/
func NewFileLogger(logWriter *FileLogWriter) *Logger {
	logger := &Logger{
		LogWriter: logWriter,
	}
	logger.start()
	return logger
}
