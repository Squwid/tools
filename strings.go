package tools

import "fmt"

// JSONLineWrap wraps a single line of json which is great for web responses
func JSONLineWrap(key, value string) string {
	return fmt.Sprintf(`{"%s": "%s"}`, key, value)
}

// WWrap wraps a line in json and returns a byte slice for quick web response
func WWrap(key string, format string, a ...interface{}) []byte {
	return []byte(JSONLineWrap(key, fmt.Sprintf(format, a...)))
}
