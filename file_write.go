package ymlog

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const DEFAULT_LOG_BUFFER_LENGTH = 2048
const DEFAULT_Rotate_Duration = time.Hour * 24

var line = []byte("\n")

type FileLoggerWriter struct {
	msgChan        chan []byte
	maxsizeCurSize int64
	maxRotateAge   time.Time

	file *os.File
	mu   sync.Mutex

	FileName     string
	realFileName string

	RotateDuration  time.Duration
	MaxSizeByteSize int64 //byte

	//buffer
	ChanBufferLength int
	WriteFileBuffer  int
}

func (w *FileLoggerWriter) start() {
	if w == nil {
		panic("log.FileLoggerWriter error")
	}

	if w.RotateDuration == 0 {
		w.RotateDuration = DEFAULT_Rotate_Duration
	}

	if w.RotateDuration < time.Second {
		//fmt.Println("RotateDuration is less one Second")
		panic(errors.New("RotateDuration is less one Second"))
	}

	w.maxRotateAge = time.Now().Add(w.RotateDuration)

	//ChanBufferLength
	if w.msgChan == nil {
		if w.ChanBufferLength > 0 {
			w.msgChan = make(chan []byte, w.ChanBufferLength)
		} else {
			w.msgChan = make(chan []byte, DEFAULT_LOG_BUFFER_LENGTH)
		}
	}

	//
	if w.WriteFileBuffer == 0 {
		w.WriteFileBuffer = 1
	}

	w.fileRotate(true)

	// check 100 millisecond
	go func() {
		for range time.Tick(time.Microsecond * 100) {
			if err := w.fileRotate(false); err != nil {
				//fmt.Println(fmt.Printf("FileLogWriter(%s): %s\n", w.realFileName, err.Error()))
				fmt.Fprintf(os.Stderr, "FileLogWriter(%q): %s\n", w.FileName, err)
				return
			}
		}
	}()

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

			//var items [][]byte
			msgBody = append(msgBody, <-w.msgChan...)
			msgBody = append(msgBody, line...)
			//items = append(items, <-w.msgChan)
			//Batch to obtain
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
			w.maxsizeCurSize += int64(n)

			// check fileRotate
			if err := w.fileRotate(false); err != nil {
				fmt.Fprintf(os.Stderr, "FileLogWriter(%q): %s\n", w.FileName, err)
				return
			}
		}
	}()
}

// This is the FileLogWriter's output method
func (w *FileLoggerWriter) writeLog(msg []byte) {
	w.msgChan <- msg
}

func (w *FileLoggerWriter) checkRotate(now time.Time) bool {
	//check size
	if w.MaxSizeByteSize > 0 && w.maxsizeCurSize >= w.MaxSizeByteSize {
		return true
	}

	if w.maxRotateAge.Before(now) {
		return true
	}

	return false
}

func (w *FileLoggerWriter) close() {
	close(w.msgChan)
	w.file.Sync()

}

/**
Real file cutting, cutting method lock, and do secondary authentication
*/
func (w *FileLoggerWriter) fileRotate(init bool) error {
	//check un
	w.mu.Lock()
	defer w.mu.Unlock()

	now := time.Now()
	if init {
		//fmt.Println(fmt.Sprintf("%s RotateLog>name>%s,Hour>%d", now.Format(time.RFC3339), w.realFileName, now.Hour()))
		w.realFileName = getActualPathReplacePattern(w.FileName)
		// Open the log file
		checkDir(w.realFileName)
		fileInfo, err := os.Lstat(w.realFileName)
		if err == nil {
			w.maxsizeCurSize = fileInfo.Size()
			w.maxRotateAge = now.Add(w.RotateDuration)
		}

		fd, err := os.OpenFile(w.realFileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
		if err == nil {
			w.file = fd
		} else {
			//fmt.Println(err)
			return fmt.Errorf("Rotate: %s\n", err)
		}
	}

	if w.checkRotate(now) {

		//fmt.Println(fmt.Sprintf("%s RotateLog>name>%s,Hour>%d",
		//	now.Format(time.RFC3339), w.realFileName, now.Hour()))

		if w.maxsizeCurSize > 0 {
			w.file.Sync()
			w.file.Close()
		}

		w.realFileName = getActualPathReplacePattern(w.FileName)
		// Open the log file
		checkDir(w.realFileName)

		_, err := os.Lstat(w.realFileName)
		if err == nil {
			num := 1
			for ; ; num++ {
				fname := w.realFileName + fmt.Sprintf(".%d", num)
				_, err := os.Lstat(fname)
				if err != nil {
					err = os.Rename(w.realFileName, fname)
					break
				}
			}
		}

		fd, err := os.OpenFile(w.realFileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
		if err != nil {
			//fmt.Println(err)
			return err
		}
		w.file = fd

		w.maxRotateAge = now.Add(w.RotateDuration)
		w.maxsizeCurSize = 0
	} else {
		if w.maxsizeCurSize > 0 {
			w.file.Sync()
		}
	}
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
	s := fmt.Sprintf("%02d", now.Second())

	pattern = strings.Replace(pattern, "%Y", Y, -1)
	pattern = strings.Replace(pattern, "%M", M, -1)
	pattern = strings.Replace(pattern, "%D", D, -1)
	pattern = strings.Replace(pattern, "%H", H, -1)
	pattern = strings.Replace(pattern, "%m", m, -1)
	pattern = strings.Replace(pattern, "%s", s, -1)
	return pattern
}
