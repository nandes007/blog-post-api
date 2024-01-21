package jwt

import "strings"

func FormatToken(token string) string {
	return strings.Replace(token, "Bearer ", "", 1)
}
