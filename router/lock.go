package router

import (
	"github.com/tiptophelmet/nomess-core/logger"
)

var lockedPatterns []string

func LockRoute(pattern string) {
	lockedPatterns = append(lockedPatterns, pattern)
}

func PeekRouteLock(assumed string) {
	for _, locked := range lockedPatterns {
		if matchLockedRoute(locked, assumed) {
			logger.Panic("route reserved for core features & cannot be modified: %v", assumed)
		}
	}
}

func matchLockedRoute(locked string, assumed string) bool {
	return locked == assumed
}
