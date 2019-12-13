// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudsearchdomain

type ContentType string

// Enum values for ContentType
const (
	ContentTypeApplicationJson ContentType = "application/json"
	ContentTypeApplicationXml  ContentType = "application/xml"
)

func (enum ContentType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ContentType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type QueryParser string

// Enum values for QueryParser
const (
	QueryParserSimple     QueryParser = "simple"
	QueryParserStructured QueryParser = "structured"
	QueryParserLucene     QueryParser = "lucene"
	QueryParserDismax     QueryParser = "dismax"
)

func (enum QueryParser) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum QueryParser) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}