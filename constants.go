package qc

const (
	// Pipe |
	Pipe = "|"
	// Tilde ~
	Tilde = "~"
	// Quote "
	Quote = "\""
	// Comma ,
	Comma = ","
	//Carat ^
	Carat = "^"
	// Ear is '\xfe'
	Ear = "\xfe"
	// Nose is Paragraph
	Nose = "\x14"
	// Bom UTF-16
	Bom = "\xef\xbb\xbf"
	// Tab \t
	Tab = "\t"
	// Empty ""
	Empty = ""
)

// Concordance delimiters
var Concordance = Delimiters{Comma: Nose, Quote: Ear}

// Piped is Pipe,Tilde
var Piped = Delimiters{Comma: Pipe, Quote: Tilde}

// PipeCarat is Pipe,Carat
var PipeCarat = Delimiters{Comma: Pipe, Quote: Carat}

// CSV delimiters
var CSV = Delimiters{Comma: Comma, Quote: Quote}

// Tabbed use tab as delimiter with no separator
var Tabbed = Delimiters{Comma: Tab, Quote: Empty}

//var Concord     =   Format(Nose, LeftEar)
