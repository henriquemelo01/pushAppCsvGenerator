package model

import "bytes"

type ReportCsvFile struct {
	FileName   string
	FileBuffer *bytes.Buffer
}
