package client

import (
	"encoding/json"
	"os"
)

type Locale map[string]string

func GetLocalization(lang string) (Locale, error) {
	locale := make(Locale)
	data, err := os.ReadFile("localization/" + lang + ".json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &locale)
	if err != nil {
		return nil, err
	}
	return locale, nil
}

func (l Locale) Get(key string) string {
	val, ok := l[key]
	if ok {
		return val
	} else {
		return key
	}
}
