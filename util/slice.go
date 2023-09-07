package util

func GetNonIntersecting(main []string, compared []string) []string {
	m := make(map[string]struct{}, len(main))
	for _, item := range main {
		m[item] = struct{}{}
	}

	nonIntersecting := make([]string, 0, len(compared))
	for _, item := range compared {
		if _, ok := m[item]; !ok {
			nonIntersecting = append(nonIntersecting, item)
		}
	}

	return nonIntersecting
}
