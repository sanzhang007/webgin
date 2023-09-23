package protocol

import (
	"log"
	"net"
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

func (ssr *Ssr) ParseTemp(urlstring string) {
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

// var (
// 	NotSSRLink error = errors.New("not a shadowsocksR link")
// )

func (ssr *Ssr) Parse(link string) {
	regex := regexp.MustCompile(`^ssr://([A-Za-z0-9+-=/_]+)`)
	res := regex.FindAllStringSubmatch(link, 1)
	b64 := ""
	if len(res) > 0 && len(res[0]) > 1 {
		b64 = res[0][1]
	}
	uri, err := utils.DecodeB64(b64)
	if err != nil {
		return
	}
	parts := strings.SplitN(uri, "/?", 2)
	links := strings.Split(parts[0], ":")
	if len(links) != 6 || len(parts) != 2 {
		return
	}
	port := links[1]
	pass, err := utils.DecodeB64(links[5])
	if err != nil {
		return
	}
	cipher := links[3]
	if cipher == "none" {
		cipher = "dummy"
	}
	ssr.Type = "ssr"
	ssr.Name = ""
	ssr.Server = links[0]
	ssr.Port = port
	ssr.Protocol = links[2]
	ssr.Cipher = cipher
	ssr.Obfs = links[4]
	ssr.Password = pass
	ssr.ObfsParam = ""
	ssr.ProtocolParam = ""
	ssr.Link = link
	query := strings.ReplaceAll(parts[1], "+", "%2B")
	if rawQuery, err := url.ParseQuery(query); err == nil {
		obfsparam, err := utils.DecodeB64(rawQuery.Get("obfsparam"))
		if err != nil {
			return
		}
		ssr.ObfsParam = obfsparam
		if obfsparam == "" {
			obfsparam, err := utils.DecodeB64(rawQuery.Get("obfs-param"))
			if err != nil {
				return
			}
			ssr.ObfsParam = obfsparam
		}
		protoparam, err := utils.DecodeB64(rawQuery.Get("protoparam"))
		if err != nil {
			return
		}
		ssr.ProtocolParam = protoparam
		remarks, err := utils.DecodeB64(rawQuery.Get("remarks"))
		if err == nil {
			if remarks == "" {
				remarks = net.JoinHostPort(ssr.Server, ssr.Port)
			}
			ssr.Name = remarks
			// ssr.Type = remarks
		}
	}
}
