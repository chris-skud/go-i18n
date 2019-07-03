go-i18n is a Go [package](#package-i18n) and a [command](#command-goi18n) that helps you translate Go programs into multiple languages.

- Supports [pluralized strings](http://cldr.unicode.org/index/cldr-spec/plural-rules) for all 200+ languages in the [Unicode Common Locale Data Repository (CLDR)](http://www.unicode.org/cldr/charts/28/supplemental/language_plural_rules.html).
  - Code and tests are [automatically generated](https://github.com/chris-skud/go-i18n/tree/master/i18n/language/codegen) from [CLDR data](http://cldr.unicode.org/index/downloads).
- Supports strings with embedded variables using [text/template](http://golang.org/pkg/text/template/) syntax.
- Supports message files of any format (e.g. JSON, TOML, YAML).


## Using the i18n package

The i18n package provides support for looking up messages according to a set of locale preferences.

```go
import "github.com/chris-skud/go-i18n/v2/i18n"
```

Create a Bundle when your application starts and load supported translation files.

```go
bundle := i18n.NewBundle(language.English)
bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
bundle.LoadMessageFile("en.toml")
bundle.LoadMessageFile("es.toml")
bundle.LoadMessageFile("es-es.toml")
bundle.LoadMessageFile("en-gb.toml")
```

Create a Localizer set to a given language preference.

```go
func(w http.ResponseWriter, r *http.Request) {
    lang := r.FormValue("lang")
    accept := r.Header.Get("Accept-Language")
    localizer := i18n.NewLocalizer(bundle, lang, accept)
}
```

Use the localizer to retrieve the approriately translated message.
```go
localizer.Localize(&i18n.LocalizeConfig{
    DefaultMessage: &i18n.Message{
        ID: "PersonCats",
        One: "{{.Name}} has {{.Count}} cat.",
        Other: "{{.Name}} has {{.Count}} cats.",
    },
    TemplateData: map[string]interface{}{
        "Name": "Nick",
        "Count": 2,
    },
    PluralCount: 2,
}) // Nick has 2 cats.
```

## Using the goi18n command

The goi18n command manages translation files used by the i18n package.

```
go get -u github.com/chris-skud/go-i18n/v2/goi18n
goi18n -help
```

### Workflows

#### New Project

1. Use `goi18n extract` to create a message file that contains the messages defined in your Go source files (e.g. `goi18n extract -format=json path/to/your/source`)

2. Use `goi18n translate` to create message files for translation (e.g. `goi18n translation -newLangs=en-US,en-GB -format=json active.*.json translate.*.json`). See `goi18n translation -help`
   1. There is also a "provider" flag that enables custom marshalling of translate.*.json files to get them in a format suitable for a given translation provider. Currently only "smartling" is supported (and its very opinionated).

3. Use `goi18n merge` to merge translated message files with your existing "active" message files (e.g. `goi18n merge -format=json active.*.json translate.*.json`).

4. Once you have added new active message files, add them to your bundle like so:
```go
		bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
		bundle.MustLoadMessageFile("active.es.toml")
```

#### Adding new messages

If you have added new messages to your program:

1. Run `goi18n extract` to update `active.{sourceLang}.{format}` with the new messages.
2. Run `goi18n translation -format={format} active.*.{format} translate.*.{format}` to generate updated `translate.*.{format}` files.
3. Translate all the messages in the `translate.*.{format}` files.
4. Run `goi18n merge active.*.{format} translate.*.{format}` to merge the translated messages into the active message files.

## For more information and examples:

- Read the [documentation](http://godoc.org/github.com/chris-skud/go-i18n/v2).
- Look at the [code examples](https://github.com/chris-skud/go-i18n/blob/master/v2/i18n/example_test.go) and [tests](https://github.com/chris-skud/go-i18n/blob/master/i18n/v2/localizer_test.go).
- Look at an example [application](https://github.com/chris-skud/go-i18n/tree/master/v2/example).

## License

go-i18n is available under the MIT license. See the [LICENSE](LICENSE) file for more info.
