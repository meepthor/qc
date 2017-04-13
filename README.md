# qc
Read and write delimited text files.

## Synopsis

This library can be used to read and write delimited text files that have headers.

- Every row in file should match header
- Delimiter and Separator are assigned based on Header row
- Rows are read into map[string]string

## Code Example

(tbd)


## Motivation

I needed a simple CSV library that could work with my particular needs:
- supports various delimiter, separator values
- handles large files
- relaxed encoding (where sheets could contain values from multiple encodings)

## License

MIT license.
