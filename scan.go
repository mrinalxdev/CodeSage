package main 

import (
	"bufio"
	"fmt"
	"io"
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

