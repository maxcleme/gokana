package model

type KanaType int

const (
	KanaTypeHiragana KanaType = iota
	KanaTypeKatakana
	KanaTypeBoth
)

func (k KanaType) String() string {
	switch k {
	case KanaTypeHiragana:
		return "Hiragana"
	case KanaTypeKatakana:
		return "Katakana"
	case KanaTypeBoth:
		return "Both"
	default:
		return "Unknown"
	}
}

// Kana represents a Japanese kana character and its romanization
type Kana struct {
	Character string
	Romaji    string
}

// FallingKana represents a kana falling in the game
type FallingKana struct {
	Kana           Kana
	FallPosition   int
	HorizontalPos  int
	ShowingCorrect bool
}

// MainHiragana contains all the main hiragana characters
var MainHiragana = []Kana{
	{"あ", "a"}, {"い", "i"}, {"う", "u"}, {"え", "e"}, {"お", "o"},
	{"か", "ka"}, {"き", "ki"}, {"く", "ku"}, {"け", "ke"}, {"こ", "ko"},
	{"さ", "sa"}, {"し", "shi"}, {"す", "su"}, {"せ", "se"}, {"そ", "so"},
	{"た", "ta"}, {"ち", "chi"}, {"つ", "tsu"}, {"て", "te"}, {"と", "to"},
	{"な", "na"}, {"に", "ni"}, {"ぬ", "nu"}, {"ね", "ne"}, {"の", "no"},
	{"は", "ha"}, {"ひ", "hi"}, {"ふ", "fu"}, {"へ", "he"}, {"ほ", "ho"},
	{"ま", "ma"}, {"み", "mi"}, {"む", "mu"}, {"め", "me"}, {"も", "mo"},
	{"や", "ya"}, {"ゆ", "yu"}, {"よ", "yo"},
	{"ら", "ra"}, {"り", "ri"}, {"る", "ru"}, {"れ", "re"}, {"ろ", "ro"},
	{"わ", "wa"}, {"を", "wo"}, {"ん", "n"},
}

// DakutenHiragana contains hiragana characters with dakuten (゛) marks
var DakutenHiragana = []Kana{
	{"が", "ga"}, {"ぎ", "gi"}, {"ぐ", "gu"}, {"げ", "ge"}, {"ご", "go"},
	{"ざ", "za"}, {"じ", "ji"}, {"ず", "zu"}, {"ぜ", "ze"}, {"ぞ", "zo"},
	{"だ", "da"}, {"ぢ", "di"}, {"づ", "du"}, {"で", "de"}, {"ど", "do"},
	{"ば", "ba"}, {"び", "bi"}, {"ぶ", "bu"}, {"べ", "be"}, {"ぼ", "bo"},
}

// HandakutenHiragana contains hiragana characters with handakuten (゜) marks
var HandakutenHiragana = []Kana{
	{"ぱ", "pa"}, {"ぴ", "pi"}, {"ぷ", "pu"}, {"ぺ", "pe"}, {"ぽ", "po"},
}

// MainKatakana contains all the main katakana characters
var MainKatakana = []Kana{
	{"ア", "a"}, {"イ", "i"}, {"ウ", "u"}, {"エ", "e"}, {"オ", "o"},
	{"カ", "ka"}, {"キ", "ki"}, {"ク", "ku"}, {"ケ", "ke"}, {"コ", "ko"},
	{"サ", "sa"}, {"シ", "shi"}, {"ス", "su"}, {"セ", "se"}, {"ソ", "so"},
	{"タ", "ta"}, {"チ", "chi"}, {"ツ", "tsu"}, {"テ", "te"}, {"ト", "to"},
	{"ナ", "na"}, {"ニ", "ni"}, {"ヌ", "nu"}, {"ネ", "ne"}, {"ノ", "no"},
	{"ハ", "ha"}, {"ヒ", "hi"}, {"フ", "fu"}, {"ヘ", "he"}, {"ホ", "ho"},
	{"マ", "ma"}, {"ミ", "mi"}, {"ム", "mu"}, {"メ", "me"}, {"モ", "mo"},
	{"ヤ", "ya"}, {"ユ", "yu"}, {"ヨ", "yo"},
	{"ラ", "ra"}, {"リ", "ri"}, {"ル", "ru"}, {"レ", "re"}, {"ロ", "ro"},
	{"ワ", "wa"}, {"ヲ", "wo"}, {"ン", "n"},
}

// DakutenKatakana contains katakana characters with dakuten (゛) marks
var DakutenKatakana = []Kana{
	{"ガ", "ga"}, {"ギ", "gi"}, {"グ", "gu"}, {"ゲ", "ge"}, {"ゴ", "go"},
	{"ザ", "za"}, {"ジ", "ji"}, {"ズ", "zu"}, {"ゼ", "ze"}, {"ゾ", "zo"},
	{"ダ", "da"}, {"ヂ", "di"}, {"ヅ", "du"}, {"デ", "de"}, {"ド", "do"},
	{"バ", "ba"}, {"ビ", "bi"}, {"ブ", "bu"}, {"ベ", "be"}, {"ボ", "bo"},
	{"ヴ", "vu"},
}

// HandakutenKatakana contains katakana characters with handakuten (゜) marks
var HandakutenKatakana = []Kana{
	{"パ", "pa"}, {"ピ", "pi"}, {"プ", "pu"}, {"ペ", "pe"}, {"ポ", "po"},
}

// GetKanaSet returns the appropriate kana slice based on the selected type and dakuten setting
func GetKanaSet(kanaType KanaType, includeDakuten bool) []Kana {
	switch kanaType {
	case KanaTypeHiragana:
		if includeDakuten {
			combined := make([]Kana, 0, len(MainHiragana)+len(DakutenHiragana)+len(HandakutenHiragana))
			combined = append(combined, MainHiragana...)
			combined = append(combined, DakutenHiragana...)
			combined = append(combined, HandakutenHiragana...)
			return combined
		}
		return MainHiragana
	case KanaTypeKatakana:
		if includeDakuten {
			combined := make([]Kana, 0, len(MainKatakana)+len(DakutenKatakana)+len(HandakutenKatakana))
			combined = append(combined, MainKatakana...)
			combined = append(combined, DakutenKatakana...)
			combined = append(combined, HandakutenKatakana...)
			return combined
		}
		return MainKatakana
	case KanaTypeBoth:
		if includeDakuten {
			combined := make([]Kana, 0, len(MainHiragana)+len(DakutenHiragana)+len(HandakutenHiragana)+len(MainKatakana)+len(DakutenKatakana)+len(HandakutenKatakana))
			combined = append(combined, MainHiragana...)
			combined = append(combined, DakutenHiragana...)
			combined = append(combined, HandakutenHiragana...)
			combined = append(combined, MainKatakana...)
			combined = append(combined, DakutenKatakana...)
			combined = append(combined, HandakutenKatakana...)
			return combined
		}
		combined := make([]Kana, 0, len(MainHiragana)+len(MainKatakana))
		combined = append(combined, MainHiragana...)
		combined = append(combined, MainKatakana...)
		return combined
	default:
		return MainHiragana
	}
}
