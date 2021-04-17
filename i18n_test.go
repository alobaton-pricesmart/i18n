package i18n

import (
	"testing"
)

func TestInitPathNotFound(t *testing.T) {
	_, err := NewTranslate().BindPath("./not-found-path").BindMainLocale("en").Init()
	if err == nil {
		t.Errorf("Should return error")
	}
}

func TestLookupMessageFound(t *testing.T) {
	translate, err := NewTranslate().BindPath("./example").BindMainLocale("en").Init()
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	expectedResult := "Hello World!"
	result, err := translate.Lookup("some.awesome.text")
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	if result != expectedResult {
		t.Errorf("Should return %s, got %s", expectedResult, result)
	}
}

func TestLookupWithLocaleMessageFound(t *testing.T) {
	translate, err := NewTranslate().BindPath("./example").BindMainLocale("en").BindLocale("es").Init()
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	expectedResult := "Hola Mundo!"
	result, err := translate.LookupWithLocale("es", "some.awesome.text")
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	if result != expectedResult {
		t.Errorf("Should return %s, got %s", expectedResult, result)
	}
}

func TestLookupFoundWithArgs(t *testing.T) {
	translate, err := NewTranslate().BindPath("./example").BindMainLocale("en").Init()
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	translate, err = translate.Init()
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	expectedResult := "Hello translate!"
	result, err := translate.Lookup("some.awesome.textWithArgs", "translate")
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	if result != expectedResult {
		t.Errorf("Should return %s, got %s", expectedResult, result)
	}
}

func TestLookupWithLocaleFoundWithArgs(t *testing.T) {
	translate, err := NewTranslate().BindPath("./example").BindMainLocale("en").BindLocale("es").Init()
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	translate, err = translate.Init()
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	expectedResult := "Hola translate!"
	result, err := translate.LookupWithLocale("es", "some.awesome.textWithArgs", "translate")
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	if result != expectedResult {
		t.Errorf("Should return %s, got %s", expectedResult, result)
	}
}

func TestLookupNotFound(t *testing.T) {
	translate, err := NewTranslate().BindPath("./example").BindMainLocale("en").Init()
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	translate, err = translate.Init()
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	expectedResult := "not.found"
	result, err := translate.Lookup("not.found")
	if err == nil {
		t.Errorf("Should return error")
	}

	if result != expectedResult {
		t.Errorf("Should return %s, got %s", expectedResult, result)
	}
}

func TestLookupNotFoundMoreLevels(t *testing.T) {
	translate, err := NewTranslate().BindPath("./example").BindMainLocale("en").Init()
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	translate, err = translate.Init()
	if err != nil {
		t.Errorf("Shouldn't return error, got %v", err)
	}

	expectedResult := "some.awesome.text.level"
	result, err := translate.Lookup("some.awesome.text.level")
	if err == nil {
		t.Errorf("Should return error")
	}

	if result != expectedResult {
		t.Errorf("Should return %s, got %s", expectedResult, result)
	}
}
