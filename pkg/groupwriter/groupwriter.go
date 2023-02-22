package groupwriter

import (
	"fmt"
	"io"
	"strings"
)

type Writer struct {
	writer io.Writer

	spaceSize   int
	indentLevel int
}

var (
	space = " "
)

func NewWriter(output io.Writer, spaceSize int) *Writer {
	return &Writer{
		writer:    output,
		spaceSize: spaceSize,
	}
}

func (w *Writer) Group(groupName string) {
	fmt.Fprintf(w.writer, "%s%s\n", strings.Repeat(space, w.indentLevel*w.spaceSize), groupName)
	w.indentLevel++
}

func (w *Writer) Println(item string) {
	fmt.Fprintf(w.writer, "%s%s\n", strings.Repeat(space, w.indentLevel*w.spaceSize), item)
}

func (w *Writer) Endgroup() {
	w.indentLevel--
}

func (w *Writer) Flush() {
	if f, ok := w.writer.(interface{ Flush() }); ok {
		f.Flush()
	}
}
