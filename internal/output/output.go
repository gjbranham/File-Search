package output

import (
	"fmt"
	"sort"
	"time"

	"golang.org/x/exp/slices"

	"github.com/gjbranham/Text-Finder/internal/concurrency"
)

func PrintResults(start time.Time, searchTerms []string, matchInfo *concurrency.MatchInfo) {
	// copy matchInfo then sort it for nice looking output
	matchInfoCopy := copyMatchInfo(matchInfo)
	sortMatchInfoByKeyThenFile(matchInfoCopy)

	matchCount := printAndCountMatches(matchInfoCopy, searchTerms)

	Print(fmt.Sprintf("Found %v matches in %v files in %v", matchInfoCopy.Count, matchCount, time.Since(start)))
}

func copyMatchInfo(matchInfo *concurrency.MatchInfo) *concurrency.MatchInfo {
	var matchInfoCopy concurrency.MatchInfo
	matchInfoCopy.Count = matchInfo.Count
	matchInfoCopy.Matches = matchInfo.Matches
	return &matchInfoCopy
}

func sortMatchInfoByKeyThenFile(matchInfo *concurrency.MatchInfo) {
	sort.Slice(matchInfo.Matches, func(i, j int) bool {
		if matchInfo.Matches[i].Key == matchInfo.Matches[j].Key {
			return matchInfo.Matches[i].File < matchInfo.Matches[j].File
		}
		return matchInfo.Matches[i].Key < matchInfo.Matches[j].Key
	})
}

func printAndCountMatches(matchInfo *concurrency.MatchInfo, searchTerms []string) int {
	padding := calcPadding(searchTerms)
	uniqFiles := make([]string, 0)
	customFmt := fmt.Sprintf("%%-%ds: %%s line %%v", padding)
	for _, item := range matchInfo.Matches {
		Print(fmt.Sprintf(customFmt, item.Key, item.File, item.LineNum))
		if !slices.Contains(uniqFiles, item.File) {
			uniqFiles = append(uniqFiles, item.File)
		}
	}
	return len(uniqFiles)
}

func calcPadding(searchTerms []string) int {
	padding := -1
	for _, t := range searchTerms {
		if len(t) > padding {
			padding = len(t)
		}
	}
	return padding
}
