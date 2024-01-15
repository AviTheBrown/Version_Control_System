package datatypes

import "os"

type SVCS map[string]string
type FileInfo struct {
	File     *os.File
	FileName string
}
type User struct {
	UserName string
	FileInfo
}
