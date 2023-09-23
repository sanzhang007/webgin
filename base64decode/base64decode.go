package base64decode

import (
	"encoding/base64"
	"strings"
)

func Base64Decode(str string) string {
	// str := "SGVsbG8sIHdvcmxkIQ=="
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(str)))
	n, err := base64.StdEncoding.Decode(dst, []byte(str))
	if err != nil {
		// fmt.Printf("err: %v\n", err)
		return str
	}
	dst = dst[:n]
	return string(dst)
}

func Base64Decode_(str string) string {
	s := strings.Split(str, "_")
	var builder strings.Builder
	for _, item := range s {
		builder.WriteString(Base64Decode(item))
	}
	return builder.String()
}
