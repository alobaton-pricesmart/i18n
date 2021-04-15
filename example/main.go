package main

import (
	"fmt"

	"github.com/alobaton/i18n"
)

var I18N *i18n.I18N

func main() {
	var err error
	I18N = i18n.NewI18N().BindPath("./en.json")
	I18N, err = I18N.BindMainLocale("en")
	if err != nil {
		panic(err)
	}

	I18N, err = I18N.Init()
	if err != nil {
		panic(err)
	}

	result, err := I18N.Lookup("some.awesome.text")
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
