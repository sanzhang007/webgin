package test

import (
	"fmt"
	"testing"

	"github.com/sanzhang007/webgin/base64decode"
)

func TestBase64Decode(t *testing.T) {
	s := base64decode.Base64Decode("SGVsbG8sIHdvcmxkIQ==")
	fmt.Println(s)
}
