package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type translatePath struct {
	Path        string `json:"path,omitempty"`
	Key         string `json:"key,omitempty"`
	Instruction string `json:"instruction,omitempty"`
}

type smartling struct{}

func (p *smartling) Marshal(v interface{}, fileFormat string) ([]byte, error) {

	if v != nil && fileFormat == "json" {
		// merge smartling and messages json together ensuring
		// smartling property is at top (required by smartling)
		smartlingBytes, err := smartlingMeta()
		if err != nil {
			return nil, err
		}

		messagesBytes, err := marshalJSONWithoutHTMLEscape(v)
		if err != nil {
			return nil, err
		}

		// trim trailing "}"
		smartlingBytes = smartlingBytes[:len(smartlingBytes)-2]

		// trim beginning "{"
		messagesBytes = messagesBytes[1:]
		buf := bytes.NewBuffer(smartlingBytes)

		// add separating comma between two json objects
		_, err = buf.WriteString(",")
		if err != nil {
			return nil, err
		}

		_, err = buf.Write(messagesBytes)
		if err != nil {
			return nil, err
		}

		return buf.Bytes(), nil
	}

	return nil, fmt.Errorf("smartling Marshal does not support %s format", fileFormat)
}

func smartlingMeta() ([]byte, error) {
	var supportedPluralRules = []string{
		"zero",
		"one",
		"two",
		"few",
		"many",
		"other",
	}

	var translatePaths []translatePath
	for _, rule := range supportedPluralRules {
		var path = translatePath{
			Path:        fmt.Sprintf("/*/%s", rule),
			Key:         fmt.Sprintf("{*}/%s", rule),
			Instruction: "*/description",
		}
		translatePaths = append(translatePaths, path)
	}

	b, err := json.Marshal(translatePaths)
	if err != nil {
		return nil, err
	}

	var smartbuf bytes.Buffer
	_, err = smartbuf.WriteString(`{"smartling": {"translate_paths": `)
	if err != nil {
		return nil, err
	}
	_, err = smartbuf.Write(b)
	if err != nil {
		return nil, err
	}
	_, err = smartbuf.WriteString(`}},`)
	if err != nil {
		return nil, err
	}

	return smartbuf.Bytes(), nil
}

func marshalJSONWithoutHTMLEscape(v interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(v)
	return buffer.Bytes(), err
}
