package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/chris-skud/go-i18n/v2/i18n"
	"github.com/chris-skud/go-i18n/v2/internal"
	"github.com/chris-skud/go-i18n/v2/internal/plural"
	"golang.org/x/text/language"
	yaml "gopkg.in/yaml.v2"
)

func usagetranslation() {
	fmt.Fprintf(os.Stderr, `usage: goi18n translation [options] [message files]

translation generates a updates to all translation files based on the deltas detected between 
the source file and existing active files.

Flags:

	-sourceLanguage tag
		translation messages from this language (e.g. en, en-US, zh-Hant-CN)
 		Default: en

	-outdir directory
		Write message files to this directory.
		Default: .

	-format format
		Output message files in this format.
		Supported formats: json, toml, yaml
		Default: toml
	
	-newLangs tags
		Comma separated list of language "tags" (e.g. en, en-US, zh-Hant-CN)
 		Default: none
`)
}

type translationCommand struct {
	messageFiles   []string
	sourceLanguage languageTag
	outdir         string
	format         string
	newLangs       []string
}

func (tc *translationCommand) name() string {
	return "translation"
}

func (tc *translationCommand) parse(args []string) error {
	flags := flag.NewFlagSet("translation", flag.ExitOnError)
	flags.Usage = usagetranslation

	flags.Var(&tc.sourceLanguage, "sourceLanguage", "en")
	flags.StringVar(&tc.outdir, "outdir", ".", "")
	flags.StringVar(&tc.format, "format", "toml", "")

	var commaSepNewLangs string
	flags.StringVar(&commaSepNewLangs, "newLangs", "", "A comma separated list of any language codes to add new")
	if err := flags.Parse(args); err != nil {
		return err
	}

	if len(commaSepNewLangs) != 0 {
		tc.newLangs = strings.Split(commaSepNewLangs, ",")
	}

	tc.messageFiles = flags.Args()
	return nil
}

func (tc *translationCommand) execute() error {
	if len(tc.messageFiles) < 1 {
		return fmt.Errorf("need at least one message file to parse")
	}

	// if there are new langs to add, create the files, but without overwriting any that currently exist
	for _, loc := range tc.newLangs {

		// ensure tag is valid before doing work
		var err error
		_, err = language.Parse(loc)
		if err != nil {
			log.Printf("failed to parse %s: %s", loc, err.Error())
			continue
		}

		translatePath := fmt.Sprintf("%s/translate.%s.%s", tc.outdir, loc, tc.format)
		_, err = os.Stat(translatePath)
		if os.IsNotExist(err) {
			_, err = os.Create(translatePath)
			if err != nil {
				return err
			}

			// add new file to tc.messageFiles so they get updated too
			tc.messageFiles = append(tc.messageFiles, translatePath)
		}

		activePath := fmt.Sprintf("%s/active.%s.%s", tc.outdir, loc, tc.format)
		_, err = os.Stat(activePath)
		if os.IsNotExist(err) {
			var f *os.File
			f, err = os.Create(activePath)
			if err != nil {
				return err
			}

			_, err = f.WriteString("{}")
			if err != nil {
				return err
			}

			// add new file to tc.messageFiles so they get updated too
			tc.messageFiles = append(tc.messageFiles, activePath)
		}
	}

	inFiles := make(map[string][]byte)
	for _, path := range tc.messageFiles {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		inFiles[path] = content
	}
	ops, err := translation(inFiles, tc.sourceLanguage.Tag(), tc.outdir, tc.format)
	if err != nil {
		return err
	}
	for path, content := range ops.writeFiles {
		if err := ioutil.WriteFile(path, content, 0666); err != nil {
			return err
		}
	}
	for _, path := range ops.deleteFiles {
		// Ignore error since it isn't guaranteed to exist.
		os.Remove(path)
	}
	return nil
}

func translation(messageFiles map[string][]byte, sourceLanguageTag language.Tag, outdir, outputFormat string) (*fileSystemOp, error) {
	untranslationd := make(map[language.Tag][]map[string]*i18n.MessageTemplate)
	sourceMessageTemplates := make(map[string]*i18n.MessageTemplate)
	unmarshalFuncs := map[string]i18n.UnmarshalFunc{
		"json": json.Unmarshal,
		"toml": toml.Unmarshal,
		"yaml": yaml.Unmarshal,
	}
	for path, content := range messageFiles {
		mf, err := i18n.ParseMessageFileBytes(content, path, unmarshalFuncs)
		if err != nil {
			return nil, fmt.Errorf("failed to load message file %s: %s", path, err)
		}
		templates := map[string]*i18n.MessageTemplate{}
		for _, m := range mf.Messages {
			templates[m.ID] = i18n.NewMessageTemplate(m)
		}

		if mf.Tag == sourceLanguageTag {
			for _, template := range templates {
				if sourceMessageTemplates[template.ID] != nil {
					return nil, fmt.Errorf("multiple source translations for id %s", template.ID)
				}
				template.Hash = hash(template)
				sourceMessageTemplates[template.ID] = template
			}
		}
		untranslationd[mf.Tag] = append(untranslationd[mf.Tag], templates)
	}

	if len(sourceMessageTemplates) == 0 {
		return nil, fmt.Errorf("no messages found for source locale %s", sourceLanguageTag)
	}

	pluralRules := plural.DefaultRules()
	all := make(map[language.Tag]map[string]*i18n.MessageTemplate)
	all[sourceLanguageTag] = sourceMessageTemplates
	for _, srcTemplate := range sourceMessageTemplates {
		for dstLangTag, messageTemplates := range untranslationd {
			if dstLangTag == sourceLanguageTag {
				continue
			}
			pluralRule := pluralRules.Rule(dstLangTag)
			if pluralRule == nil {
				// Non-standard languages not supported because
				// we don't know if translations are complete or not.
				continue
			}
			if all[dstLangTag] == nil {
				all[dstLangTag] = make(map[string]*i18n.MessageTemplate)
			}
			dstMessageTemplate := all[dstLangTag][srcTemplate.ID]
			if dstMessageTemplate == nil {
				dstMessageTemplate = &i18n.MessageTemplate{
					Message: &i18n.Message{
						ID:          srcTemplate.ID,
						Description: srcTemplate.Description,
						Hash:        srcTemplate.Hash,
					},
					PluralTemplates: make(map[plural.Form]*internal.Template),
				}
				all[dstLangTag][srcTemplate.ID] = dstMessageTemplate
			}

			// Check all untranslationd message templates for this message id.
			for _, messageTemplates := range messageTemplates {
				untranslationdTemplate := messageTemplates[srcTemplate.ID]
				if untranslationdTemplate == nil {
					continue
				}
				// Ignore empty hashes for v1 backward compatibility.
				if untranslationdTemplate.Hash != "" && untranslationdTemplate.Hash != srcTemplate.Hash {
					// This was translationd from different content so discard.
					continue
				}

				// translation in the translationd messages.
				for pluralForm := range pluralRule.PluralForms {
					dt := untranslationdTemplate.PluralTemplates[pluralForm]
					if dt != nil && dt.Src != "" {
						dstMessageTemplate.PluralTemplates[pluralForm] = dt
					}
				}
			}
		}
	}

	translation := make(map[language.Tag]map[string]*i18n.MessageTemplate)
	active := make(map[language.Tag]map[string]*i18n.MessageTemplate)
	for langTag, messageTemplates := range all {
		active[langTag] = make(map[string]*i18n.MessageTemplate)
		if langTag == sourceLanguageTag {
			active[langTag] = messageTemplates
			continue
		}
		pluralRule := pluralRules.Rule(langTag)
		if pluralRule == nil {
			// Non-standard languages not supported because
			// we don't know if translations are complete or not.
			continue
		}
		for _, messageTemplate := range messageTemplates {
			srcMessageTemplate := sourceMessageTemplates[messageTemplate.ID]
			activeMessageTemplate, translationMessageTemplate := activeDst(srcMessageTemplate, messageTemplate, pluralRule)
			if translationMessageTemplate != nil {
				if translation[langTag] == nil {
					translation[langTag] = make(map[string]*i18n.MessageTemplate)
				}
				translation[langTag][messageTemplate.ID] = translationMessageTemplate
			}
			if activeMessageTemplate != nil {
				active[langTag][messageTemplate.ID] = activeMessageTemplate
			}
		}
	}

	writeFiles := make(map[string][]byte, len(translation)+len(active))
	for langTag, messageTemplates := range translation {
		path, content, err := writeFile(outdir, "translate", langTag, outputFormat, messageTemplates, false)
		if err != nil {
			return nil, err
		}
		writeFiles[path] = content
	}
	deleteFiles := []string{}
	// for langTag, messageTemplates := range active {
	// 	path, content, err := writeFile(outdir, "active", langTag, outputFormat, messageTemplates, langTag == sourceLanguageTag)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if len(content) > 0 {
	// 		writeFiles[path] = content
	// 	} else {
	// 		deleteFiles = append(deleteFiles, path)
	// 	}
	// }
	return &fileSystemOp{writeFiles: writeFiles, deleteFiles: deleteFiles}, nil
}
