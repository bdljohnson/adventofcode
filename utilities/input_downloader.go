package utilities

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Rather than download everytime, let's cache in the users homedir?

var (
	linkPattern         = "https://adventofcode.com/%v/day/%v/input"
	cacheDir            = "~/.adventofcode/cache/"
	credentialsLocation = "~/.adventofcode/credentials"
)

type InputFetcher struct {
	linkPattern     string
	cacheDir        string
	credentialsFile string
	directToken     bool
	useCache        bool
}

// DownloadInput
func (f *InputFetcher) GetInput(year string, day string) string { // YYYYDD or YYYY, DD?
	url := fmt.Sprintf(f.linkPattern, year, day)
	if f.useCache {
		cacheFile := fmt.Sprintf(cacheDir+"%v/%v/data.txt", year, day)
		f.readThrough(cacheFile, url)
	} else {
		f.download(url)
	}
}

func (f *InputFetcher) readThrough(file string, url string) {
	if checkCache(file) {
		return f.readCache(file)
	} else {
		f.download(url)
	}
	return
}

func (f *InputFetcher) download(url string) {
	return
}

// checkCache is a naive attempt to check the cache.
// No validation is done on the file itself, so an invalid input could be pulled
// from the cache.
func checkCache(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}
	}
	return true
}

func (f *InputFetcher) readCache(cacheFile string) string {
	input, err := ioutil.ReadFile(cacheFile)

	if err != nil {
		log.Fatal("Couldn't read cache file, this should not happen!")
	}
	return string(input)
}

// GetInput with defaults
func GetInput(year string, day string) string {
	downloader := InputFetcher{
		linkPattern:     linkPattern,
		credentialsFile: credentialsLocation,
		cacheDir:        cacheDir,
		directToken:     false,
		useCache:        true,
	}

	downloader.readThrough(year, day)
}
