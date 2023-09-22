package sku

type ErrorKind string

const (
	parse      ErrorKind = "parseError"
	fileFormat ErrorKind = "fileFormatError"
	csvHead    ErrorKind = "csvHeadError"
	csvRow     ErrorKind = "csvRowError"
	unknown    ErrorKind = "unknownError"
)

type ParseError struct {
	kind    ErrorKind
	message string
}

func newParseError(kind ErrorKind, message string) ParseError {
	return ParseError{kind, message}
}

func (p ParseError) Kind() ErrorKind {
	return p.kind
}

func (p ParseError) Error() string {
	return p.message
}
