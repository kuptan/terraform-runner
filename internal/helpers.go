package internal

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

// checks whether a string is part of an array of strings
func arrayContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func varFileExtensionAllowed(file fs.FileInfo) bool {
	ext := filepath.Ext(file.Name())

	return ext == ".tfvars" || ext == ".tf"
}

// lists all files in a directory
func listFilesInDir(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}

	fileNames := []string{}

	for _, file := range files {
		if !file.IsDir() && varFileExtensionAllowed(file) {
			fileNames = append(fileNames, fmt.Sprintf("%s/%s", path, file.Name()))
		}
	}

	return fileNames, nil
}

// checks whether if file exist
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
