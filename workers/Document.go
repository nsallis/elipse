package workers


type Document struct {
  Value []byte
  Errors []error
  Source string
  SourceType string // TODO enumerate to: disk, sftp, generated, mixed...
}
