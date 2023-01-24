package quran

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	SURAH_NUM = 114 // the number of surahs in the holy quran
	LANG      = 1   // AR -> 0, EN -> 1
)

// get all the ayaht in a particular page by scanning the pages for a certain page number
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
