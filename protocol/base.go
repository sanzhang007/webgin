package protocol

import (
	"strings"

	"github.com/sanzhang007/webgin/utils"
)

type Baser interface {
	Parse(string)
}

func ClashParse(link *string) (Baser, error) {
	matches, err := utils.CheckLink(*link)
	var ctx Baser
	if err != nil {
		return nil, err
	}
	switch strings.ToLower(matches[1]) {
	case "vmess":
		ctx = new(Vmess)
		ctx.Parse(*link)
	case "trojan":
		ctx = new(Trojan)
		ctx.Parse(*link)
	case "ss":
		ctx = new(Ss)
		ctx.Parse(*link)
	case "ssr":
		// fmt.Println(*link)

		ctx = new(Ssr)
		ctx.Parse(*link)

	}

	return ctx, nil
}
