package main

import (
	"fmt"

	"github.com/alobaton/i18n"
)

var translate *i18n.Translate

func main() {
	var err error
	translate, err = i18n.NewTranslate().BindPath("./example").BindMainLocale("en").Init()
	if err != nil {
		panic(err)
	}

	result, err := translate.Lookup("some.awesome.text")
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
