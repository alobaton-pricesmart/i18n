package i18n

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// I18N
type I18N struct {
	path       []string
	mainLocale string
	languages  map[string]interface{}
}

// Create a new Config with default value.
func NewI18N() *I18N {
	i18n := new(I18N)
	return i18n
}

// Store the config file json path value.
func (i *I18N) BindPath(path string) *I18N {
	i.path = append(i.path, path)
	return i
}

// Store the default language code value.
func (i *I18N) BindMainLocale(mainLocale string) (*I18N, error) {
	if !IsValid(mainLocale) {
		return nil, fmt.Errorf("invalid language %s. supported languages : %v", mainLocale, All)
	}

	i.mainLocale = mainLocale
	return i, nil
}

// Store the message to map variable.
func (i *I18N) Init() (*I18N, error) {
	if i.languages == nil {
		i.languages = make(map[string]interface{})
	}

	for _, p := range i.path {
		f, err := os.Open(p)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		_, file := filepath.Split(f.Name())
		ext := filepath.Ext(file)
		file = strings.Replace(file, ext, "", 1)
		if !IsValid(file) {
			return nil, fmt.Errorf("invalid language %s. supported languages : %v", file, All)
		}

		content, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}

		messages := make(map[string]interface{})
		err = json.Unmarshal(content, &messages)
		if err != nil {
			return nil, err
		}

		i.languages[file] = messages
	}

	return i, nil
}

func (i *I18N) Exists(locale string) (bool, error) {
	if !IsValid(locale) {
		return false, fmt.Errorf("invalid language %s. supported languages : %v", locale, All)
	}

	if len(i.languages[locale].(map[string]interface{})) < 1 {
		return false, nil
	}

	return true, nil
}

func (i *I18N) Lookup(key string, args ...interface{}) (string, error) {
	locale := i.mainLocale

	return i.LookupWithLocale(locale, key, args...)
}

// Lookup the destination message based on language code & message.
func (i *I18N) LookupWithLocale(locale, key string, args ...interface{}) (string, error) {
	exists, err := i.Exists(locale)
	if err != nil {
		return "", err
	}

	if !exists {
		return "", fmt.Errorf("locales soen't exists %s. supported languages : %v", locale, All)
	}

	keys := strings.Split(key, ".")

	message, found := lookup(keys[0], keys[1:], i.languages[locale].(map[string]interface{}), args)
	if !found {
		return key, nil
	}

	return message, nil
}

func lookup(key string, keys []string, m map[string]interface{}, args []interface{}) (string, bool) {
	v := m[key]

	strc, ok := v.(string)
	if ok {
		if len(keys) < 1 {
			if len(args) < 1 {
				return strc, true
			} else {
				return fmt.Sprintf(strc, args...), true
			}
		} else {
			return "", false
		}
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
