# qc
Read and write delimited text files.

## Synopsis

This library can be used to read and write delimited text files that have headers.

- Rows in file should contain cells that correspond with cells in header
- Delimiter and Separator are determined based on Header row
- Rows are read into map[string]string

## Usage

main/qc.go provides an executable that exposes three functions.

List columns included in header

`qc -h sample.csv`

Subset text file with selected columns

`qc sample.csv name city state zipcode`

Redelimit text with selected delimiter and separator.

`qc -f piped sample.csv > sample.piped`



## Motivation

I needed a simple CSV library that could work with my particular needs:
- supports various delimiter, separator values
- handles large files
- relaxed encoding (where sheets could contain values from multiple encodings)

## License

MIT license.
