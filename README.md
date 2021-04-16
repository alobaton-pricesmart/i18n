# i18n

[![Build Status](https://travis-ci.org/alobaton/i18n.svg?branch=main)](https://travis-ci.org/alobaton/i18n) [![Go Report Card](https://goreportcard.com/badge/github.com/alobaton/i18n)](https://goreportcard.com/report/github.com/alobaton/i18n)

Simple `i18n` support that relies on standard go libraries

## How to start?

The `i18n` package mainly includes a set of methods for managing the data. Start by creating a `en.json` file.
```bash
{
    "some": {
        "awesome": {
            "text": "Hello World!"

        }
    }
}
```
Create a new `Translate` instance as follows.
```bash
var err error
translate, err = i18n.NewTranslate().BindPath("./example").BindMainLocale("en").Init()
if err != nil {
	...
}
```

Once you setup the i18n instance, you should be able to lookup for messages.
```bash
result, err := translate.Lookup("some.awesome.text")
if err != nil {
    ...
}
fmt.Println(result)
```

The program should print `Hello World!`

### Lookup for a specific locale

```bash
result, err := translate.LookupWithLocale("en", "some.awesome.text")
if err != nil {
    ...
}
fmt.Println(result)
```

The program should print `Hello World!`

### Lookup with arguments

`i18n` relies on `fmt.Sprintf(...)` to apply replacements, so you should be able to use it as follows.

Your `.json` file should look like this. 
```bash
{
    "some": {
        "awesome": {
            "textWithArgs": "Hello %s!"
        }
    }
}
```

Lookup for messages like this.
```bash
result, err := translate.Lookup("some.awesome.textWithArgs", "i18n")
if err != nil {
    ...
}
fmt.Println(result)
```

The program should print `Hello i18n!`

## JSON Format

```bash
{
    "some": {
        "awesome": {
            "text": "Hello World!",
            "textWithArgs": "Hello %s!"
        }
    }
}
```

## How to install?

```bash
go get github.com/alobaton/i18n
```

## How to test?

```bash
$ go test ./...
```

## Example

Here you can find an [example](https://github.com/alobaton/i18n/blob/main/example/main.go)
