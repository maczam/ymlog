package ymlog

import (
	"io"
)

//start()
//writeLog([]byte)
//close()

type OutLoggerWriter struct {
	Out io.Writer
}

func (w *OutLoggerWriter) start() {

}

func (w *OutLoggerWriter) writeLog(msg []byte) {
	w.Out.Write(msg)
	w.Out.Write([]byte("\n"))
}

func (w *OutLoggerWriter) close() {
}
