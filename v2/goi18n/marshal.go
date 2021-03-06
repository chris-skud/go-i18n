package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/chris-skud/go-i18n/v2/goi18n/providers"
	"github.com/chris-skud/go-i18n/v2/i18n"
	"github.com/chris-skud/go-i18n/v2/internal/plural"
	"golang.org/x/text/language"
	yaml "gopkg.in/yaml.v2"
)

func writeFile(outdir, label string, langTag language.Tag, format string, messageTemplates map[string]*i18n.MessageTemplate, sourceLanguage bool, provider providers.Provider) (path string, content []byte, err error) {
	v := marshalValue(messageTemplates, sourceLanguage)
	if provider != nil {
		content, err = provider.Marshal(v, format)
	} else {
		content, err = marshal(v, format)
	}
	if err != nil {
		return "", nil, fmt.Errorf("failed to marshal %s strings to %s: %s", langTag, format, err)
	}
	path = filepath.Join(outdir, fmt.Sprintf("%s.%s.%s", label, langTag, format))

	// hack to strip extra backslash json encoder adds to '\n' in json string
	content = bytes.Replace(content, []byte(`\\n`), []byte(`\n`), -1)

	return path, content, nil
}

func marshalValue(messageTemplates map[string]*i18n.MessageTemplate, sourceLanguage bool) interface{} {
	v := make(map[string]interface{}, len(messageTemplates))
	for id, template := range messageTemplates {
		if other := template.PluralTemplates[plural.Other]; sourceLanguage && len(template.PluralTemplates) == 1 &&
			other != nil && template.Description == "" && template.LeftDelim == "" && template.RightDelim == "" {
			v[id] = other.Src
		} else {
			m := map[string]string{}
			if template.Description != "" {
				m["description"] = template.Description
			}
			if !sourceLanguage {
				m["hash"] = template.Hash
			}
			for pluralForm, template := range template.PluralTemplates {
				m[string(pluralForm)] = template.Src
			}
			v[id] = m
		}
	}
	return v
}

func marshalJSONWithoutHTMLEscape(v interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(v)
	return buffer.Bytes(), err
}

func marshal(v interface{}, format string) ([]byte, error) {
	switch format {
	case "json":
		return marshalJSONWithoutHTMLEscape(v)
	case "toml":
		var buf bytes.Buffer
		enc := toml.NewEncoder(&buf)
		enc.Indent = ""
		err := enc.Encode(v)
		return buf.Bytes(), err
	case "yaml":
		return yaml.Marshal(v)
	}
	return nil, fmt.Errorf("unsupported format: %s", format)
}
