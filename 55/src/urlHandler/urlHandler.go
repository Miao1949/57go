package urlHandler

import (
	"dao"
	"strings"
	"sort"
)

//---------------------------------------------------------
// TYPES
//---------------------------------------------------------
type StringSlice []string

func (p StringSlice) Len() int {return len(p)}
func (p StringSlice) Less(i, j int) bool {
	if len(p[i]) != len(p[j]) { // First sort by length...
		return len(p[i]) < len(p[j])
	} else { // .. sort strings of equal length alphabetically.
		return p[i] < p[j]
	}
}
func (p StringSlice) Swap(i, j int) {p[i], p[j] = p[j], p[i]}


//---------------------------------------------------------
// FUNCTIONS
//---------------------------------------------------------

func GenerateAndAddSnippet(text string) (url string, err error) {
	urls, err := dao.GetAllUrls()

	if err != nil {
		return
	}

	if len(urls) == 0 {
		url ="a"
	} else {
		SortSliceByLength(urls)
		url = nextIdentifierAfter(urls[len(urls) -1])
	}

	err = dao.AddTextSnippet(url, text)

	return
}

func urlExists(url string) (found bool, err error) {
	return dao.UrlExists(url)
}

func SortSliceByLength(sliceToSort []string) {sort.Sort(StringSlice(sliceToSort))}

func nextIdentifierAfter(s string) (nextIdentifier string) {
	nextIdentifierAsSlice := []rune(s)
	indexOfLastRune := len(nextIdentifierAsSlice) - 1

	if nextIdentifierAsSlice[indexOfLastRune] < 'z' {
		nextIdentifierAsSlice[indexOfLastRune] = nextIdentifierAsSlice[indexOfLastRune] + 1
	} else {
		found := false
		for i := indexOfLastRune - 1; i >= 0; i-- {
			if nextIdentifierAsSlice[i] < 'z' {
				nextIdentifierAsSlice[i]++
				found = true
				break
			}
		}
		if !found {
			nextIdentifierAsSlice = []rune(strings.Repeat("a", len(nextIdentifierAsSlice) + 1))
		}
	}
	nextIdentifier = string(nextIdentifierAsSlice)
	return
}