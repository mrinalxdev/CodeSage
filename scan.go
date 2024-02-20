package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

func getDotFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dotFile := usr.HomeDir + "/.goailocalstats"
	return dotFile
}

func openFile(filePath string) *os.File {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0755)

	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create(filePath)
			if err != nil {
				panic(err)
			}

		} else {
			panic(err)
		}
	}

	return f
}

func parseFileLinestoSlice(filePath string) []string {
	f := openFile(filePath)
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		if err != io.EOF{
			panic(err)
		}
	}


	return lines
}

func sliceContains(slice []string, value string) bool {

	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func joinSlices(new []string, existing []string) []string {
	for _, i := range new {
		if !sliceContains(existing, i) {
			existing = append(existing, i)
		}
	}
	return existing
}

func dumpStringSlicetoFile(repos []string, filePath string){
	content := strings.Join(repos, "\n")
	os.WriteFile(filePath, []byte(content), 0755)
}

func addNewSliceElementsToFile(filePath string, newRepos []string) {
	existingRepos := parseFileLinestoSlice(filePath)
	repos := joinSlices(newRepos, existingRepos)
	dumpStringSlicetoFile(repos, filePath)
}

func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

func scanGitFolders(folders []string, folder string) []string {
	folder = strings.TrimSuffix(folder, "/")

	return folders
}