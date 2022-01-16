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

// filters files on tfvar files
func filterTfVarsFiles(path string, files []fs.FileInfo) []string {
	fileNames := []string{}

	for _, file := range files {
		if !file.IsDir() && varFileExtensionAllowed(file) {
			fileNames = append(fileNames, fmt.Sprintf("%s/%s", path, file.Name()))
		}
	}

	return fileNames
}

// lists all files in a directory
func listFilesInDir(path string) ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}

	return files, nil
}

// return a list of tfvar files
func getTfVarFilesPaths(path string) ([]string, error) {
	files, err := listFilesInDir(path)

	if err != nil {
		return nil, err
	}

	return filterTfVarsFiles(path, files), nil
}

// checks whether if file exist
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
