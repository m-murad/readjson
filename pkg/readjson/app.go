package readjson

import (
	"io"
	"log"
)

type App struct {
	// internal
	processChan chan JSONEntity

	l     *log.Logger
	out   io.Writer
	files string
}

func NewApp(l *log.Logger, files string, out io.Writer) *App {
	return &App{
		processChan: make(chan JSONEntity),
		l:           l,
		out:         out,   // Using this so that the code is testable.
		files:       files, // Maybe this can be changed for better testing, include this as a file(s)
	}
}

func (a *App) Start() {

}

func (a *App) processFiles() {
	// use filepath.Walk to read all files in `a.files`` and send the bytes to the `processChan`

	// dec := json.NewDecoder(bufio.NewReader(a.file))
	// dec.Token()
	// for dec.More() {
	// 		// unmarshal the data and send via channel or print error
	// 		a.processChan <- JSONEntity{}
	// }

	// close(a.processChan)

}

func (a *App) printResults() {
	// read the data from `processChan`. Print it to `a.out`

}
