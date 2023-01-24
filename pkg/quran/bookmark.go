package quran

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadBookmark(path string) (*Bookmark, error) {
	// check the extension type
	ext := filepath.Ext(path)
	if ext != ".json" {
		return nil, errors.New("Unsupported bookmark filetype!")
	}

	bookmark := Bookmark{}
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(f, &bookmark)

	return &bookmark, nil
}

// create a default json bookmark file with some default data
// the default path is the root of the project
func MakeBookmark(path string) {
	defaultData := `{"currentPage" : 1,"savedPages" : [1, 2, 3],"savedAyaht" : []}`
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		f, _ := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		fmt.Fprintln(f, defaultData)
	}
}
