package i18n

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	i18n := new(Translate)
	return i18n
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
	temp := append(t.locales, locale)
	check := make(map[string]int)
	for _, l := range temp {
		check[l] = 1
	}

	for l := range check {
		t.locales = append(t.locales, l)
	}

	return t
}

// Init store the message to map variable.
func (t *Translate) Init() (*Translate, error) {
	if !IsValid(t.mainLocale) {
		return nil, fmt.Errorf("invalid language %s. supported languages : %v", t.mainLocale, All)
	}

	for _, l := range t.locales {
		if !IsValid(l) {
			return nil, fmt.Errorf("invalid language %s. supported languages : %v", l, All)
		}
	}

	if t.languages == nil {
		t.languages = make(map[string]interface{})
	}

	for _, p := range t.path {
		for _, l := range t.locales {
			// for now only supports json files
			path := fmt.Sprintf("%s/%s%s", p, l, ".json")

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

	if t.loaded() {
		return nil, errors.New("no languages loaded")
	}

	t.initialized = true

	return t, nil
}

// Exists return true and nil if locale exists on Translate struct,
// otherwise return false and an error if applies.
func (t *Translate) Exists(locale string) (bool, error) {
	if !IsValid(locale) {
		return false, fmt.Errorf("invalid language %s. supported languages : %v", locale, All)
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
		return key, fmt.Errorf("locales soen't exists %s. supported languages : %v", locale, All)
	}

	keys := strings.Split(key, ".")

	message, found := lookup(keys[0], keys[1:], t.languages[locale].(map[string]interface{}), args)
	if !found {
		return key, fmt.Errorf("message not found for key : %v", key)
	}

	return message, nil
}

// loaded check if at least one laguage was loaded.
func (t *Translate) loaded() bool {
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
