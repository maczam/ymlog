package ymlog

import "fmt"

//start()
//writeLog([]byte)
//close()

type ConsoleLoggerWriter struct {
}

func (w *ConsoleLoggerWriter) start() {

}

func (w *ConsoleLoggerWriter) writeLog(msg []byte) {
	fmt.Println(*bytesToString(msg))
}

func (w *ConsoleLoggerWriter) close() {

}
