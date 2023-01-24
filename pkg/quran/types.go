package quran

type Surah struct {
	NumberSurah int    `json:"number"`
	NameAr      string `json:"name"`
	NameEn      string `json:"englishName"`
	NumberAyaht int    `json:"numberOfAyahs"`
	Ayaht       []Ayah `json:"ayahs"`
}

type Ayah struct {
	Text          string `json:"text"`
	NumberInSurah int    `json:"numberInSurah"`
	NumberInQuran int    `json:"number"`
	Juz           int    `json:"juz"`
	Verse         int    `json:"chapter"`
	Page          int    `json:"page"`
}

type Bookmark struct {
	CurrentPage int   `json:"currentPage"`
	SavedPages  []int `json:"savedPages"`
	SavedAyaht  []int `json:"savedAyaht"`
}
