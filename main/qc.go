package main

import (
	"flag"
	"fmt"
	"meepthor/qc"
)

func main() {

	var showHeader = flag.Bool("h", false, "List Header")
	var format = flag.String("f", "pipe", "concord csv hat pipe tab")
	var delimiters qc.Delimiters

	flag.Parse()

	switch *format {
	case "concord":
		delimiters = qc.Concordance
	case "pipe":
		delimiters = qc.Piped
	case "hat":
		delimiters = qc.PipeCarat
	case "tab":
		delimiters = qc.Tabbed
	case "csv":
		delimiters = qc.CSV
	default:
		delimiters = qc.Piped
	}

	if *showHeader && flag.NArg() > 0 {
		dat := flag.Args()[0]
		hdr, _ := qc.Lines(dat)
		for _, h := range hdr {
			fmt.Println(h)
		}
	} else if flag.NArg() >= 2 {
		dat := flag.Args()[0]
		cols := flag.Args()[1:]
		qc.WriteSelected(dat, delimiters, cols...)
	} else if flag.NArg() >= 1 {
		dat := flag.Args()[0]
		qc.Reformat(dat, delimiters)
	} else {
		flag.Usage()
	}

}
