package application

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/gjbranham/Text-Finder/internal/args"
	"github.com/gjbranham/Text-Finder/internal/concurrency"
	out "github.com/gjbranham/Text-Finder/internal/output"
)

type TextFinder struct {
	Args      *args.Arguments
	MatchInfo *concurrency.MatchInfo
}

func (a *TextFinder) FindFiles(rootPath string) {
	var wg sync.WaitGroup

	files, err := os.ReadDir(rootPath)
	if err != nil {
		out.Print(fmt.Sprintf("error occurred while walking root dir: %v\n", err))
	}

	for _, fo := range files {
		path := path.Join(rootPath, fo.Name())

		if fo.IsDir() && a.Args.RecursiveSearch {
			wg.Add(1)
			go func(path string) {
				defer wg.Done()
				a.FindFiles(path)
			}(path)
		}

		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			a.CheckFileForMatch(path)
		}(path)
	}
	wg.Wait()
}

func (a *TextFinder) CheckFileForMatch(file string) {
	fileObj, err := os.Open(file)
	if err != nil {
		out.Print(fmt.Sprintf("Failed to open file '%v': %v\n", file, err))
		return
	}
	defer fileObj.Close()

	lineNum := 1
	localMatchCnt := 0
	localMatchList := []concurrency.FileInfo{}

	r := bufio.NewScanner(fileObj)
	for r.Scan() {
		line := r.Text()
		for _, key := range a.Args.SearchTerms {
			if strings.Contains(line, "\x00") {
				out.Print(fmt.Sprintf("Ignoring binary file %v", file))
				return
			}
			if strings.Contains(line, key) || (strings.Contains(strings.ToLower(line), strings.ToLower(key)) && a.Args.CaseInsensitive) {
				localMatchCnt++
				localMatchList = append(localMatchList, concurrency.FileInfo{Key: key, File: file, LineNum: lineNum})
			}
		}
		lineNum++
	}
	a.MatchInfo.CounterInc(localMatchCnt)
	a.MatchInfo.AddMatch(localMatchList...)
}
