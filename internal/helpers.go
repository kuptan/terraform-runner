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

func varFileExtensionAllowed(file string) bool {
	ext := filepath.Ext(file)

	return ext == ".tfvars" || ext == ".tf" || ext == ".json"
}

// filters files on tfvar files
func filterTfVarsFiles(files []string) []string {
	fileNames := []string{}

	for _, file := range files {

		if varFileExtensionAllowed(file) {
			fileNames = append(fileNames, file)
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

// lists all files in a directory and its sub directories
func listFilesInDirNested(basePath string) ([]string, error) {
	files := []string{}
	all, err := listFilesInDir(basePath)

	if err != nil {
		return nil, err
	}

	for _, f := range all {
		namePath := fmt.Sprintf("%s/%s", basePath, f.Name())

		if f.IsDir() {
			listed, err := listFilesInDirNested(namePath)

			if err != nil {
				return nil, err
			}

			files = append(files, listed...)
		} else {
			files = append(files, fmt.Sprintf("%s/%s", basePath, f.Name()))
		}
	}

	return files, nil
}

// return a list of tfvar files
func getTfVarFilesPaths(path string) ([]string, error) {
	files, err := listFilesInDirNested(path)

	if err != nil {
		return nil, err
	}

	return filterTfVarsFiles(files), nil
}

// checks whether if file exist
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
