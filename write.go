package qc

import (
	"fmt"
	"io"
	"os"
)

func contains(ss []string, t string) bool {
	for _, s := range ss {
		if s == t {
			return true
		}
	}
	return false
}

// QCWriter manages the writing of a delimited file
type QCWriter struct {
	Out    io.Writer
	Del    Delimiters
	Errors int
}

// Line creates string from []string using delimiters set in QCWriter
func (w QCWriter) Line(row []string) string {
	return fmt.Sprintf("%s\r\n", w.Del.Join(row))
}

// WriteRow writes []string as string to Out of QCWriter
func (w QCWriter) WriteRow(row []string) {
	line := w.Line(row)
	_, err := io.WriteString(w.Out, line)
	if err != nil {
		w.Errors += 1
	}
}

// Select subsets header from desired colums
func Select(hdr []string, desired []string) []string {
	selected := make([]string, len(desired))

	for i, col := range desired {
		if !contains(hdr, col) {
			println(col, ": Not found")
		}

		selected[i] = col
	}

	return selected
}

// SubsetRow builds list of values for selected keys
func SubsetRow(selected []string, kv map[string]string) []string {
	row := make([]string, len(selected))
	for i, c := range selected {
		if v, ok := kv[c]; ok {
			row[i] = v
		} else {
			row[i] = ""
		}
	}
	return row
}

// WriteSelected writes the selected columns found in provided dat to stdout
func WriteSelected(dat string, format Delimiters, columns ...string) {

	hdr, lines := Lines(dat)
	selection := Select(hdr, columns)
	writer := QCWriter{Out: os.Stdout, Del: format}

	if len(selection) > 0 {

		if len(selection) > 1 {
			writer.WriteRow(selection)
		}

		for line := range lines {
			row := SubsetRow(selection, line)
			writer.WriteRow(row)
		}
	}
}

// Reformat will echo dat with selected delimiters to stdout
func Reformat(dat string, format Delimiters) {

	hdr, lines := Lines(dat)
	sz := len(hdr)

	writer := QCWriter{Out: os.Stdout, Del: format}
	writer.WriteRow(hdr)

	for line := range lines {
		row := make([]string, sz)

		for i, h := range hdr {
			if col, ok := line[h]; ok {
				row[i] = col
			} else {
				row[i] = ""
			}
		}
		writer.WriteRow(row)
	}
}
