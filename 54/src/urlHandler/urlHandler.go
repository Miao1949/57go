package urlHandler

import (
	"strings"
	"dao"
)

func GetAllUrlMappings() (urlMappings []*dao.UrlMapping, err error){
	return dao.GetAllUrlMappings()
}

func GetAllShortUrls() (shortUrls []string, err error) {
	return dao.GetAllShortUrls()
}

func GetShortUrlForLongUrl(longUrl string) (shortUrl string, err error) {
	return dao.GetShortUrlForLongUrl(longUrl)
}

func GetLongUrlForShortUrl(shortUrl string) (longUrl string, err error) {
	return dao.GetLongUrlForShortUrl(shortUrl)
}

func AddMapping(shortUrl string, longUrl string) (err error) {
	return dao.AddMapping(shortUrl, longUrl)
}

func GetNumberOfUrlInvocations(shortUrl string) (numberOfInvocations int, err error) {
	return dao.GetNumberOfUrlInvocations(shortUrl)
}

func ShortUrlInvoked(shortUrl string) (err error) {
	return dao.ShortUrlInvoked(shortUrl)
}

func MappingExists(longUrl string) (found bool, err error) {
	return dao.MappingExists(longUrl)
}

func GenerateAndAdd(longUrl string) (shortUrl string, err error) {
	found, err := MappingExists(longUrl)
	if err != nil {
		return
	}

	if found {
		return dao.GetShortUrlForLongUrl(longUrl)
	}

	shortUrls, err := dao.GetAllShortUrls()

	if err != nil {
		return
	}

	if len(shortUrls) == 0 {
		shortUrl ="a"
	} else {
		existingShortUrlsSortedByLength := SortSliceByLength(shortUrls)
		shortUrl = NextIdentifierAfter(existingShortUrlsSortedByLength[len(existingShortUrlsSortedByLength) -1])
	}

	err = dao.AddMapping(shortUrl, longUrl)

	return
}

func NextIdentifierAfter(s string) (nextIdentifier string) {
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


func SortSliceByLength(sliceToSort []string) (sortedSlice []string) {
	sortedSlice = make([]string, len(sliceToSort))
	for _, stringToInsert := range sliceToSort {
		insertionIndex := 0
		for indexOfSortedElement, sortedElement := range sortedSlice {
			if len(stringToInsert) < len(sortedElement) || len(sortedElement) == 0 {
				insertionIndex = indexOfSortedElement
				break
			}
		}

		copy(sortedSlice[insertionIndex + 1 :], sortedSlice[insertionIndex: len(sortedSlice) - 1])
		sortedSlice[insertionIndex] = stringToInsert
	}

	return sortedSlice
}
