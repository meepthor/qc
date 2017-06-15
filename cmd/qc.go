package main

import (
	"flag"
	"fmt"
	"github.com/meepthor/qc"
)

func main() {

	var showHeader = flag.Bool("h", false, "List Header")
	var format = flag.String("f", "pipe", "concord csv hat pipe tab")
	
	flag.Parse()

	delimiters := qc.NamedDelimiters(*format)

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
