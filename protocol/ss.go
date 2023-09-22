package protocol

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/sanzhang007/webgin/base64decode"
)

type Ss struct {
	Id       int `gorm:"primaryKey"`
	Name     string
	Password string
	Port     string
	Server   string `gorm:"unique"`
	Cipher   string
	Type     string
	Link     string `gorm:"size:1024"`
}

func (ss *Ss) Parse(urlstring string) {
	builder := new(strings.Builder)
	urlSpilt := strings.Split(urlstring, "://")
	builder.WriteString(urlSpilt[0])
	builder.WriteString("://")
	r := regexp.MustCompile("[a-zA-Z0-9+/=]+")
	match := r.FindAllString(urlSpilt[1], -1)[0]
	urlSpilt[1] = strings.Replace(urlSpilt[1], match, base64decode.Base64Decode(match), -1)
	builder.WriteString(base64decode.Base64Decode(urlSpilt[1]))
	urlstring = builder.String()
	u, err := url.Parse(urlstring)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	userString := strings.Split(base64decode.Base64Decode(u.User.String()), ":")

	ss.Name = u.Fragment
	// ss.Name = "unkownName"
	ss.Server = u.Hostname()
	ss.Port = u.Port()
	ss.Type = "ss"
	ss.Link = urlstring

	if len(userString) == 1 {

		b, _ := base64.RawStdEncoding.DecodeString(userString[0])
		tmp := strings.Split(string(b), ":")
		if len(tmp) == 1 {
			return
		}
		ss.Password = tmp[1]
		ss.Cipher = tmp[0]
	} else {
		ss.Cipher = userString[0]
		ss.Password = userString[1]
	}
}
