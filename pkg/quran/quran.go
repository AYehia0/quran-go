package quran

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	SURAH_NUM = 114 // the number of surahs in the holy quran
	PAGES     = 604 // the number of pages in the holy quran
	LANG      = 1   // AR -> 0, EN -> 1
)

// get all the ayaht in a particular page by scanning the pages for a certain page number
// TODO: make it map for faster access
func AyahtInPage(surahs *[][]Surah, pageNum int) []Ayah {
	ayaht := make([]Ayah, 0)
	for _, surah := range *surahs {
		for _, ayah := range surah[LANG].Ayaht {
			if ayah.Page == pageNum {
				ayaht = append(ayaht, ayah)
			}
		}
	}
	return ayaht
}

// return a map of page number with all the ayaht inside
func AyahtInPages(surahs *[][]Surah) map[int][]Ayah {
	ayahtMap := make(map[int][]Ayah, 0)
	for _, surah := range *surahs {
		for _, ayah := range surah[LANG].Ayaht {
			ayahtMap[ayah.Page] = append(ayahtMap[ayah.Page], ayah)
		}
	}
	return ayahtMap
}

// return the left and right pages. (as a list)
func GetPages(ayaht map[int][]Ayah, currentPage int) ([]Ayah, []Ayah) {

	var leftPageNum, rightPageNum int

	// check the boundries
	// TODO: handle this !!
	if currentPage > PAGES || currentPage < 1 {
		return nil, nil
	}

	// check the page order
	// if pagenum is even : left
	if currentPage%2 == 0 {
		leftPageNum = currentPage
		rightPageNum = currentPage - 1
	} else {
		// if pagenum is odd : right
		rightPageNum = currentPage
		leftPageNum = currentPage + 1
	}

	return ayaht[leftPageNum], ayaht[rightPageNum]

}

func ParseQuranData(dataFilePath string) *[][]Surah {

	// the first index for surah number/order in mushaf, the second is either 0,1 (ar, en)
	chapters := [][]Surah{}

	f, err := os.Open(dataFilePath)
	if err != nil {
		fmt.Printf("Error reading the data! : %v\n", err)
		os.Exit(1)
	}

	data := json.NewDecoder(f)
	if err := data.Decode(&chapters); err != nil {
		fmt.Printf("Error decoding the file! : %v\n", err)
		os.Exit(1)
	}
	return &chapters

}
