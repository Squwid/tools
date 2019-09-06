package tools

import "fmt"

// JSONLineWrap wraps a single line of json which is great for web responses
func JSONLineWrap(key, value string) string {
	return fmt.Sprintf(`{"%s": "%s"}`, key, value)
}
