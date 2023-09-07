package locales

import "strings"

var supportedLocales []string

func Register(locales []string) {
	supportedLocales = locales
}

func IsSupportedLocale(tag string) bool {
	for _, lang := range supportedLocales {
		if strings.HasPrefix(tag, lang) {
			return true
		}
	}
	return false
}
