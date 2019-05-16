package ymlog

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//start()
//writeLog([]byte)
//close()
const DEFAULT_LOG_BUFFER_LENGTH = 2048

var line = []byte("\n")

type FileLoggerWriter struct {
	msgChan         chan []byte
	maxsizeCurSize  int
	maxSizeByteSize int
	dailyOpenDate   int
	hourlyOpenDate  int
	file            *os.File

	FileName string

	// Rotate
	RotateSize   bool
	RotateDaily  bool
	RotateHourly bool

	// Keep old logfiles (.001, .002, etc)
	MaxBackup int
	MaxSize   int //单位为M

	//buffer
	ChanBufferLength int
	WriteFileBuffer  int
}

func (w *FileLoggerWriter) start() {
	if w == nil {
		panic("log.FileLoggerWriter error")
	}

	if w.MaxSize > 0 {
		w.maxSizeByteSize = w.MaxSize * 1024 * 1024
	}

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
			// check fileRotate
			now := time.Now()
			if w.checkRotate(&now) {
				fmt.Println(fmt.Sprintf("%s RotateLog>name>%s,Hour>%d,hourly_opendate>%d,Day>%d,daily_opendate>%d",
					now.Format(time.RFC3339), w.FileName, now.Hour(), w.hourlyOpenDate, now.Day(), w.dailyOpenDate))
				if err := w.fileRotate(false, &now); err != nil {
					fmt.Fprintf(os.Stderr, "FileLogWriter(%q): %s\n", w.FileName, err)
					panic(err)
					return
				}
			}

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
			w.maxsizeCurSize += n
		}
	}()
}

// This is the FileLogWriter's output method
func (w *FileLoggerWriter) writeLog(msg []byte) {
	w.msgChan <- msg
}

func (w *FileLoggerWriter) checkRotate(now *time.Time) bool {
	if (w.RotateSize && w.maxSizeByteSize > 0 && w.maxsizeCurSize >= w.maxSizeByteSize) ||
		(w.RotateHourly && now.Hour() != w.hourlyOpenDate) ||
		(w.RotateDaily && now.Day() != w.dailyOpenDate) {
		return true
	} else {
		return false
	}
}

func (w *FileLoggerWriter) close() {
	close(w.msgChan)
	w.file.Sync()

}

/**
Real file cutting, cutting method lock, and do secondary authentication
*/
func (w *FileLoggerWriter) fileRotate(init bool, now *time.Time) error {
	//拿到做
	if !w.checkRotate(now) {
		fileName := getActualPathReplacePattern(w.FileName)
		checkDir(fileName)
		fd, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
		if err != nil {
			return err
		}
		w.maxsizeCurSize = 0
		w.file = fd
		w.dailyOpenDate = now.Day()
		w.hourlyOpenDate = now.Hour()
		return nil
	}

	if w.file != nil {
		w.file.Sync()
		w.file.Close()
	}

	fileName := getActualPathReplacePattern(w.FileName)
	fmt.Println(fmt.Sprintf("%s fileRotate>name>%s,Hour>%d,hourly_opendate>%d,Day>%d,daily_opendate>%d,fileName>%s",
		now.Format(time.RFC3339), w.FileName, now.Hour(), w.hourlyOpenDate, now.Day(), w.dailyOpenDate, fileName))

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

	w.dailyOpenDate = now.Day()
	w.hourlyOpenDate = now.Hour()
	w.maxsizeCurSize = 0
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
