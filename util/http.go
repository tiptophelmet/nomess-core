package util

import "net/http"

func GetHeaderMaps(w http.ResponseWriter) (singleValHeaders map[string]string, multiValHeaders map[string][]string) {
	singleValHeaders = make(map[string]string)
	multiValHeaders = make(map[string][]string)

	for k, v := range w.Header() {
		if len(v) == 1 {
			singleValHeaders[k] = v[0]
		} else if len(v) > 1 {
			multiValHeaders[k] = v
		}
	}

	return singleValHeaders, multiValHeaders
}
