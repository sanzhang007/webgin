package protocol

import (
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/sanzhang007/webgin/base64decode"

	"github.com/xxf098/lite-proxy/utils"
)

type Ssr struct {
	Id            int    `gorm:"primaryKey"`
	Name          string `yaml:"name"`
	Password      string `yaml:"password"`
	Port          string `yaml:"port"`
	Server        string `yaml:"server" gorm:"unique"`
	Cipher        string `yaml:"cipher"`
	Type          string `yaml:"type"`
	Protocol      string `yaml:"protocol"`
	Obfs          string `yaml:"obfs"`
	ObfsParam     string `yaml:"obfs-param"`
	ProtocolParam string `yaml:"protocol-param"`
	Link          string
}

func (ssr *Ssr) Parse(urlstring string) {
	builder := new(strings.Builder)
	linkSplit := strings.Split(urlstring, "://")
	// ctx := base64decode.Base64Decode(linkSplit[1])
	// ctx := base64decode.Base64Decode_(linkSplit[1])

	ctx, _ := utils.DecodeB64(linkSplit[1])
	builder.WriteString(linkSplit[0])
	builder.WriteString("://")
	r := regexp.MustCompile(`.*:\d+(:.*)/\?`)
	matches := r.FindAllStringSubmatch(ctx, -1)
	if len(matches) == 0 {
		log.Println("ssr error")
		return
	}
	ctx2 := matches[0][1]
	ctx3 := strings.Split(ctx2, ":")
	ctx = strings.Replace(ctx, ctx2, "", 1)
	builder.WriteString(ctx)
	u, err := url.Parse(builder.String())
	if err != nil {
		return
	}
	query := u.Query()
	ssr.Type = "ssr"
	ssr.Name = base64decode.Base64Decode(query.Get("remarks"))
	ssr.Name = "UnkownName"
	ssr.Server = u.Hostname()
	ssr.Port = u.Port()
	ssr.Protocol = base64decode.Base64Decode(ctx3[1])
	ssr.Cipher = base64decode.Base64Decode(ctx3[2])
	ssr.Obfs = base64decode.Base64Decode(ctx3[3])
	ssr.Password = base64decode.Base64Decode(ctx3[4])
	ssr.ObfsParam = base64decode.Base64Decode(query.Get("obfsparam"))
	ssr.ProtocolParam = base64decode.Base64Decode(query.Get("protoparam"))
	ssr.ObfsParam = ""
	ssr.ProtocolParam = ""
	ssr.Link = urlstring

}
