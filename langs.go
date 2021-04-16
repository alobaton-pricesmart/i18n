package i18n

const (
	// CA Catalan
	CA = "ca"
	// DE German
	DE = "de"
	// EN English
	EN = "en"
	// ES Spanish
	ES = "es"
	// EU Basque
	EU = "eu"
	// FR French
	FR = "fr"
	// IT Italian
	IT = "it"
	// JA Japanese
	JA = "ja"
	// PT Portuguese
	PT = "pt"
	// RU Russian
	RU = "ru"
)

// All supported languages
var All = []string{
	CA,
	DE,
	EN,
	ES,
	EU,
	FR,
	IT,
	JA,
	PT,
	RU,
}

var native = map[string]string{
	"CA": "Català",
	"DE": "Deutsch",
	"EN": "English",
	"ES": "Español",
	"EU": "Euskera",
	"FR": "Français",
	"IT": "Italiano",
	"JA": "日本語",
	"PT": "Portugues",
	"RU": "русский",
}

// IsValid true if lang is a valid language, false otherwise.
func IsValid(lang string) bool {
	for _, l := range All {
		if l == lang {
			return true
		}
	}
	return false
}

// NativeName return the lang native name
func NativeName(lang string) string {
	return native[lang]
}
