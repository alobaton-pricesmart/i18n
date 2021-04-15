# i18n

> :warning: **Still under construction**

simple i18n support that relies on standard go libraries

## How to start?

The i18n package mainly includes a set of methods for managing the data. Start by creating a new i18n instance.

```bash
i18n := NewI18N().BindPath("./example/en.json")
i18n, err := i18n.BindMainLocale("en")
if err != nil {
	...
}

i18n, err = i18n.Init()
if err != nil {
	...
}
```

Once you setup the i18n instance, you must be able to lookup for messages.
```bash
result, err := i18n.Lookup("some.awesome.text")
if err != nil {
    ...
}
fmt.Println(result)
```

## How to install?

```bash
go get github.com/alobaton/i18n
```

## How to test?

```bash
$ go test i18n/...
```
