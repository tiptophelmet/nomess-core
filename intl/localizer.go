package intl

import (
	"fmt"
	"sync"

	"github.com/pelletier/go-toml"
)

var lz *localizer

func Init(defaultLocale string) {
	if lz != nil {
		return
	}

	lz = &localizer{locale: defaultLocale}

	loadLocale()
}

func loadLocale() {
	var err error

	lz.localeTree, err = toml.LoadFile(fmt.Sprintf("../../locales/%s.toml", lz.locale))

	if err != nil {
		panic(err)
	}
}

type localizer struct {
	locale     string
	localeTree *toml.Tree
	mu         sync.Mutex
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
