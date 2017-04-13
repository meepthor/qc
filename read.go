package qc

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"
)

// Zip header columns with row columns
func Zip(hs, cs []string) (map[string]string, bool) {

	row := make(map[string]string, len(hs))

	if len(cs) != len(hs) {
		return row, false
	} else {
		for i := 0; i < len(hs); i++ {
			v := cs[i]
			if len(v) > 0 {
				row[hs[i]] = v
			}
		}
	}

	return row, true

}

// NormHeader normalizes header columns names
func NormHeader(cs []string) []string {
	row := make([]string, len(cs))
	for i := 0; i < len(cs); i++ {
		c := strings.ToLower(cs[i])
		c = strings.Replace(c, "_", "", -1)
		c = strings.Replace(c, " ", "", -1)
		row[i] = c
	}
	return row
}

// LineIterator returns a generator for lines of text in a text file.
func LineIterator(fname string) <-chan string {

	ch := make(chan string)
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)

	go func() {

		var err error
		var txt []byte
		var isPrefix bool

		var buf bytes.Buffer

		for {
			txt, isPrefix, err = reader.ReadLine()
			buf.Write(txt)

			if isPrefix {
				continue
			}

			if err == nil {
				ch <- buf.String()
				buf.Reset()
			} else {

				s := buf.String()
				if len(s) > 0 {
					ch <- buf.String()
				}

				close(ch) // Remember to close or the loop never ends!
				break
			}
		}

		defer file.Close()
	}()

	return ch
}

//  Lines returns header and line iterator for provided textfile.
func Lines(fname string) ([]string, <-chan map[string]string) {

	lines := LineIterator(fname)
	first := strings.Replace(<-lines, Bom, "", 1)

	d := GetDelimiters(first)
	hdr := NormHeader(d.Split(first))

	rc := make(chan map[string]string)
	rowCount := 1

	go func() {
		for line := range lines {
			rowCount++
			cols := d.Split(line)
			row, ok := Zip(hdr, cols)
			if ok {
				rc <- row
			} else {
				log.Printf(
					"[ %s ] => Line: %d, Row: %d, Header: %d\n",
					fname, rowCount, len(cols), len(hdr))
			}
		}
		close(rc)
	}()

	return hdr, rc

}
