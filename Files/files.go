package files

import (
	datatypes "Version_Control_System/DataTypes"
	"os"
)

func CreateFile(filePath string) (*os.File, error) {
	createdFile := new(datatypes.FileInfo)
	var err error
	createdFile.File, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return createdFile.File, nil
}
