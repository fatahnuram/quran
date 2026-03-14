package data

type Surah struct {
	No        string `json:"no"`
	NameAR    string `json:"name_ar"`
	NameID    string `json:"name_id"`
	AyatCount string `json:"ayat_count"`
	Ayats     []Ayat `json:"-"`
}

type Ayat struct {
	SurahNo string `json:"surah_no"`
	AyatNo  string `json:"ayat_no"`
	AR      string `json:"ar"`
	ID      string `json:"id"`
}

var Quran = []Surah{
	{
		No:        "1",
		NameAR:    "abce",
		NameID:    "asdasd",
		AyatCount: "11",
		Ayats: []Ayat{
			{
				SurahNo: "1",
				AyatNo:  "1",
				AR:      "11asdasd",
				ID:      "11asdasd",
			},
			{
				SurahNo: "1",
				AyatNo:  "2",
				AR:      "11asdasd",
				ID:      "11asdasd",
			},
		},
	},
	{
		No:        "2",
		NameAR:    "2abce",
		NameID:    "2asdasd",
		AyatCount: "22",
		Ayats: []Ayat{
			{
				SurahNo: "2",
				AyatNo:  "1",
				AR:      "2asdasd",
				ID:      "2asdasd",
			},
			{
				SurahNo: "2",
				AyatNo:  "2",
				AR:      "2asdasd",
				ID:      "2asdasd",
			},
		},
	},
}
