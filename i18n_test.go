package i18n

import (
	"testing"
)

func TestLookupMessageFound(t *testing.T) {
	i18n := NewI18N().BindPath("./example/en.json")
	i18n, err := i18n.BindMainLocale("en")
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	i18n, err = i18n.Init()
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	expectedResult := "Hello World!"
	result, err := i18n.Lookup("some.awesome.text")
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	if result != expectedResult {
		t.Errorf("Should return %s, got %s", expectedResult, result)
	}
}

func TestLookupMessageFoundWithArgs(t *testing.T) {
	i18n := NewI18N().BindPath("./example/en.json")
	i18n, err := i18n.BindMainLocale("en")
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	i18n, err = i18n.Init()
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	expectedResult := "Hello i18n!"
	result, err := i18n.Lookup("some.awesome.textWithArgs", "i18n")
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	if result != expectedResult {
		t.Errorf("Should return %s, got %s", expectedResult, result)
	}
}

func TestLookupMessageNotFound(t *testing.T) {
	i18n := NewI18N().BindPath("./example/en.json")
	i18n, err := i18n.BindMainLocale("en")
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	i18n, err = i18n.Init()
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	expectedResult := "not.found"
	result, err := i18n.Lookup("not.found")
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	if result != expectedResult {
		t.Errorf("Should return %s, got %s", expectedResult, result)
	}
}

func TestLookupMessageNotFoundMoreLevels(t *testing.T) {
	i18n := NewI18N().BindPath("./example/en.json")
	i18n, err := i18n.BindMainLocale("en")
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	i18n, err = i18n.Init()
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	expectedResult := "some.awesome.text.level"
	result, err := i18n.Lookup("some.awesome.text.level")
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	if result != expectedResult {
		t.Errorf("Should return %s, got %s", expectedResult, result)
	}
}
