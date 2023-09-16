package intl

import (
	"embed"
	"fmt"
	"sync"

	"github.com/pelletier/go-toml"
)

var lz *localizer

func Init(defaultLocale string, localesDir embed.FS) {
	if lz != nil {
		return
	}

	lz = &localizer{locale: defaultLocale, localesDir: localesDir}

	loadLocale()
}

type localizer struct {
	locale     string
	localeTree *toml.Tree
	localesDir embed.FS
	mu         sync.Mutex
}

func loadLocale() {
	var err error

	localeFileBytes, err := lz.localesDir.ReadFile(fmt.Sprintf("locales/%s.toml", lz.locale))
	if err != nil {
		panic(err)
	}

	lz.localeTree, err = toml.LoadBytes(localeFileBytes)

	if err != nil {
		panic(err)
	}
}

func SetLocale(locale string) {
	lz.mu.Lock()
	defer lz.mu.Unlock()

	lz.locale = locale
	loadLocale()
}

func GetLocale() string {
	return lz.locale
}

func Localize(key string) string {
	if value := lz.localeTree.Get(key); value != nil {
		if s, ok := value.(string); ok {
			return s
		}
	}
	return ""
}
