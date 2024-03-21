package filesync

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func copyFile(src, dst string) {
	srcInfo, err := os.Stat(src)
	if err != nil {
		fmt.Printf("there was problem finiding file: %s", srcInfo)
		return
	}
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Printf("There was a proble opening: %v", srcFile)
		return
	}

	_, err = os.Stat(dst)
	if err != nil {
		fmt.Printf("there was a problem finding file: %v", dst)
		return
	}

	dstFile, err := os.Create(dst)
	dstFile.Close()
	if err != nil {
		fmt.Printf("there was a problem with creating or finding the file %v: ", dst)
		return
	}

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		fmt.Printf("there was problem copying %v to %v ", srcFile, dstFile)
		return
	}
	fmt.Printf("File %v Copied to %s\n", src, dst)
}

func SyncFiles(originalDir, copiedDir string, filename string) bool {
	var hasChanged bool
	// ./ <fileName>
	originalFilePath := filepath.Join(originalDir, filename)
	// vcs/commits/<fileName>
	copiedFIlepath := filepath.Join(copiedDir, filename)
	for {
		origFilePathModTime := getModTime(originalFilePath)
		copyFilePathModTime := getModTime(copiedFIlepath)

		if origFilePathModTime.After(copyFilePathModTime) {
			copyFile(originalFilePath, copiedFIlepath)
			hasChanged = true
			break
		}
		fmt.Println("the has been no change to the file.")

		time.Sleep(time.Second * 1)
	}
	return hasChanged
}

func getModTime(filePath string) time.Time {

	fileData, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("there was a proble retrieving")
	}

	return fileData.ModTime()
}
