package i18n

const (
	CA = "ca" // CA Catalan
	DE = "de" // DE German
	EN = "en" // EN English
	ES = "es" // ES Spanish
	EU = "eu" // EU Basque
	FR = "fr" // FR French
	IT = "it" // IT Italian
	JA = "ja" // JA Japanese
	PT = "pt" // PT Portuguese
	RU = "ru" // RU Russian
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
