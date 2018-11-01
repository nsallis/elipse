package workers

import (
	// "encoding/json"
	"fmt"
)

type Document struct {
	Value      []byte
	Errors     []error
	Source     string // url or file path
	SourceType string // TODO enumerate to: disk, sftp, generated, mixed...
}

func (doc Document) ToString() (string, error) {
	docJson := fmt.Sprintf("{Source: %s, SourceType: %s, Errors: %+v, Value: '%v'}", doc.Source, doc.SourceType, doc.Errors, string(doc.Value))
	return docJson, nil
}
