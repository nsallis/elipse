package workers

import (
	// "encoding/json"
	"fmt"
	"os"
)

type Document struct {
	Value             []byte
	Errors            []error
	Source            string      // url or file path
	SourceType        string      // TODO enumerate to: disk, sftp, generated, mixed...
	SourcePermissions os.FileMode // original permissions on file
	TotalFragments    int         // Total number of fragments (splits) to the original file
	FragmentNumber    int         // which fragment (line/entity) this document represents
}

func (doc Document) ToString() (string, error) {
	docJSON := fmt.Sprintf("{Source: %s, SourceType: %s, Errors: %+v, Value: '%v', SourcePermissions: %v}", doc.Source, doc.SourceType, doc.Errors, string(doc.Value), doc.SourcePermissions.String())
	return docJSON, nil
}
