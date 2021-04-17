package i18n

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/alobaton/i18n/extensions"
	"github.com/alobaton/i18n/langs"
)

// Translate handles directories, default locale and messages.
type Translate struct {
	path       []string
	mainLocale string
	locales    []string
	languages  map[string]interface{}

	initialized bool
}

// NewTranslate create a new Config with default value.
func NewTranslate() *Translate {
	return new(Translate)
}

// BindPath store the config file json path value.
func (t *Translate) BindPath(path string) *Translate {
	t.path = append(t.path, path)
	return t
}

// BindMainLocale store the default language code value.
func (t *Translate) BindMainLocale(mainLocale string) *Translate {
	t.mainLocale = mainLocale
	return t.BindLocale(mainLocale)
}

// BindLocale store a locale code value.
func (t *Translate) BindLocale(locale string) *Translate {
	t.locales = append(t.locales, locale)
	return t
}

// Init store the message to map variable.
func (t *Translate) Init() (*Translate, error) {
	if !langs.IsValid(t.mainLocale) {
		return nil, fmt.Errorf("invalid language %s. supported languages : %v", t.mainLocale, langs.All)
	}

	for _, l := range t.locales {
		if !langs.IsValid(l) {
			return nil, fmt.Errorf("invalid language %s. supported languages : %v", l, langs.All)
		}
	}

	if t.languages == nil {
		t.languages = make(map[string]interface{})
	}

	for _, p := range t.path {
		for _, l := range t.locales {
			for _, ext := range extensions.All {
				path := fmt.Sprintf("%s/%s%s", p, l, ext)

				f, err := os.Open(path)
				if err != nil {
					if os.IsNotExist(err) {
						continue
					}

					return nil, err
				}
				defer f.Close()

				content, err := ioutil.ReadAll(f)
				if err != nil {
					return nil, err
				}

				messages := make(map[string]interface{})
				err = json.Unmarshal(content, &messages)
				if err != nil {
					return nil, err
				}

				t.languages[l] = messages
			}
		}
	}

	if t.IsLoaded() {
		return nil, errors.New("no languages loaded")
	}

	t.initialized = true

	return t, nil
}

// IsLoaded check if at least one language was loaded.
func (t *Translate) IsLoaded() bool {
	for _, v := range t.languages {
		if v != nil {
			m, ok := v.(map[string]interface{})
			if ok {
				if len(m) > 0 {
					return false
				}
			}
		}
	}

	return true
}

// Exists return true and nil if locale exists on Translate struct,
// otherwise return false and an error if applies.
func (t *Translate) Exists(locale string) (bool, error) {
	if !langs.IsValid(locale) {
		return false, fmt.Errorf("invalid language %s. supported languages : %v", locale, langs.All)
	}

	if t.languages[locale] == nil {
		return false, fmt.Errorf("no messages detected for language : %v", locale)
	}

	if len(t.languages[locale].(map[string]interface{})) < 1 {
		return false, nil
	}

	return true, nil
}

// Lookup the destination message based on a message key.
func (t *Translate) Lookup(key string, args ...interface{}) (string, error) {
	if !t.initialized {
		return key, errors.New("Translate instance not initialized")
	}

	locale := t.mainLocale

	return t.LookupWithLocale(locale, key, args...)
}

// LookupWithLocale lookup the destination message based
// on language code and message key.
func (t *Translate) LookupWithLocale(locale, key string, args ...interface{}) (string, error) {
	if !t.initialized {
		return key, errors.New("Translate instance not initialized")
	}

	exists, err := t.Exists(locale)
	if err != nil {
		return key, err
	}

	if !exists {
		return key, fmt.Errorf("locales soen't exists %s. supported languages : %v", locale, langs.All)
	}

	keys := strings.Split(key, ".")

	message, found := lookup(keys[0], keys[1:], t.languages[locale].(map[string]interface{}), args)
	if !found {
		return key, fmt.Errorf("message not found for key : %v", key)
	}

	return message, nil
}

// lookup for a key k in a struct m. Replace args if founded
func lookup(key string, keys []string, m map[string]interface{}, args []interface{}) (string, bool) {
	v := m[key]

	strc, ok := v.(string)
	if ok {
		if len(keys) < 1 {
			if len(args) < 1 {
				return strc, true
			}

			return fmt.Sprintf(strc, args...), true
		}

		return "", false
	}

	mapc, ok := v.(map[string]interface{})
	if !ok {
		return "", false
	}

	if mapc == nil {
		return "", false
	}

	if len(mapc) < 1 {
		return "", false
	}

	return lookup(keys[0], keys[1:], mapc, args)
}
