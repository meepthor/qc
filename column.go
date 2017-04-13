package qc

import (
	"bytes"
	"fmt"
	"strings"
)

// Delimiters defines delimiter and separator
type Delimiters struct {
	// Comma is delimiter
	Comma string
	// Quote is separator
	Quote string
}

func (d Delimiters) cstream(s string) <-chan string {
	cs := make(chan string)
	go func() {
		for strings.Contains(s, d.Comma) {
			found := strings.Index(s, d.Comma)
			cs <- s[:found]
			s = s[found+1:]
		}
		cs <- s // whether or not len(s) > 0
		close(cs)
	}()
	return cs
}

func (d Delimiters) qstream(s string) <-chan string {

	qs := make(chan string)

	var buf = make([]string, 0)
	var yield = func() { qs <- strings.Join(buf, d.Comma); buf = nil }
	var quoted = false

	go func() {
		for c := range d.cstream(s) {
			buf = append(buf, c)
			if strings.HasSuffix(c, d.Quote) {
				if quoted || (strings.HasPrefix(c, d.Quote) && len(c) > 1) {
					yield()
					quoted = false
				} else {
					quoted = true
				}
			} else if !quoted {
				if strings.HasPrefix(c, d.Quote) {
					quoted = true
				} else {
					yield()
				}
			}
		}
		if len(buf) > 0 {
			yield()
		}
		close(qs)
	}()
	return qs
}

// Trim removes separator from column.
func (d Delimiters) Trim(col string) string {
	if len(col) > 1 {
		if strings.HasPrefix(col, d.Quote) {
			if strings.HasSuffix(col, d.Quote) {
				return col[1 : len(col)-1]
			}
		}
	}
	return col

}

// Split returns []string for provided line of text.
func (d Delimiters) Split(line string) []string {
	if d.Quote == "" {
		return strings.Split(line, d.Comma)
	}

	buffer := make([]string, 0)
	for q := range d.qstream(line) {
		buffer = append(buffer, d.Trim(q)) //undouble(trim(q)))
	}

	return buffer

}

// GetDelimiters assigns Quote and Comma based on string s.
func GetDelimiters(s string) Delimiters {

	qcq := func(c, q string) string { return fmt.Sprintf("%s%s%s", q, c, q) }

	found := func(c, q string) bool {
		return strings.Contains(s, qcq(c, q))
	}

	var q, c string

	if found(Nose, Ear) {
		c = Nose
		q = Ear
	} else if found(Pipe, Tilde) {
		c = Pipe
		q = Tilde
	} else if found(Pipe, Carat) {
		c = Pipe
		q = Carat
	} else if found(Comma, Quote) {
		c = Comma
		q = Quote
	} else if found(Tab, "") {
		c = "\t"
		q = ""
	} else {

		// Probably no delimiter
		first := ""
		if len(s) > 0 {
			first = string(s[0])
		}
		c = Pipe
		q = first
	}

	return Delimiters{Comma: c, Quote: q}
}

func (d Delimiters) simpleJoin(row []string) string {
	qcq := fmt.Sprintf("%s%s%s", d.Quote, d.Comma, d.Quote)
	return fmt.Sprintf("%s%s%s", d.Quote, strings.Join(row, qcq), d.Quote)
}

func (d Delimiters) quoteCol(col string) string {
	if strings.Contains(col, d.Comma) || strings.Contains(col, d.Quote) {
		col = fmt.Sprintf("%s%s%s", d.Quote, col, d.Quote)
	}
	return col
}

func (d Delimiters) joinCSV(row []string) string {

	bs := bytes.NewBufferString("")
	i := 0
	max := len(row) - 1

	for ; i < max; i++ {
		bs.WriteString(d.quoteCol(row[i]))
		bs.WriteString(d.Comma)
	}

	bs.WriteString(d.quoteCol(row[i]))
	return bs.String()
}

// Join joins row with quote and comma defined in d
func (d Delimiters) Join(row []string) string {
	if len(row) == 1 {
		return row[0]
	} else if d.Comma == Comma {
		return d.joinCSV(row)
	}
	return d.simpleJoin(row)
}